package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"project3/config"
	"project3/constant"
	"project3/middlewares"
	"project3/models"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

type UserResponse struct {
	Message string
	Data    models.Users
}

type Login struct {
	Email    string
	Password string
}

// data dummy
var (
	mock_data_user = models.Users{
		Nama:         "andri",
		Email:        "andri@gmail.com",
		Password:     "bismillah",
		Phone_Number: "081296620776",
	}
	mock_data_user_pass_error = models.Users{
		Nama:         "andri",
		Email:        "andri@gmail.com",
		Password:     "123",
		Phone_Number: "081296620776",
	}
	mock_data_user_pass_empty = models.Users{
		Nama:         "andri",
		Email:        "andri@gmail.com",
		Password:     "",
		Phone_Number: "081296620776",
	}
	mock_data_user_phone_empty = models.Users{
		Nama:         "andri",
		Email:        "andri@gmail.com",
		Password:     "bismillah",
		Phone_Number: "",
	}
	mock_data_user_name_error = models.Users{
		Nama:         "",
		Email:        "andri@gmail.com",
		Password:     "bismillah",
		Phone_Number: "081296620776",
	}
	mock_data_user_email_error = models.Users{
		Nama:         "andri",
		Email:        "",
		Password:     "bismillah",
		Phone_Number: "081296620776",
	}
	mock_data_user_email_format_error = models.Users{
		Nama:         "andri",
		Email:        "andri",
		Password:     "bismillah",
		Phone_Number: "081296620776",
	}
	mock_data_user_email_pass_empty = models.Users{
		Nama:         "andri",
		Email:        "",
		Password:     "",
		Phone_Number: "081296620776",
	}
	mock_data_user_all_empty = models.Users{
		Nama:         "",
		Email:        "",
		Password:     "",
		Phone_Number: "",
	}
	mock_data_login = models.Users{
		Email:    "andri@gmail.com",
		Password: "bismillah",
	}
	mock_data_login_fail = models.Users{
		Email:    "andri@gmail.com",
		Password: "abcdefg",
	}
	mock_data_login_email_error = models.Users{
		Email:    "andri",
		Password: "abcdefg",
	}
	mock_data_login_email_empty = models.Users{
		Email:    "",
		Password: "abcdefg",
	}
	mock_data_login_pass_empty = models.Users{
		Email:    "andri",
		Password: "",
	}
	mock_data_login_email_pass_empty = models.Users{
		Email:    "",
		Password: "",
	}
)

// inisialisasi echo
func InitEcho() *echo.Echo {
	config.InitDBTest()
	e := echo.New()

	return e
}

// menambahkan user
func InsertUser() error {
	if err := config.DB.Save(&mock_data_user).Error; err != nil {
		return err
	}
	return nil
}

// test create user success
func TestCreateUserController(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "Successful Operation",
		path:       "/signup",
		expectCode: http.StatusOK,
	}

	e := InitEcho()

	body, err := json.Marshal(mock_data_user)
	if err != nil {
		t.Error(t, err, "error")
	}

	// send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPost, testCases.path, bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, CreateUserControllers(c)) {
		bodyrecponses := rec.Body.String()
		var user UserResponse

		err := json.Unmarshal([]byte(bodyrecponses), &user)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.expectCode, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	}
}

// test create user password error
func TestCreateUserControllerPasswordError(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "password cannot less than 5 character",
		path:       "/signup",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()

	body, err := json.Marshal(mock_data_user_pass_error)
	if err != nil {
		t.Error(t, err, "error")
	}

	// send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPost, testCases.path, bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, CreateUserControllers(c)) {
		bodyrecponses := rec.Body.String()
		var user UserResponse

		err := json.Unmarshal([]byte(bodyrecponses), &user)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.expectCode, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	}
}

// test create user name error
func TestCreateUserControllerNameError(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "name cannot be empty",
		path:       "/signup",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()

	body, err := json.Marshal(mock_data_user_name_error)
	if err != nil {
		t.Error(t, err, "error")
	}

	// send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPost, testCases.path, bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, CreateUserControllers(c)) {
		bodyrecponses := rec.Body.String()
		var user UserResponse

		err := json.Unmarshal([]byte(bodyrecponses), &user)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.expectCode, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	}
}

