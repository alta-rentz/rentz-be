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

	// JWT Group
	r := e.Group("/jwt")
	r.Use(m.JWT([]byte(constant.SECRET_JWT)))

	// Users JWT
	r.GET("/users", controllers.GetUserControllers)
	r.PUT("/users", controllers.UpdateUserControllers)
	r.DELETE("/users", controllers.DeleteUserControllers)

	return e
}
