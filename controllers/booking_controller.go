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
	var user models.Users
	var productUser models.Products
	if user.ID == productUser.UsersID {
		return c.JSON(http.StatusBadRequest, response.CannotBookingSelfProductResponse())
	}

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

	rent, err := databases.CreateBooking(input, int(cart.ID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}

	databases.AccumulatedDays(input.Time_In, input.Time_Out, rent.ID)
	input.Total = databases.AddPriceBooking(input.ProductsID, rent.ID)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success create new booking",
		"idBook":  rent.ID,
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
	if uint(logged) == 0 {
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
