package controllers

import (
	"net/http"
	"project3/lib/databases"
	"project3/middlewares"
	"project3/models"
	"project3/response"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

var format_date string = "2006-01-02"

// Membuat reservasi baru
func CreateBookingControllers(c echo.Context) error {
	body := models.BookingBody{}
	c.Bind(&body)
	logged := middlewares.ExtractTokenUserId(c)

	input := models.Booking{}

	cart, _ := databases.GetCartId(logged)
	input.CartID = uint(cart.ID)
	input.ProductsID = body.ProductsID
	input.Qty = body.Qty
	input.TransactionID = nil
	input.Time_In, _ = time.Parse(format_date, body.Time_In)
	input.Time_Out, _ = time.Parse(format_date, body.Time_Out)
	if input.Time_In.Unix() > input.Time_Out.Unix() {
		return c.JSON(http.StatusBadRequest, response.DateInvalidResponse())
	}

	// Mendapatkan seluruh tanggal rental product tertentu
	dateList, err := databases.ProductRentList(int(input.ProductsID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}

	// Pengecekan ketersediaan product untuk tanggal time_in dan time_out yang diinginkan
	for _, date := range dateList {
		if (input.Time_In.Unix() >= date.Time_In.Unix() && input.Time_In.Unix() <= date.Time_Out.Unix()) || (input.Time_Out.Unix() >= date.Time_In.Unix() && input.Time_Out.Unix() <= date.Time_Out.Unix()) {
			return c.JSON(http.StatusBadRequest, response.CheckFailedResponse())
		}
	}

	ownProduct, _ := databases.GetProductOwner(int(body.ProductsID))
	if logged == ownProduct {
		return c.JSON(http.StatusBadRequest, response.BookingOwnProductsFailed())
	}

	rent, err := databases.CreateBooking(input, int(cart.ID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}

	databases.AccumulatedDays(input.Time_In, input.Time_Out, rent.ID)
	input.Total = databases.AddPriceBooking(input.ProductsID, rent.ID)

	hasil, _ := databases.GetBookingById(int(rent.ID))
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success create new booking",
		"idBook":  hasil,
	})
}

func GetBookingByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	rent, err := databases.GetBookingById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	if rent == nil {
		return c.JSON(http.StatusBadRequest, response.BookingNotFoundResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData(rent))
}

func CancelBookingController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	logged := middlewares.ExtractTokenUserId(c)
	bookingOwner, err := databases.GetBookingOwner(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	if bookingOwner != logged {
		return c.JSON(http.StatusBadRequest, response.AccessForbiddenResponse())
	}
	databases.CancelBooking(id)
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

func GetBookingByCartIDController(c echo.Context) error {
	logged := middlewares.ExtractTokenUserId(c)
	cart, err := databases.GetCartId(logged)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	rent, err := databases.GetBookingByCartID(int(cart.ID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	if rent == nil {
		return c.JSON(http.StatusBadRequest, response.BookingNotFoundResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData(rent))
}

func GetHistoryByCartIDController(c echo.Context) error {
	logged := middlewares.ExtractTokenUserId(c)
	cart, err := databases.GetCartId(logged)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	rent, err := databases.GetHistoryByCartID(int(cart.ID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	if rent == nil {
		return c.JSON(http.StatusBadRequest, response.BookingNotFoundResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData(rent))
}

type BodyDate struct {
	Time_In  string `json:"time_in" form:"time_in"`
	Time_Out string `json:"time_out" form:"time_out"`
}

type InputDate struct {
	Time_In  time.Time `json:"time_in" form:"time_in"`
	Time_Out time.Time `json:"time_out" form:"time_out"`
}

// Fungsi untuk melakukan pengecekan availability suatu product
func ProductRentCheckController(c echo.Context) error {
	body := BodyDate{}
	c.Bind(&body)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	// Mendapatkan seluruh tanggal rental product tertentu
	dateList, err := databases.ProductRentList(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}

	input := InputDate{}
	input.Time_In, _ = time.Parse(format_date, body.Time_In)
	input.Time_Out, _ = time.Parse(format_date, body.Time_Out)

	if input.Time_In.Unix() < time.Now().Unix() || input.Time_Out.Unix() < time.Now().Unix() {
		return c.JSON(http.StatusBadRequest, response.DateInvalidResponse())
	}
	if input.Time_In.Unix() > input.Time_Out.Unix() {
		return c.JSON(http.StatusBadRequest, response.DateInvalidResponse())
	}
	// Pengecekan ketersediaan product untuk tanggal time_in dan time_out yang diinginkan
	for _, date := range dateList {
		if (input.Time_In.Unix() >= date.Time_In.Unix() && input.Time_In.Unix() <= date.Time_Out.Unix()) || (input.Time_Out.Unix() >= date.Time_In.Unix() && input.Time_Out.Unix() <= date.Time_Out.Unix()) {
			return c.JSON(http.StatusBadRequest, response.CheckFailedResponse())
		}
	}
	return c.JSON(http.StatusOK, response.CheckSuccessResponse())
}