// test create user email error
func TestCreateUserControllerEmailError(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "email cannot be empty",
		path:       "/signup",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()

	body, err := json.Marshal(mock_data_user_email_error)
	if err != nil {
		t.Error(t, err, "error")
	}

	// send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPost, testCases.path, bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, CreateUserControllers(c)) {
		bodyrecponses := rec.Body.String()
		var user UserResponse

		err := json.Unmarshal([]byte(bodyrecponses), &user)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.expectCode, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	}
}

// test create user email format error
func TestCreateUserControllerEmailFormatError(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "Email must contain email format",
		path:       "/signup",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()

	body, err := json.Marshal(mock_data_user_email_format_error)
	if err != nil {
		t.Error(t, err, "error")
	}

	// send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPost, testCases.path, bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, CreateUserControllers(c)) {
		bodyrecponses := rec.Body.String()
		var user UserResponse

		err := json.Unmarshal([]byte(bodyrecponses), &user)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.expectCode, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	}
}

// test create user phone empty
func TestCreateUserControllerPhoneEmpty(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "phone number cannot be empty",
		path:       "/signup",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()

	body, err := json.Marshal(mock_data_user_phone_empty)
	if err != nil {
		t.Error(t, err, "error")
	}

	// send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPost, testCases.path, bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, CreateUserControllers(c)) {
		bodyrecponses := rec.Body.String()
		var user UserResponse

		err := json.Unmarshal([]byte(bodyrecponses), &user)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.expectCode, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	}
}

// test create pass pass empty
func TestCreateUserControllerPassEmpty(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "password cannot be empty",
		path:       "/signup",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()

	body, err := json.Marshal(mock_data_user_pass_empty)
	if err != nil {
		t.Error(t, err, "error")
	}

	// send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPost, testCases.path, bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, CreateUserControllers(c)) {
		bodyrecponses := rec.Body.String()
		var user UserResponse

		err := json.Unmarshal([]byte(bodyrecponses), &user)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.expectCode, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	}
}

// test create user isExist
func TestCreateUserControllerIsExist(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "Email or Phone Number is Exist",
		path:       "/signup",
		expectCode: http.StatusInternalServerError,
	}

	e := InitEcho()
	InsertUser()

	body, err := json.Marshal(mock_data_user)
	if err != nil {
		t.Error(t, err, "error")
	}

	// send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPost, testCases.path, bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, CreateUserControllers(c)) {
		bodyrecponses := rec.Body.String()
		var user UserResponse

		err := json.Unmarshal([]byte(bodyrecponses), &user)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.expectCode, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	}
}

// test create user all empty
func TestCreateUserControllerAllEmpty(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "Bad Request",
		path:       "/signup",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()
	InsertUser()

	body, err := json.Marshal(mock_data_user_all_empty)
	if err != nil {
		t.Error(t, err, "error")
	}

	// send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPost, testCases.path, bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, CreateUserControllers(c)) {
		bodyrecponses := rec.Body.String()
		var user UserResponse

		err := json.Unmarshal([]byte(bodyrecponses), &user)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.expectCode, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	}
}

// test login user success
// func TestLoginUserController(t *testing.T) {
// 	var testCases = struct {
// 		name       string
// 		path       string
// 		expectCode int
// 	}{

// 		name:       "Successful Operation",
// 		path:       "/signin",
// 		expectCode: http.StatusOK,
// 	}

// 	e := InitEcho()
// 	InsertUser()

// 	body, err := json.Marshal(mock_data_login)
// 	if err != nil {
// 		t.Error(t, err, "error")
// 	}

// 	// send data using request body with HTTP Method POST
// 	req := httptest.NewRequest(http.MethodPost, testCases.path, bytes.NewBuffer(body))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()

// 	c := e.NewContext(req, rec)

// 	if assert.NoError(t, LoginUsersController(c)) {
// 		bodyrecponses := rec.Body.String()
// 		var user UserResponse

// 		err := json.Unmarshal([]byte(bodyrecponses), &user)
// 		if err != nil {
// 			assert.Error(t, err, "error")
// 		}

// 		assert.Equal(t, testCases.expectCode, rec.Code)
// 		assert.Equal(t, testCases.name, user.Message)
// 	}
// }

