package controllers

import (
	"project3/config"
	"project3/middlewares"
	"project3/models"
	"time"

	"github.com/labstack/echo/v4"
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
	mock_data_province_b = models.Province{
		Province_Name: "Jawa Barat",
	}
	mock_data_city_b = models.City{
		City_Name:  "Subang",
		ProvinceID: 1,
	}
	mock_data_category_b = models.Category{
		Category_Name: "Elektronik",
	}
	mock_data_subcategory_b = models.Subcategory{
		Subcategory_Name: "Kamera",
		CategoryID:       1,
	}
	mock_data_guarantee_b = models.Guarantee{
		Guarantee_Name: "KK",
	}
	mock_data_product_b = models.Products{
		Name:          "Kamera DSLR Canon",
		UsersID:       1,
		SubcategoryID: 1,
		CityID:        1,
		Price:         50000,
		Description:   "Murah yang terbaik",
		Stock:         5,
		Longitude:     107,
		Latitude:      -6,
	}
	mock_data_product2_b = models.Products{
		Name:          "Kamera DSLR Nikon",
		UsersID:       2,
		SubcategoryID: 1,
		CityID:        1,
		Price:         50000,
		Description:   "Murah yang terbaik",
		Stock:         2,
		Longitude:     107,
		Latitude:      -6,
	}
	mock_data_photo_b = models.Photos{
		Photo_Name: "inicontohfoto.jpg",
		Url:        "https://googlecloud/inicontohfoto.jgp",
		ProductsID: 1,
	}
	mock_data_product_guarantee_b = models.ProductsGuarantee{
		GuaranteeID: 1,
		ProductsID:  1,
	}
	mock_data_user_b = models.Users{
		Nama:         "alfy",
		Email:        "alfy@gmail.com",
		Password:     "12345678",
		Phone_Number: "081296620776",
	}
	mock_data_user2_b = models.Users{
		Nama:         "alfy1",
		Email:        "alfy1@gmail.com",
		Password:     "12345678",
		Phone_Number: "081296627876",
	}
	mock_data_login_b = models.Users{
		Email:    "alfy@gmail.com",
		Password: "12345678",
	}
	mock_data_cart = models.Cart{
		UsersID: 1,
	}
	mock_data_cart2 = models.Cart{
		UsersID: 2,
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
	mock_data_booking2 = models.Booking{
		ProductsID:     1,
		CartID:         1,
		Time_In:        time1,
		Time_Out:       time2,
		Total_Day:      2,
		Qty:            1,
		Total:          100000,
		Status_Payment: "succes",
	}
	mock_data_booking3 = models.Booking{
		ProductsID:     1,
		CartID:         2,
		Time_In:        time1,
		Time_Out:       time2,
		Total_Day:      2,
		Qty:            1,
		Total:          100000,
		Status_Payment: "succes",
	}
	mock_data_transaction = models.Transaction{
		Total: 100000,
	}
)

// Fungsi untuk memasukkan data user test ke dalam database
func InsertMockDataUsersBToDB() error {
	query := config.DB.Save(&mock_data_user_b)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

// Fungsi untuk memasukkan data user test ke dalam database
func InsertMockDataBToDB() {
	config.DB.Save(&mock_data_category_b)
	config.DB.Save(&mock_data_subcategory_b)
	config.DB.Save(&mock_data_province_b)
	config.DB.Save(&mock_data_city_b)
	config.DB.Save(&mock_data_guarantee_b)
}

// inisialisasi echo
func InitEchoB() *echo.Echo {
	config.InitDBTest()
	e := echo.New()

	return e
}

// Fungsi untuk melakukan login dan ekstraksi token JWT
func UsingJWTB() (string, error) {
	// Melakukan login data user test
	InsertMockDataUsersBToDB()
	var user models.Users
	tx := config.DB.Where("email = ? AND password = ?", mock_data_login_b.Email, mock_data_login_b.Password).First(&user)
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

// // Fungsi testing CreateProductController
// func GetBookingByIDControllerTesting() echo.HandlerFunc {
// 	return GetBookingByIdController
// }

// // Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// // kondisi request success
// func TestGetBookingByIdControllerSuccess(t *testing.T) {
// 	var testCase = TestCaseCart{
// 		Name:       "success to get booking by id",
// 		Path:       "/booking",
// 		ExpectCode: http.StatusOK,
// 	}

// 	e := InitEchoB()
// 	// Mendapatkan token
// 	token, err := UsingJWTB()
// 	if err != nil {
// 		panic(err)
// 	}

// 	InsertMockDataBToDB()
// 	config.DB.Save(&mock_data_user2_b)
// 	config.DB.Save(&mock_data_product_b)
// 	config.DB.Save(&mock_data_product2_b)
// 	config.DB.Save(&mock_data_cart)
// 	config.DB.Save(&mock_data_booking)

// 	req := httptest.NewRequest(http.MethodGet, "/booking", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	context.SetParamNames("id")
// 	context.SetParamValues("1")
// 	middleware.JWT([]byte(constant.SECRET_JWT))(GetBookingByIDControllerTesting())(context)

// 	res_body := res.Body.String()
// 	var response CartBookResponse
// 	er := json.Unmarshal([]byte(res_body), &response)
// 	if er != nil {
// 		assert.Error(t, er, "error")
// 	}
// 	t.Run("GET /jwt/booking", func(t *testing.T) {
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "Successful Operation", response.Message)
// 	})

// }

// // Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// // kondisi request failed
// func TestGetBookingByIDControllerNotFound(t *testing.T) {
// 	var testCase = TestCaseCart{
// 		Name:       "get booking not found",
// 		Path:       "/booking",
// 		ExpectCode: http.StatusBadRequest,
// 	}

// 	e := InitEchoB()
// 	// Mendapatkan token
// 	token, err := UsingJWTB()
// 	if err != nil {
// 		panic(err)
// 	}

// 	req := httptest.NewRequest(http.MethodGet, "/booking", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	context.SetParamNames("id")
// 	context.SetParamValues("1")
// 	middleware.JWT([]byte(constant.SECRET_JWT))(GetBookingByIDControllerTesting())(context)

// 	res_body := res.Body.String()
// 	var response CartBookResponse
// 	er := json.Unmarshal([]byte(res_body), &response)
// 	if er != nil {
// 		assert.Error(t, er, "error")
// 	}
// 	t.Run("GET /jwt/booking", func(t *testing.T) {
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "Booking not found", response.Message)
// 	})

// }

// // Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// // kondisi request failed
// func TestGetBookingByIDControllerFailed(t *testing.T) {
// 	var testCase = TestCaseCart{
// 		Name:       "failed to get booking",
// 		Path:       "/booking",
// 		ExpectCode: http.StatusBadRequest,
// 	}

// 	e := InitEchoB()
// 	// Mendapatkan token
// 	token, err := UsingJWTB()
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Melakukan penghapusan tabel untuk membuat request failed
// 	config.DB.Migrator().DropTable(&models.Booking{})

// 	req := httptest.NewRequest(http.MethodGet, "/booking", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	context.SetParamNames("id")
// 	context.SetParamValues("1")
// 	middleware.JWT([]byte(constant.SECRET_JWT))(GetBookingByIDControllerTesting())(context)

// 	res_body := res.Body.String()
// 	var response CartBookResponse
// 	er := json.Unmarshal([]byte(res_body), &response)
// 	if er != nil {
// 		assert.Error(t, er, "error")
// 	}
// 	t.Run("GET /jwt/booking", func(t *testing.T) {
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "Bad Request", response.Message)
// 	})

