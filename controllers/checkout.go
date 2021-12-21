package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"project3/lib/databases"
	"project3/middlewares"
	"project3/models"
	"project3/response"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func CreateCheckoutController(c echo.Context) error {
	input := models.CheckOut{}
	user_id := middlewares.ExtractTokenUserId(c)
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}

	if input.Booking_ID == nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status":  "failed",
			"message": "Input Your Booking",
		})
	}

	cart_id, _ := databases.GetCartId(user_id)
	bookings, err := databases.GetBookings(input.Booking_ID, int(cart_id.ID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BookingNotFoundResponse())
	}
	var totPrice int
	for i := 0; i < len(bookings); i++ {
		totPrice += bookings[i].Total
		fmt.Println(totPrice)
	}

	if totPrice > 10000000 && (input.CheckoutMethod == "ID_DANA" || input.CheckoutMethod == "ID_LINKAJA" || input.CheckoutMethod == "ID_SHOPEEPAY") {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "Your Billing Exceed 10,000,000 ",
		})
	}

	input.User_ID = user_id
	input.Total = totPrice
	CheckOut, err := databases.CreateCheckout(input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}

	if input.CheckoutMethod == "" {
		return c.JSON(http.StatusBadRequest, response.CheckOutMissingResponse())
	}

	if input.CheckoutMethod == "COD" {
		return c.JSON(http.StatusOK, response.CheckoutCODSuccessResponse(CheckOut.ID))
	}
	// Payment Xendit
	var body2 = models.RequestBodyStruct{
		ReferenceID:    strconv.Itoa(int(CheckOut.ID)),
		Currency:       "IDR",
		Amount:         float64(CheckOut.Total),
		CheckoutMethod: "ONE_TIME_PAYMENT",
		ChannelCode:    input.CheckoutMethod,
		ChannelProperties: models.ChannelProperties{
			Success_redirect_url: "https://redirect.me/payment",
		},

		Metadata: models.Metadata{
			BranchArea: "PLUIT",
			BranchCity: "JAKARTA",
		},
	}

	reqBody, err := json.Marshal(body2)
	if err != nil {
		print(err)
	}

	req, _ := http.NewRequest(http.MethodPost, "https://api.xendit.co/ewallets/charges", bytes.NewBuffer(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.SetBasicAuth(os.Getenv("SECRET_KEY"), os.Getenv("PASS_XENDIT"))

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body4, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var responsePayment models.ResponsePayment
	json.Unmarshal([]byte(body4), &responsePayment)
	// respon
	return c.JSON(http.StatusOK, response.CheckoutSuccessResponse(CheckOut.ID, responsePayment))
}

func CreateCheckoutOVOController(c echo.Context) error {
	input := models.CheckOut{}
	user_id := middlewares.ExtractTokenUserId(c)
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}

	if input.Booking_ID == nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status":  "failed",
			"message": "Input Your Booking",
		})
	}

	cart_id, _ := databases.GetCartId(user_id)
	bookings, err := databases.GetBookings(input.Booking_ID, int(cart_id.ID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BookingNotFoundResponse())
	}
	var totPrice int
	for i := 0; i < len(bookings); i++ {
		totPrice += bookings[i].Total
		fmt.Println(totPrice)
	}
	if totPrice > 10000000 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "Your Billing Exceed 10,000,000 ",
		})
	}

	input.User_ID = user_id
	input.Total = totPrice
	CheckOut, err := databases.CreateCheckout(input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}

	// Payment Xendit
	var body2 = models.RequestBodyStructOVO{
		ReferenceID:    strconv.Itoa(int(CheckOut.ID)),
		Currency:       "IDR",
		Amount:         float64(CheckOut.Total),
		CheckoutMethod: "ONE_TIME_PAYMENT",
		ChannelCode:    "ID_OVO",
		ChannelProperties: models.ChannelPropertiesOVO{
			MobileNumber: input.Phone,
		},

		Metadata: models.Metadata{
			BranchArea: "PLUIT",
			BranchCity: "JAKARTA",
		},
	}

	reqBody, err := json.Marshal(body2)
	if err != nil {
		print(err)
	}

	req, _ := http.NewRequest(http.MethodPost, "https://api.xendit.co/ewallets/charges", bytes.NewBuffer(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.SetBasicAuth(os.Getenv("SECRET_KEY"), os.Getenv("PASS_XENDIT"))

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body4, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var responsePaymentOVO models.ResponsePaymentOVO
	json.Unmarshal([]byte(body4), &responsePaymentOVO)
	// respon
	return c.JSON(http.StatusOK, response.CheckoutSuccessResponseOVO(CheckOut.ID, responsePaymentOVO))
}