// test login user fail
func TestLoginUserControllerFail(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "Email or Password Invalid",
		path:       "/signin",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()
	InsertUser()

	body, err := json.Marshal(mock_data_login_fail)
	if err != nil {
		t.Error(t, err, "error")
	}

	// send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPost, testCases.path, bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, LoginUsersController(c)) {
		bodyrecponses := rec.Body.String()
		var user UserResponse

		err := json.Unmarshal([]byte(bodyrecponses), &user)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.expectCode, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	}
}

// test create user login format error
func TestLoginUserControllerEmailFormatError(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "Email must contain email format",
		path:       "/signin",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()

	body, err := json.Marshal(mock_data_login_email_error)
	if err != nil {
		t.Error(t, err, "error")
	}

	// send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPost, testCases.path, bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, LoginUsersController(c)) {
		bodyrecponses := rec.Body.String()
		var user UserResponse

		err := json.Unmarshal([]byte(bodyrecponses), &user)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.expectCode, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	}
}

// test create user login email pass empty
func TestLoginUserControllerEmailPasswordEmpty(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "email or password cannot be empty",
		path:       "/signin",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()

	body, err := json.Marshal(mock_data_login_email_pass_empty)
	if err != nil {
		t.Error(t, err, "error")
	}

	// send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPost, testCases.path, bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, LoginUsersController(c)) {
		bodyrecponses := rec.Body.String()
		var user UserResponse

		err := json.Unmarshal([]byte(bodyrecponses), &user)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.expectCode, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	}
}

// test create user login email empty
func TestLoginUserControllerEmailEmpty(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "email or password cannot be empty",
		path:       "/signin",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()

	body, err := json.Marshal(mock_data_login_email_empty)
	if err != nil {
		t.Error(t, err, "error")
	}

	// send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPost, testCases.path, bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, LoginUsersController(c)) {
		bodyrecponses := rec.Body.String()
		var user UserResponse

		err := json.Unmarshal([]byte(bodyrecponses), &user)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.expectCode, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	}
}

// test create user login pass empty
func TestLoginUserControllerPasswordEmpty(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "email or password cannot be empty",
		path:       "/signin",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()

	body, err := json.Marshal(mock_data_user_pass_empty)
	if err != nil {
		t.Error(t, err, "error")
	}

	// send data using request body with HTTP Method POST
	req := httptest.NewRequest(http.MethodPost, testCases.path, bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, LoginUsersController(c)) {
		bodyrecponses := rec.Body.String()
		var user UserResponse

		err := json.Unmarshal([]byte(bodyrecponses), &user)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.expectCode, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	}
}

// test get user by jwt success
func TestGetUserControllers(t *testing.T) {
	testCases := struct {
		name string
		path string
		code int
	}{

		name: "Successful Operation",
		path: "jwt/users",
		code: http.StatusOK,
	}

	e := InitEcho()
	InsertUser()
	var userDB models.Users
	tx := config.DB.Where("email = ? AND password = ?", mock_data_login.Email, mock_data_login.Password).First(&userDB)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDB.ID))
	if err != nil {
		panic(err)
	}
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetPath(testCases.path)
	middleware.JWT([]byte(constant.SECRET_JWT))(GetUserControllersTest())(context)

	var user UserResponse
	rec_body := rec.Body.String()
	json.Unmarshal([]byte(rec_body), &user)
	if err != nil {
		assert.Error(t, err, "error")
	}

	t.Run("GET /jwt/users", func(t *testing.T) {
		assert.Equal(t, testCases.code, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
		assert.Equal(t, "andri", user.Data.Nama)
	})

}

// test get user by jwt fail
func TestGetUserControllersFail(t *testing.T) {
	testCases := struct {
		name string
		path string
		code int
	}{

		name: "Bad Request",
		path: "jwt/users",
		code: http.StatusBadRequest,
	}

	e := InitEcho()
	InsertUser()
	var userDB models.Users
	tx := config.DB.Where("email = ? AND password = ?", mock_data_login.Email, mock_data_login.Password).First(&userDB)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDB.ID))
	if err != nil {
		panic(err)
	}
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetPath(testCases.path)
	config.DB.Migrator().DropTable(models.Users{})
	middleware.JWT([]byte(constant.SECRET_JWT))(GetUserControllersTest())(context)

	var user UserResponse
	rec_body := rec.Body.String()
	json.Unmarshal([]byte(rec_body), &user)
	if err != nil {
		assert.Error(t, err, "error")
	}

	t.Run("GET /jwt/users", func(t *testing.T) {
		assert.Equal(t, testCases.code, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	})

}