// }

// // Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// // kondisi request failed
// func TestGetBookingByIDControlleFailed(t *testing.T) {
// 	var testCase = TestCaseCart{
// 		Name:       "failed to get booking",
// 		Path:       "/booking",
// 		ExpectCode: http.StatusBadRequest,
// 	}

// 	e := InitEchoB()
// 	// Mendapatkan token
// 	token, err := UsingJWTB()
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Melakukan penghapusan tabel untuk membuat request failed
// 	config.DB.Migrator().DropTable(&models.Booking{})

// 	req := httptest.NewRequest(http.MethodGet, "/booking", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	context.SetParamNames("id")
// 	context.SetParamValues("!")
// 	middleware.JWT([]byte(constant.SECRET_JWT))(GetBookingByIDControllerTesting())(context)

// 	res_body := res.Body.String()
// 	var response CartBookResponse
// 	er := json.Unmarshal([]byte(res_body), &response)
// 	if er != nil {
// 		assert.Error(t, er, "error")
// 	}
// 	t.Run("GET /jwt/booking", func(t *testing.T) {
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "False Param", response.Message)
// 	})

// }

//===============================================================================================================================

// // Fungsi testing CreateProductController
// func GetBookingByCartIDControllerTesting() echo.HandlerFunc {
// 	return GetBookingByCartIDController
// }

// // Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// // kondisi request success
// func TestGetBookingByCartIDControllerSuccess(t *testing.T) {
// 	var testCase = TestCaseCart{
// 		Name:       "success to get booking by cart id",
// 		Path:       "/cart",
// 		ExpectCode: http.StatusOK,
// 	}

// 	e := InitEchoB()
// 	// Mendapatkan token
// 	token, err := UsingJWTB()
// 	if err != nil {
// 		panic(err)
// 	}

// 	InsertMockDataBToDB()
// 	config.DB.Save(&mock_data_user2_b)
// 	config.DB.Save(&mock_data_product_b)
// 	config.DB.Save(&mock_data_product2_b)
// 	config.DB.Save(&mock_data_cart)
// 	config.DB.Save(&mock_data_booking)

// 	req := httptest.NewRequest(http.MethodGet, "/cart", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	middleware.JWT([]byte(constant.SECRET_JWT))(GetBookingByCartIDControllerTesting())(context)

// 	res_body := res.Body.String()
// 	var response CartBookResponse
// 	er := json.Unmarshal([]byte(res_body), &response)
// 	if er != nil {
// 		assert.Error(t, er, "error")
// 	}
// 	t.Run("GET /jwt/cart", func(t *testing.T) {
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "Successful Operation", response.Message)
// 	})

// }

// // Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// // kondisi request failed
// func TestGetBookingByCartIDControllerNotFound(t *testing.T) {
// 	var testCase = TestCaseCart{
// 		Name:       "get booking not found",
// 		Path:       "/cart",
// 		ExpectCode: http.StatusBadRequest,
// 	}

// 	e := InitEchoB()
// 	// Mendapatkan token
// 	token, err := UsingJWTB()
// 	if err != nil {
// 		panic(err)
// 	}

// 	req := httptest.NewRequest(http.MethodGet, "/cart", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	middleware.JWT([]byte(constant.SECRET_JWT))(GetBookingByCartIDControllerTesting())(context)

// 	res_body := res.Body.String()
// 	var response CartBookResponse
// 	er := json.Unmarshal([]byte(res_body), &response)
// 	if er != nil {
// 		assert.Error(t, er, "error")
// 	}
// 	t.Run("GET /jwt/cart", func(t *testing.T) {
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "Booking not found", response.Message)
// 	})

// }

// // Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// // kondisi request failed
// func TestGetBookingByCartIDControllerFailed(t *testing.T) {
// 	var testCase = TestCaseCart{
// 		Name:       "failed to get booking",
// 		Path:       "/cart",
// 		ExpectCode: http.StatusBadRequest,
// 	}

// 	e := InitEchoB()
// 	// Mendapatkan token
// 	token, err := UsingJWTB()
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Melakukan penghapusan tabel untuk membuat request failed
// 	config.DB.Migrator().DropTable(&models.Booking{})

