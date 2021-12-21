package routes

import (
	"net/http"

	"project3/constant"
	"project3/controllers"

	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {

	e := echo.New()
	e.Use(m.CORSWithConfig(m.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// User Sign Up & Sign In
	e.POST("/signup", controllers.CreateUserControllers)
	e.POST("/signin", controllers.LoginUsersController)
	e.GET("/products", controllers.GetAllProductsController)
	e.GET("/products/:id", controllers.GetProductByIDController)
	e.GET("/products/subcategory/:id", controllers.GetProductsBySubcategoryIDController)
	e.POST("/booking/check/:id", controllers.ProductRentCheckController)

	// JWT Group
	r := e.Group("/jwt")
	r.Use(m.JWT([]byte(constant.SECRET_JWT)))

	// Users JWT
	r.GET("/users", controllers.GetUserControllers)
	r.PUT("/users", controllers.UpdateUserControllers)
	r.DELETE("/users", controllers.DeleteUserControllers)

	// Product JWT
	r.POST("/products", controllers.CreateProductControllers)
	r.GET("/products", controllers.GetProductsByUserIDController)
	r.DELETE("/products/:id", controllers.DeleteProductByIDController)

	// Booking JWt
	r.POST("/booking", controllers.CreateBookingControllers)
	r.GET("/booking/:id", controllers.GetBookingByIdController)
	r.GET("/cart", controllers.GetBookingByCartIDController)
	r.DELETE("/booking/:id", controllers.CancelBookingController)

	//Review JWT
	r.POST("/reviews/:id", controllers.AddReviewsController)

	// CheckOut JWT
	r.POST("/checkout", controllers.CreateCheckoutController)
	r.POST("/checkout/ovo", controllers.CreateCheckoutOVOController)

	return e
}
