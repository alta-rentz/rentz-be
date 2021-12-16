package controllers

import (
	"net/http"
	"project3/lib/databases"
	"project3/middlewares"
	"project3/models"
	"project3/plugins"
	"project3/response"
	"regexp"
	"strings"

	"github.com/labstack/echo/v4"
)

func CreateUserControllers(c echo.Context) error {
	user := models.Users{}
	c.Bind(&user)
	spaceEmpty := strings.TrimSpace(user.Nama)
	if user.Nama == "" && user.Email == "" && user.Password == "" && user.Phone_Number == "" {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	if user.Password == "" {
		return c.JSON(http.StatusBadRequest, response.PasswordCannotEmpty())
	}
	if len(user.Password) < 5 {
		return c.JSON(http.StatusBadRequest, response.PasswordCannotLess5())
	}
	newPass, _ := plugins.Encrypt(user.Password)
	user.Password = newPass
	if spaceEmpty == "" {
		return c.JSON(http.StatusBadRequest, response.NameCannotEmpty())
	}
	if user.Email == "" {
		return c.JSON(http.StatusBadRequest, response.EmailCannotEmpty())
	}
	if user.Phone_Number == "" {
		return c.JSON(http.StatusBadRequest, response.PhoneNumberCannotEmpty())
	}
	pattern := `^\w+@\w+\.\w+$`
	matched, _ := regexp.Match(pattern, []byte(user.Email))
	if !matched {
		return c.JSON(http.StatusBadRequest, response.FormatEmailInvalid())
	}
	createdUser, err := databases.CreateUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.IsExist())
	}
	cart := models.Cart{
		UsersID: createdUser.ID,
	}
	_, err = databases.CreateCartUser(&cart)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
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
	pattern := `^\w+@\w+\.\w+$`
	matched, _ := regexp.Match(pattern, []byte(user.Email))
	if user.Email == "" && user.Password == "" {
		return c.JSON(http.StatusBadRequest, response.EmailPasswordCannotEmpty())
	} else if user.Email == "" || user.Password == "" {
		return c.JSON(http.StatusBadRequest, response.EmailPasswordCannotEmpty())
	} else if !matched {
		return c.JSON(http.StatusBadRequest, response.FormatEmailInvalid())
	} else if err != nil || users == 0 {
		return c.JSON(http.StatusBadRequest, response.LoginFailedResponse())
	}

	return c.JSON(http.StatusOK, response.LoginSuccessResponse(users))
}

func GetUserControllersTest() echo.HandlerFunc {
	return GetUserControllers
}

func UpdateUserControllersTest() echo.HandlerFunc {
	return UpdateUserControllers
}

func DeleteUserControllersTest() echo.HandlerFunc {
	return DeleteUserControllers
}