// 	req := httptest.NewRequest(http.MethodGet, "/cart", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	middleware.JWT([]byte(constant.SECRET_JWT))(GetBookingByCartIDControllerTesting())(context)

// 	res_body := res.Body.String()
// 	var response CartBookResponse
// 	er := json.Unmarshal([]byte(res_body), &response)
// 	if er != nil {
// 		assert.Error(t, er, "error")
// 	}
// 	t.Run("GET /jwt/cart", func(t *testing.T) {
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "Bad Request", response.Message)
// 	})

// }

// // Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// // kondisi request failed
// func TestGetBookingByCartIDControllerCartNotFound(t *testing.T) {
// 	var testCase = TestCaseCart{
// 		Name:       "failed to get booking",
// 		Path:       "/cart",
// 		ExpectCode: http.StatusBadRequest,
// 	}

// 	e := InitEchoB()
// 	// Mendapatkan token
// 	token, err := UsingJWTB()
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Melakukan penghapusan tabel untuk membuat request failed
// 	config.DB.Migrator().DropTable(&models.Cart{})

// 	req := httptest.NewRequest(http.MethodGet, "/cart", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	middleware.JWT([]byte(constant.SECRET_JWT))(GetBookingByCartIDControllerTesting())(context)

// 	res_body := res.Body.String()
// 	var response CartBookResponse
// 	er := json.Unmarshal([]byte(res_body), &response)
// 	if er != nil {
// 		assert.Error(t, er, "error")
// 	}
// 	t.Run("GET /jwt/cart", func(t *testing.T) {
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "Bad Request", response.Message)
// 	})

// }

// //====================================================================================================================================
// // Fungsi testing CreateProductController
// func GetHistoryByCartIDControllerTesting() echo.HandlerFunc {
// 	return GetHistoryByCartIDController
// }

// // Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// // kondisi request success
// func TestGetHistoryByCartIDControllerSuccess(t *testing.T) {
// 	var testCase = TestCaseCart{
// 		Name:       "success to get history by cart id",
// 		Path:       "/history",
// 		ExpectCode: http.StatusOK,
// 	}

// 	e := InitEchoB()
// 	// Mendapatkan token
// 	token, err := UsingJWTB()
// 	if err != nil {
// 		panic(err)
// 	}

// 	InsertMockDataBToDB()
// 	config.DB.Save(&mock_data_user2_b)
// 	config.DB.Save(&mock_data_product_b)
// 	config.DB.Save(&mock_data_product2_b)
// 	config.DB.Save(&mock_data_cart)
// 	config.DB.Save(&mock_data_booking2)

// 	req := httptest.NewRequest(http.MethodGet, "/history", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	middleware.JWT([]byte(constant.SECRET_JWT))(GetHistoryByCartIDControllerTesting())(context)

// 	res_body := res.Body.String()
// 	var response CartBookResponse
// 	er := json.Unmarshal([]byte(res_body), &response)
// 	if er != nil {
// 		assert.Error(t, er, "error")
// 	}
// 	t.Run("GET /jwt/history", func(t *testing.T) {
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "Successful Operation", response.Message)
// 	})

// }

// // Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// // kondisi request failed
// func TestGetHistoryByCartIDControllerNotFound(t *testing.T) {
// 	var testCase = TestCaseCart{
// 		Name:       "get history not found",
// 		Path:       "/history",
// 		ExpectCode: http.StatusBadRequest,
// 	}

// 	e := InitEchoB()
// 	// Mendapatkan token
// 	token, err := UsingJWTB()
// 	if err != nil {
// 		panic(err)
// 	}

// 	req := httptest.NewRequest(http.MethodGet, "/history", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	middleware.JWT([]byte(constant.SECRET_JWT))(GetHistoryByCartIDControllerTesting())(context)

