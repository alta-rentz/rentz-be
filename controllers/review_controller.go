package controllers

import (
	"net/http"
	"project3/lib/databases"
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
	if review.Rating <= 0 || review.Rating > 5 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "please choose between 1 - 5",
		})
	}
	productID, err := databases.GetProductID(int(review.BookingID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	review.ProductsID = uint(productID)
	_, err = databases.AddReviews(&review)
	databases.AddRatingToProduct(int(review.ProductsID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}
