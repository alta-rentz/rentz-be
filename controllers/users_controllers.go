package controllers

import (
	"net/http"
	"os"
	"project3/lib/databases"
	"project3/middlewares"
	"project3/models"
	"project3/plugins"
	"project3/response"
	"regexp"

	"github.com/labstack/echo/v4"
)

func CreateUserControllers(c echo.Context) error {
	user := models.Users{}
	c.Bind(&user)
	if len(user.Password) < 5 {
		return c.JSON(http.StatusBadRequest, response.PasswordCannotLess5())
	}
	newPass, _ := plugins.Encrypt(user.Password)
	user.Password = newPass
	if user.Nama == "" {
		return c.JSON(http.StatusBadRequest, response.NameCannotEmpty())
	}
	if user.Email == "" {
		return c.JSON(http.StatusBadRequest, response.EmailCannotEmpty())
	}
	pattern := `^\w+@\w+\.\w+$`
	matched, tx := regexp.Match(pattern, []byte(user.Email))
	if tx != nil {
		os.Exit(1)
		return c.JSON(http.StatusBadRequest, response.FormatEmailInvalid())
	}
	if !matched {
		return c.JSON(http.StatusBadRequest, response.FormatEmailInvalid())
	}
	_, err := databases.CreateUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.IsExist())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

func GetUserControllers(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)
	user, err := databases.GetUser(id)
	if err != nil || user == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData(user))
}

func UpdateUserControllers(c echo.Context) error {
	user := models.Users{}
	id := middlewares.ExtractTokenUserId(c)
	c.Bind(&user)
	_, err := databases.UpdateUser(id, &user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

func DeleteUserControllers(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)
	_, err := databases.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

//login users
func LoginUsersController(c echo.Context) error {
	user := models.UserLogin{}
	c.Bind(&user)
	users, err := databases.LoginUser(user)
	if err != nil || users == 0 {
		return c.JSON(http.StatusBadRequest, response.LoginFailedResponse())
	}

	return c.JSON(http.StatusOK, response.LoginSuccessResponse(users))
}

func GetUserControllerTest() echo.HandlerFunc {
	return GetUserControllers
}

func UpdateUserControllerTest() echo.HandlerFunc {
	return UpdateUserControllers
}

func DeleteUserControllerTest() echo.HandlerFunc {
	return DeleteUserControllers
}