// 	res_body := res.Body.String()
// 	var response CartBookResponse
// 	er := json.Unmarshal([]byte(res_body), &response)
// 	if er != nil {
// 		assert.Error(t, er, "error")
// 	}
// 	t.Run("GET /jwt/history", func(t *testing.T) {
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "Booking not found", response.Message)
// 	})

// }

// // Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// // kondisi request failed
// func TestGetBookingByIDControllerFailed(t *testing.T) {
// 	var testCase = TestCaseCart{
// 		Name:       "failed to get history",
// 		Path:       "/history",
// 		ExpectCode: http.StatusBadRequest,
// 	}

// 	e := InitEchoB()
// 	// Mendapatkan token
// 	token, err := UsingJWTB()
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Melakukan penghapusan tabel untuk membuat request failed
// 	config.DB.Migrator().DropTable(&models.Booking{})

// 	req := httptest.NewRequest(http.MethodGet, "/history", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	middleware.JWT([]byte(constant.SECRET_JWT))(GetHistoryByCartIDControllerTesting())(context)

// 	res_body := res.Body.String()
// 	var response CartBookResponse
// 	er := json.Unmarshal([]byte(res_body), &response)
// 	if er != nil {
// 		assert.Error(t, er, "error")
// 	}
// 	t.Run("GET /jwt/history", func(t *testing.T) {
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "Bad Request", response.Message)
// 	})

// }

// // Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// // kondisi request failed
// func TestGetHistoryByCartIDControllerNoCart(t *testing.T) {
// 	var testCase = TestCaseCart{
// 		Name:       "failed to get history",
// 		Path:       "/history",
// 		ExpectCode: http.StatusBadRequest,
// 	}

// 	e := InitEchoB()
// 	// Mendapatkan token
// 	token, err := UsingJWTB()
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Melakukan penghapusan tabel untuk membuat request failed
// 	config.DB.Migrator().DropTable(&models.Cart{})

// 	req := httptest.NewRequest(http.MethodGet, "/history", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	middleware.JWT([]byte(constant.SECRET_JWT))(GetHistoryByCartIDControllerTesting())(context)

// 	res_body := res.Body.String()
// 	var response CartBookResponse
// 	er := json.Unmarshal([]byte(res_body), &response)
// 	if er != nil {
// 		assert.Error(t, er, "error")
// 	}
// 	t.Run("GET /jwt/history", func(t *testing.T) {
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "Bad Request", response.Message)
// 	})

// }

//====================================================================================================================================
// // Fungsi testing DeleteProductController
// func CancelBookingControllerTesting() echo.HandlerFunc {
// 	return CancelBookingController
// }

// // Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// // kondisi request failed
// func TestCancelBookingControllerSuccess(t *testing.T) {
// 	var testCase = TestCaseCart{
// 		Name:       "Success cancel booking",
// 		Path:       "/booking/:id",
// 		ExpectCode: http.StatusOK,
// 	}

// 	e := InitEchoB()
// 	// Mendapatkan token
// 	token, err := UsingJWTB()
// 	if err != nil {
// 		panic(err)
// 	}

// 	InsertMockDataBToDB()
// 	config.DB.Save(&mock_data_product_b)
// 	config.DB.Save(&mock_data_cart)
// 	config.DB.Save(&mock_data_booking)

// 	req := httptest.NewRequest(http.MethodDelete, "/booking/:id", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	context.SetParamNames("id")
// 	context.SetParamValues("1")
// 	middleware.JWT([]byte(constant.SECRET_JWT))(CancelBookingControllerTesting())(context)

// 	res_body := res.Body.String()
// 	var response CartBookResponse
// 	er := json.Unmarshal([]byte(res_body), &response)
// 	if er != nil {
// 		assert.Error(t, er, "error")
// 	}
// 	t.Run("DELETE /jwt/booking/:id", func(t *testing.T) {
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "Successful Operation", response.Message)
// 	})