// test update user by jwt success
func TestUpdateUserControllers(t *testing.T) {
	testCases := struct {
		name string
		path string
		code int
	}{

		name: "Successful Operation",
		path: "jwt/users",
		code: http.StatusOK,
	}

	e := InitEcho()
	InsertUser()
	var userDB models.Users
	tx := config.DB.Where("email = ? AND password = ?", mock_data_login.Email, mock_data_login.Password).First(&userDB)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDB.ID))
	if err != nil {
		panic(err)
	}
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetPath(testCases.path)
	middleware.JWT([]byte(constant.SECRET_JWT))(UpdateUserControllersTest())(context)

	var user UserResponse
	rec_body := rec.Body.String()
	json.Unmarshal([]byte(rec_body), &user)
	if err != nil {
		assert.Error(t, err, "error")
	}

	t.Run("UPDATE /jwt/users", func(t *testing.T) {
		assert.Equal(t, testCases.code, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	})

}

// test update user by jwt fail
func TestUpdateUserControllersFail(t *testing.T) {
	testCases := struct {
		name string
		path string
		code int
	}{

		name: "Bad Request",
		path: "jwt/users",
		code: http.StatusBadRequest,
	}

	e := InitEcho()
	InsertUser()
	var userDB models.Users
	tx := config.DB.Where("email = ? AND password = ?", mock_data_login.Email, mock_data_login.Password).First(&userDB)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDB.ID))
	if err != nil {
		panic(err)
	}
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetPath(testCases.path)
	config.DB.Migrator().DropTable(models.Users{})
	middleware.JWT([]byte(constant.SECRET_JWT))(UpdateUserControllersTest())(context)

	var user UserResponse
	rec_body := rec.Body.String()
	json.Unmarshal([]byte(rec_body), &user)
	if err != nil {
		assert.Error(t, err, "error")
	}

	t.Run("UPDATE /jwt/users", func(t *testing.T) {
		assert.Equal(t, testCases.code, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	})

}

// test delete user by jwt success
func TestDeleteUserControllers(t *testing.T) {
	testCases := struct {
		name string
		path string
		code int
	}{

		name: "Successful Operation",
		path: "jwt/users",
		code: http.StatusOK,
	}

	e := InitEcho()
	InsertUser()
	var userDB models.Users
	tx := config.DB.Where("email = ? AND password = ?", mock_data_login.Email, mock_data_login.Password).First(&userDB)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDB.ID))
	if err != nil {
		panic(err)
	}
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetPath(testCases.path)
	middleware.JWT([]byte(constant.SECRET_JWT))(DeleteUserControllersTest())(context)

	var user UserResponse
	rec_body := rec.Body.String()
	json.Unmarshal([]byte(rec_body), &user)
	if err != nil {
		assert.Error(t, err, "error")
	}

	t.Run("DELETE /jwt/users", func(t *testing.T) {
		assert.Equal(t, testCases.code, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	})

}

// test delete user by jwt fail
func TestDeleteUserControllersFail(t *testing.T) {
	testCases := struct {
		name string
		path string
		code int
	}{

		name: "Bad Request",
		path: "jwt/users",
		code: http.StatusBadRequest,
	}

	e := InitEcho()
	InsertUser()
	var userDB models.Users
	tx := config.DB.Where("email = ? AND password = ?", mock_data_login.Email, mock_data_login.Password).First(&userDB)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDB.ID))
	if err != nil {
		panic(err)
	}
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetPath(testCases.path)
	config.DB.Migrator().DropTable(models.Users{})
	middleware.JWT([]byte(constant.SECRET_JWT))(DeleteUserControllersTest())(context)

	var user UserResponse
	rec_body := rec.Body.String()
	json.Unmarshal([]byte(rec_body), &user)
	if err != nil {
		assert.Error(t, err, "error")
	}

	t.Run("DELETE /jwt/users", func(t *testing.T) {
		assert.Equal(t, testCases.code, rec.Code)
		assert.Equal(t, testCases.name, user.Message)
	})

}
