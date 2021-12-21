package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"project3/config"
	"project3/constant"
	"project3/middlewares"
	"project3/models"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

// Struct yang digunakan ketika test request success, dapat menampung banyak data
type CartBookResponse struct {
	Code    string
	Message string
}

// Struct untuk menampung data test case
type TestCaseCart struct {
	Name       string
	Path       string
	ExpectCode int
}

// data dummy
var (
	mock_data_province_c = models.Province{
		Province_Name: "Jawa Barat",
	}
	mock_data_city_c = models.City{
		City_Name:  "Subang",
		ProvinceID: 1,
	}
	mock_data_category_c = models.Category{
		Category_Name: "Elektronik",
	}
	mock_data_subcategory_c = models.Subcategory{
		Subcategory_Name: "Kamera",
		CategoryID:       1,
	}
	mock_data_guarantee_c = models.Guarantee{
		Guarantee_Name: "KK",
	}
	mock_data_product_c = models.Products{
		Name:          "Kamera DSLR Canon",
		UsersID:       1,
		SubcategoryID: 1,
		CityID:        1,
		Price:         50000,
		Description:   "Murah yang terbaik",
		Stock:         5,
	}
	mock_data_product2_c = models.Products{
		Name:          "Kamera DSLR Nikon",
		UsersID:       2,
		SubcategoryID: 1,
		CityID:        1,
		Price:         50000,
		Description:   "Murah yang terbaik",
		Stock:         5,
	}
	mock_data_photo_c = models.Photos{
		Photo_Name: "inicontohfoto.jpg",
		Url:        "https://googlecloud/inicontohfoto.jgp",
		ProductsID: 1,
	}
	mock_data_product_guarantee_c = models.ProductsGuarantee{
		GuaranteeID: 1,
		ProductsID:  1,
	}
	mock_data_user_c = models.Users{
		Nama:         "alfy",
		Email:        "alfy@gmail.com",
		Password:     "12345678",
		Phone_Number: "081296620776",
	}
	mock_data_user2_c = models.Users{
		Nama:         "alfy1",
		Email:        "alfy1@gmail.com",
		Password:     "12345678",
		Phone_Number: "0819866620776",
	}
	mock_data_login_product_c = models.Users{
		Email:    "alfy@gmail.com",
		Password: "12345678",
	}
	mock_data_cart = models.Cart{
		UsersID: 1,
	}
	time1             = time.Now()
	time2             = time.Now().AddDate(0, 0, 2)
	mock_data_booking = models.Booking{
		ProductsID:     1,
		CartID:         1,
		Time_In:        time1,
		Time_Out:       time2,
		Total_Day:      2,
		Qty:            1,
		Total:          100000,
		Status_Payment: "waiting",
	}
	mock_data_transaction = models.Transaction{
		Total: 100000,
	}
	mock_data_checkoutmethod = models.CheckoutMethod{
		ID:                   1,
		Checkout_Name:        "ovo",
		CheckoutMethodTypeID: 1,
	}
	mock_data_methodtype = models.CheckoutMethodType{
		ID:            1,
		Checkout_Type: "kredit",
	}
)