// }

// // Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// // kondisi request failed
// func TestCancelBookingControllerFailed(t *testing.T) {
// 	var testCase = TestCaseCart{
// 		Name:       "Failed to cancel booking",
// 		Path:       "/booking/:id",
// 		ExpectCode: http.StatusBadRequest,
// 	}

// 	e := InitEchoB()
// 	// Mendapatkan token
// 	token, err := UsingJWTB()
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Melakukan penghapusan tabel untuk membuat request failed
// 	config.DB.Migrator().DropTable(&models.Booking{})

// 	req := httptest.NewRequest(http.MethodDelete, "/booking/:id", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	context.SetParamNames("id")
// 	context.SetParamValues("1")
// 	middleware.JWT([]byte(constant.SECRET_JWT))(CancelBookingControllerTesting())(context)

// 	res_body := res.Body.String()
// 	var response CartBookResponse
// 	er := json.Unmarshal([]byte(res_body), &response)
// 	if er != nil {
// 		assert.Error(t, er, "error")
// 	}
// 	t.Run("DELETE /jwt/booking/:id", func(t *testing.T) {
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "Bad Request", response.Message)
// 	})

// }

// // Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// // kondisi request failed
// func TestDeleteProductsControllerFalseParam(t *testing.T) {
// 	var testCase = TestCaseCart{
// 		Name:       "False param",
// 		Path:       "/booking/:id",
// 		ExpectCode: http.StatusBadRequest,
// 	}

// 	e := InitEchoB()
// 	// Mendapatkan token
// 	token, err := UsingJWTB()
// 	if err != nil {
// 		panic(err)
// 	}

// 	req := httptest.NewRequest(http.MethodDelete, "/booking/:id", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	context.SetParamNames("id")
// 	context.SetParamValues("!")
// 	middleware.JWT([]byte(constant.SECRET_JWT))(CancelBookingControllerTesting())(context)

// 	res_body := res.Body.String()
// 	var response CartBookResponse
// 	er := json.Unmarshal([]byte(res_body), &response)
// 	if er != nil {
// 		assert.Error(t, er, "error")
// 	}
// 	t.Run("DELETE /jwt/booking/:id", func(t *testing.T) {
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "False Param", response.Message)
// 	})

// }

// // Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// // kondisi request failed
// func TestDeleteProductsControllerNotAllowed(t *testing.T) {
// 	var testCase = TestCaseCart{
// 		Name:       "Not allowed to delete",
// 		Path:       "/booking/:id",
// 		ExpectCode: http.StatusBadRequest,
// 	}

// 	e := InitEchoB()
// 	// Mendapatkan token
// 	token, err := UsingJWTB()
// 	if err != nil {
// 		panic(err)
// 	}

// 	InsertMockDataBToDB()
// 	config.DB.Save(&mock_data_cart)
// 	config.DB.Save(&mock_data_user2_b)
// 	config.DB.Save(&mock_data_cart2)
// 	config.DB.Save(&mock_data_product_b)
// 	config.DB.Save(&mock_data_booking)
// 	config.DB.Save(&mock_data_booking3)

// 	req := httptest.NewRequest(http.MethodDelete, "/booking/:id", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	context.SetParamNames("id")
// 	context.SetParamValues("2")
// 	middleware.JWT([]byte(constant.SECRET_JWT))(CancelBookingControllerTesting())(context)

// 	res_body := res.Body.String()
// 	var response CartBookResponse
// 	er := json.Unmarshal([]byte(res_body), &response)
// 	if er != nil {
// 		assert.Error(t, er, "error")
// 	}
// 	t.Run("DELETE /jwt/booking/:id", func(t *testing.T) {
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "Access Forbidden", response.Message)
// 	})

// }

// //====================================================================================================================================

