package controllers

import (
	"net/http"
	"project3/lib/databases"
	"project3/middlewares"
	"project3/models"
	"project3/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddReviewsController(c echo.Context) error {
	bookingId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	var review models.Reviews
	c.Bind(&review)
	review.BookingID = uint(bookingId)
	userId := middlewares.ExtractTokenUserId(c)
	bookingOwner, err := databases.GetBookingOwner(int(review.BookingID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}

	if bookingOwner != userId {
		return c.JSON(http.StatusBadRequest, response.AccessForbiddenResponse())
	}
	bookingStatus, _ := databases.GetBookingStatus(int(review.BookingID))
	if bookingStatus == "waiting" {
		return c.JSON(http.StatusBadRequest, response.AccessForbiddenResponse())
	}

	if review.Rating <= 0 || review.Rating > 5 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "please choose between 1 - 5",
		})
	}
	productID, _ := databases.GetProductID(int(review.BookingID))
	review.ProductsID = uint(productID)
	_, err = databases.AddReviews(&review)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "you have already reviewed this booking",
		})
	}
	databases.AddRatingToProduct(int(review.ProductsID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}