// Fungsi untuk memasukkan data user test ke dalam database
func InsertMockDataUsersCToDB() error {
	query := config.DB.Save(&mock_data_user_c)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

// Fungsi untuk memasukkan data user test ke dalam database
func InsertMockDataCToDB() {
	config.DB.Save(&mock_data_category_c)
	config.DB.Save(&mock_data_subcategory_c)
	config.DB.Save(&mock_data_province_c)
	config.DB.Save(&mock_data_city_c)
	config.DB.Save(&mock_data_guarantee_c)
	config.DB.Save(&mock_data_methodtype)
	config.DB.Save(&mock_data_checkoutmethod)
}

// // inisialisasi echo
// func InitEcho() *echo.Echo {
// 	config.InitDBTest()
// 	e := echo.New()

// 	return e
// }

// Fungsi untuk melakukan login dan ekstraksi token JWT
func UsingJWTC() (string, error) {
	// Melakukan login data user test
	InsertMockDataUsersCToDB()
	var user models.Users
	tx := config.DB.Where("email = ? AND password = ?", mock_data_login_product_c.Email, mock_data_login_product_c.Password).First(&user)
	if tx.Error != nil {
		return "", tx.Error
	}
	// Mengektraksi token data user test
	token, err := middlewares.CreateToken(int(user.ID))
	if err != nil {
		return "", err
	}
	return token, nil
}

// Fungsi testing CreateProductController
func GetBookingByCartIDControllerTesting() echo.HandlerFunc {
	return GetBookingByCartIDController
}

// Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// kondisi request success
func TestGetBookingByCartIDControllerSuccess(t *testing.T) {
	var testCase = TestCaseCart{
		Name:       "success to get booking by cart id",
		Path:       "/cart",
		ExpectCode: http.StatusOK,
	}

	e := InitEcho()
	// Mendapatkan token
	token, err := UsingJWTC()
	if err != nil {
		panic(err)
	}

	InsertMockDataCToDB()
	config.DB.Save(&mock_data_user2_c)
	config.DB.Save(&mock_data_product_c)
	config.DB.Save(&mock_data_product2_c)
	config.DB.Save(&mock_data_cart)
	config.DB.Save(&mock_data_checkoutmethod)
	config.DB.Save(&mock_data_methodtype)
	config.DB.Save(&mock_data_transaction)
	config.DB.Save(&mock_data_booking)

	req := httptest.NewRequest(http.MethodGet, "/cart", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCase.Path)
	middleware.JWT([]byte(constant.SECRET_JWT))(GetBookingByCartIDControllerTesting())(context)

	res_body := res.Body.String()
	var response CartBookResponse
	er := json.Unmarshal([]byte(res_body), &response)
	if er != nil {
		assert.Error(t, er, "error")
	}
	t.Run("GET /jwt/cart", func(t *testing.T) {
		assert.Equal(t, testCase.ExpectCode, res.Code)
		assert.Equal(t, "Successful Operation", response.Message)
	})

}

// Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// kondisi request failed
func TestGetBookingByCartIDControllerNotFound(t *testing.T) {
	var testCase = TestCaseCart{
		Name:       "get booking not found",
		Path:       "/cart",
		ExpectCode: http.StatusBadRequest,
	}

	e := InitEcho()
	// Mendapatkan token
	token, err := UsingJWTC()
	if err != nil {
		panic(err)
	}

	req := httptest.NewRequest(http.MethodGet, "/cart", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCase.Path)
	middleware.JWT([]byte(constant.SECRET_JWT))(GetBookingByCartIDControllerTesting())(context)

	res_body := res.Body.String()
	var response CartBookResponse
	er := json.Unmarshal([]byte(res_body), &response)
	if er != nil {
		assert.Error(t, er, "error")
	}
	t.Run("GET /jwt/cart", func(t *testing.T) {
		assert.Equal(t, testCase.ExpectCode, res.Code)
		assert.Equal(t, "Booking not found", response.Message)
	})

}

// Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// kondisi request failed
func TestGetBookingByCartIDControllerFailed(t *testing.T) {
	var testCase = TestCaseCart{
		Name:       "failed to get booking",
		Path:       "/cart",
		ExpectCode: http.StatusBadRequest,
	}

	e := InitEcho()
	// Mendapatkan token
	token, err := UsingJWTC()
	if err != nil {
		panic(err)
	}

	// Melakukan penghapusan tabel untuk membuat request failed
	config.DB.Migrator().DropTable(&models.Booking{})

	req := httptest.NewRequest(http.MethodGet, "/cart", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCase.Path)
	middleware.JWT([]byte(constant.SECRET_JWT))(GetBookingByCartIDControllerTesting())(context)

	res_body := res.Body.String()
	var response CartBookResponse
	er := json.Unmarshal([]byte(res_body), &response)
	if er != nil {
		assert.Error(t, er, "error")
	}
	t.Run("GET /jwt/cart", func(t *testing.T) {
		assert.Equal(t, testCase.ExpectCode, res.Code)
		assert.Equal(t, "Bad Request", response.Message)
	})

}

// Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// kondisi request failed
func TestGetBookingByCartIDControllerCartNotFound(t *testing.T) {
	var testCase = TestCaseCart{
		Name:       "failed to get booking",
		Path:       "/cart",
		ExpectCode: http.StatusBadRequest,
	}

	e := InitEcho()
	// Mendapatkan token
	token, err := UsingJWTC()
	if err != nil {
		panic(err)
	}

	// Melakukan penghapusan tabel untuk membuat request failed
	config.DB.Migrator().DropTable(&models.Cart{})

	req := httptest.NewRequest(http.MethodGet, "/cart", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCase.Path)
	middleware.JWT([]byte(constant.SECRET_JWT))(GetBookingByCartIDControllerTesting())(context)

	res_body := res.Body.String()
	var response CartBookResponse
	er := json.Unmarshal([]byte(res_body), &response)
	if er != nil {
		assert.Error(t, er, "error")
	}
	t.Run("GET /jwt/cart", func(t *testing.T) {
		assert.Equal(t, testCase.ExpectCode, res.Code)
		assert.Equal(t, "Bad Request", response.Message)
	})

}

//====================================================================================================================================

// // Fungsi untuk melakukan testing fungsi ProductRentCheckController
// // kondisi request success
// func TestProductRentCheckControllerrSuccess(t *testing.T) {
// 	var testCase = TestCaseCart{
// 		Name:       "items available",
// 		Path:       "/booking/check/:id",
// 		ExpectCode: http.StatusOK,
// 	}

// 	e := InitEcho()

// 	InsertMockDataCToDB()
// 	config.DB.Save(&mock_data_user_c)
// 	config.DB.Save(&mock_data_product_c)
// 	config.DB.Save(&mock_data_cart)
// 	config.DB.Save(&mock_data_checkoutmethod)
// 	config.DB.Save(&mock_data_methodtype)
// 	config.DB.Save(&mock_data_transaction)
// 	config.DB.Save(&mock_data_booking)

// 	check := BodyDate{
// 		Time_In:  "2022-01-02",
// 		Time_Out: "2022-01-04",
// 	}
// 	fmt.Println("ini test: ", check)

// 	body, err := json.Marshal(check)
// 	if err != nil {
// 		t.Error(t, err, "error")
// 	}
// 	fmt.Println("ini test: ", body)
// 	req := httptest.NewRequest(http.MethodPost, "/booking/check/:id", bytes.NewBuffer(body))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	context.SetParamNames("id")
// 	context.SetParamValues("1")
// 	ProductRentCheckController(context)

// 	res_body := res.Body.String()
// 	var response CartBookResponse
// 	er := json.Unmarshal([]byte(res_body), &response)
// 	if er != nil {
// 		assert.Error(t, er, "error")
// 	}

// 	assert.Equal(t, testCase.ExpectCode, res.Code)
// 	assert.Equal(t, "Item available", response.Message)

// }