// // Fungsi untuk melakukan testing fungsi ProductRentCheckController
// // kondisi request success
// func TestProductRentCheckControllerrSuccess(t *testing.T) {
// 	var testCase = TestCaseCart{
// 		Name:       "items available",
// 		Path:       "/booking/check/:id",
// 		ExpectCode: http.StatusOK,
// 	}

// 	e := InitEchoB()

// 	InsertMockDataBToDB()
// 	config.DB.Save(&mock_data_product_b)

// 	check := BodyDate{
// 		Time_In:  "2022-01-02",
// 		Time_Out: "2022-01-04",
// 	}
// 	fmt.Println("ini test: ", check)

// 	body, err := json.Marshal(check)
// 	if err != nil {
// 		t.Error(t, err, "error")
// 	}

// 	req := httptest.NewRequest(http.MethodPost, "/booking/check/:id", bytes.NewBuffer(body))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	context.SetParamNames("id")
// 	context.SetParamValues("1")
// 	ProductRentCheckController(context)

// 	res_body := res.Body.String()
// 	fmt.Println("res_body")
// 	var response CartBookResponse
// 	er := json.Unmarshal([]byte(res_body), &response)
// 	if er != nil {
// 		assert.Error(t, er, "error")
// 	}

// 	assert.Equal(t, testCase.ExpectCode, res.Code)
// 	assert.Equal(t, "Item available", response.Message)

// }

// // Fungsi untuk melakukan testing fungsi ProductRentCheckController
// // kondisi request success
// func TestProductRentCheckControllerFailed(t *testing.T) {
// 	var testCase = TestCaseCart{
// 		Name:       "booking empty",
// 		Path:       "/booking/check/:id",
// 		ExpectCode: http.StatusBadRequest,
// 	}

// 	e := InitEchoB()

// 	// Melakukan penghapusan tabel untuk membuat request failed
// 	config.DB.Migrator().DropTable(&models.Booking{})

// 	check := BodyDate{
// 		Time_In:  "2022-01-02",
// 		Time_Out: "2022-01-04",
// 	}
// 	fmt.Println("ini test: ", check)

// 	body, err := json.Marshal(check)
// 	if err != nil {
// 		t.Error(t, err, "error")
// 	}

// 	req := httptest.NewRequest(http.MethodPost, "/booking/check/:id", bytes.NewBuffer(body))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	context.SetParamNames("id")
// 	context.SetParamValues("1")
// 	ProductRentCheckController(context)

// 	res_body := res.Body.String()
// 	fmt.Println("res_body")
// 	var response CartBookResponse
// 	er := json.Unmarshal([]byte(res_body), &response)
// 	if er != nil {
// 		assert.Error(t, er, "error")
// 	}

// 	assert.Equal(t, testCase.ExpectCode, res.Code)
// 	assert.Equal(t, "Bad Request", response.Message)

// }

// // Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// // kondisi request failed
// func TestProductRentCheckControllerFalseParam(t *testing.T) {
// 	var testCase = TestCaseCart{
// 		Name:       "False param",
// 		Path:       "/booking/check/:id",
// 		ExpectCode: http.StatusBadRequest,
// 	}

// 	e := InitEchoB()

// 	check := BodyDate{
// 		Time_In:  "2022-01-02",
// 		Time_Out: "2022-01-04",
// 	}
// 	fmt.Println("ini test: ", check)

// 	body, err := json.Marshal(check)
// 	if err != nil {
// 		t.Error(t, err, "error")
// 	}

// 	req := httptest.NewRequest(http.MethodDelete, "/booking/check/:id", body)
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	context.SetParamNames("id")
// 	context.SetParamValues("!")
// 	ProductRentCheckController()(context)

// 	res_body := res.Body.String()
// 	var response CartBookResponse
// 	er := json.Unmarshal([]byte(res_body), &response)
// 	if er != nil {
// 		assert.Error(t, er, "error")
// 	}
// 	t.Run("DELETE /jwt/booking/:id", func(t *testing.T) {
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "False Param", response.Message)
// 	})

// }
