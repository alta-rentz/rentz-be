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

// Struct yang digunakan ketika test request success, dapat menampung banyak data
type ProductsResponse struct {
	Code    string
	Message string
}

// Struct untuk menampung data test case
type TestCase struct {
	Name       string
	Path       string
	ExpectCode int
}

type Login struct {
	Email    string
	Password string
}

// data dummy
var (
	mock_data_province = models.Province{
		Province_Name: "Jawa Barat",
	}
	mock_data_city = models.City{
		City_Name:  "Subang",
		ProvinceID: 1,
	}
	mock_data_category = models.Category{
		Category_Name: "Elektronik",
	}
	mock_data_subcategory = models.Subcategory{
		Subcategory_Name: "Kamera",
		CategoryID:       1,
	}
	mock_data_guarantee = models.Guarantee{
		Guarantee_Name: "KK",
	}
	mock_data_product = models.Products{
		Name:          "Kamera DSLR Canon",
		UsersID:       1,
		SubcategoryID: 1,
		CityID:        1,
		Price:         50000,
		Description:   "Murah yang terbaik",
		Stock:         5,
	}
	mock_data_product2 = models.Products{
		Name:          "Kamera DSLR Nikon",
		UsersID:       2,
		SubcategoryID: 1,
		CityID:        1,
		Price:         50000,
		Description:   "Murah yang terbaik",
		Stock:         2,
	}
	mock_data_photo = models.Photos{
		Photo_Name: "inicontohfoto.jpg",
		Url:        "https://googlecloud/inicontohfoto.jgp",
		ProductsID: 1,
	}
	mock_data_product_guarantee = models.ProductsGuarantee{
		GuaranteeID: 1,
		ProductsID:  1,
	}
	mock_data_user2 = models.Users{
		Nama:         "alfy1",
		Email:        "alfy1@gmail.com",
		Password:     "12345678",
		Phone_Number: "081296627876",
	}
	mock_data_user = models.Users{
		Nama:         "alfy",
		Email:        "alfy@gmail.com",
		Password:     "12345678",
		Phone_Number: "081296620776",
	}
	mock_data_login = models.Users{
		Email:    "alfy@gmail.com",
		Password: "12345678",
	}
)

// inisialisasi echo
func InitEcho() *echo.Echo {
	config.InitDBTest()
	e := echo.New()

	return e
}

// Fungsi untuk memasukkan data user test ke dalam database
func InsertMockDataUsersToDB() error {
	query := config.DB.Save(&mock_data_user)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

// Fungsi untuk memasukkan data user test ke dalam database
func InsertMockDataToDB() {
	config.DB.Save(&mock_data_category)
	config.DB.Save(&mock_data_subcategory)
	config.DB.Save(&mock_data_province)
	config.DB.Save(&mock_data_city)
	config.DB.Save(&mock_data_guarantee)
}

// Fungsi untuk melakukan login dan ekstraksi token JWT
func UsingJWT() (string, error) {
	// Melakukan login data user test
	InsertMockDataUsersToDB()
	var user models.Users
	tx := config.DB.Where("email = ? AND password = ?", mock_data_login.Email, mock_data_login.Password).First(&user)
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

// Fungsi untuk melakukan testing fungsi GetAllProductsController
// kondisi request success
func TestGetAllProductsControllerSuccess(t *testing.T) {
	var testCases = []struct {
		Name       string
		Path       string
		ExpectCode int
		ExpectSize int
	}{
		{
			Name:       "success to get all data products",
			Path:       "/products",
			ExpectCode: http.StatusOK,
			ExpectSize: 1,
		},
	}

	e := InitEcho()

	InsertMockDataToDB()
	config.DB.Save(&mock_data_user)
	config.DB.Save(&mock_data_product)
	config.DB.Save(&mock_data_product_guarantee)
	config.DB.Save(&mock_data_photo)

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)

	for _, testCase := range testCases {
		context.SetPath(testCase.Path)

		if assert.NoError(t, GetAllProductsController(context)) {
			res_body := res.Body.String()
			var response ProductsResponse
			er := json.Unmarshal([]byte(res_body), &response)
			if er != nil {
				assert.Error(t, er, "error")
			}
			assert.Equal(t, testCase.ExpectCode, res.Code)
			assert.Equal(t, "Successful Operation", response.Message)
		}
	}
}

// // Fungsi untuk melakukan testing fungsi GetProductController
// // kondisi request failed
// func TestGetAllProductsControllerFailed(t *testing.T) {
// 	var testCase = TestCase{
// 		Name:       "failed to get products",
// 		Path:       "/products",
// 		ExpectCode: http.StatusBadRequest,
// 	}

// 	e := InitEcho()

// 	// Melakukan penghapusan tabel untuk membuat request failed
// 	config.DB.Migrator().DropTable(&models.Products{})

// 	req := httptest.NewRequest(http.MethodGet, "/products", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)

// 	if assert.NoError(t, GetAllProductsController(context)) {
// 		res_body := res.Body.String()
// 		var response ProductsResponse
// 		er := json.Unmarshal([]byte(res_body), &response)
// 		if er != nil {
// 			assert.Error(t, er, "error")
// 		}
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "Bad Request", response.Message)
// 	}
// }

// // Fungsi untuk melakukan testing fungsi GetProductController
// // kondisi request success
// func TestGetAllProductsControllerNotFound(t *testing.T) {
// 	var testCase = TestCase{
// 		Name:       "products not found",
// 		Path:       "/products",
// 		ExpectCode: http.StatusBadRequest,
// 	}

// 	e := InitEcho()

// 	req := httptest.NewRequest(http.MethodGet, "/products", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)

// 	if assert.NoError(t, GetAllProductsController(context)) {
// 		res_body := res.Body.String()
// 		var response ProductsResponse
// 		er := json.Unmarshal([]byte(res_body), &response)
// 		if er != nil {
// 			assert.Error(t, er, "error")
// 		}
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "Items not found", response.Message)
// 	}
// }

// // Fungsi untuk melakukan testing fungsi GetProductByIDControllers
// // kondisi request success
// func TestGetProductsByIDControllerSuccess(t *testing.T) {
// 	var testCase = TestCase{
// 		Name:       "success to get one data product by id",
// 		Path:       "/products/:id",
// 		ExpectCode: http.StatusOK,
// 	}

// 	e := InitEcho()

// 	InsertMockDataToDB()
// 	config.DB.Save(&mock_data_user)
// 	config.DB.Save(&mock_data_product)
// 	config.DB.Save(&mock_data_product_guarantee)
// 	config.DB.Save(&mock_data_photo)

// 	req := httptest.NewRequest(http.MethodGet, "/products/:id", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	context.SetParamNames("id")
// 	context.SetParamValues("1")

// 	if assert.NoError(t, GetProductByIDController(context)) {
// 		res_body := res.Body.String()
// 		var response ProductsResponse
// 		er := json.Unmarshal([]byte(res_body), &response)
// 		if er != nil {
// 			assert.Error(t, er, "error")
// 		}
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "Successful Operation", response.Message)
// 	}

// }

// // Fungsi untuk melakukan testing fungsi GetProductByIDControllers
// // kondisi request failed
// func TestGetProductsByIDControllerFalseParam(t *testing.T) {
// 	var testCase = TestCase{
// 		Name:       "false param to get one data product by id",
// 		Path:       "/products/:id",
// 		ExpectCode: http.StatusBadRequest,
// 	}

// 	e := InitEcho()

// 	InsertMockDataToDB()
// 	config.DB.Save(&mock_data_user)
// 	config.DB.Save(&mock_data_product)
// 	config.DB.Save(&mock_data_product_guarantee)
// 	config.DB.Save(&mock_data_photo)

// 	req := httptest.NewRequest(http.MethodGet, "/products/:id", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	context.SetParamNames("id")
// 	context.SetParamValues("!")

// 	if assert.NoError(t, GetProductByIDController(context)) {
// 		res_body := res.Body.String()
// 		var response ProductsResponse
// 		er := json.Unmarshal([]byte(res_body), &response)
// 		if er != nil {
// 			assert.Error(t, er, "error")
// 		}
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "False Param", response.Message)
// 	}

// }

// // Fungsi untuk melakukan testing fungsi GetProductByIDControllers
// // kondisi request failed
// func TestGetProductsByIDControllerWrongID(t *testing.T) {
// 	var testCase = TestCase{
// 		Name:       "product by id not found",
// 		Path:       "/products/:id",
// 		ExpectCode: http.StatusBadRequest,
// 	}

// 	e := InitEcho()

// 	InsertMockDataToDB()
// 	config.DB.Save(&mock_data_user)
// 	config.DB.Save(&mock_data_product)
// 	config.DB.Save(&mock_data_product_guarantee)
// 	config.DB.Save(&mock_data_photo)

// 	req := httptest.NewRequest(http.MethodGet, "/products/:id", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	context.SetParamNames("id")
// 	context.SetParamValues("2")

// 	if assert.NoError(t, GetProductByIDController(context)) {
// 		res_body := res.Body.String()
// 		var response ProductsResponse
// 		er := json.Unmarshal([]byte(res_body), &response)
// 		if er != nil {
// 			assert.Error(t, er, "error")
// 		}
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "Items not found", response.Message)
// 	}

// }

// // Fungsi untuk melakukan testing fungsi GetProductByIDControllers
// // kondisi request failed
// func TestGetProductsByIDControllerFailed(t *testing.T) {
// 	var testCase = TestCase{
// 		Name:       "failed to get product",
// 		Path:       "/products/:id",
// 		ExpectCode: http.StatusBadRequest,
// 	}

// 	e := InitEcho()

// 	// Melakukan penghapusan tabel untuk membuat request failed
// 	config.DB.Migrator().DropTable(&models.Products{})

// 	req := httptest.NewRequest(http.MethodGet, "/products/:id", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	context.SetParamNames("id")
// 	context.SetParamValues("1")

// 	if assert.NoError(t, GetProductByIDController(context)) {
// 		res_body := res.Body.String()
// 		var response ProductsResponse
// 		er := json.Unmarshal([]byte(res_body), &response)
// 		if er != nil {
// 			assert.Error(t, er, "error")
// 		}
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "Bad Request", response.Message)
// 	}
// }

// // Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// // kondisi request success
// func TestGetProductsBySubcategoryIDControllerSuccess(t *testing.T) {
// 	var testCase = TestCase{
// 		Name:       "success to get one data product by subcategory id",
// 		Path:       "/products/subcategory/:id",
// 		ExpectCode: http.StatusOK,
// 	}

// 	e := InitEcho()

// 	InsertMockDataToDB()
// 	config.DB.Save(&mock_data_user)
// 	config.DB.Save(&mock_data_product)
// 	config.DB.Save(&mock_data_product_guarantee)
// 	config.DB.Save(&mock_data_photo)

// 	req := httptest.NewRequest(http.MethodGet, "/products/subcategory/:id", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	context.SetParamNames("id")
// 	context.SetParamValues("1")

// 	if assert.NoError(t, GetProductsBySubcategoryIDController(context)) {
// 		res_body := res.Body.String()
// 		var response ProductsResponse
// 		er := json.Unmarshal([]byte(res_body), &response)
// 		if er != nil {
// 			assert.Error(t, er, "error")
// 		}
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "Successful Operation", response.Message)
// 	}

// }

// // Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// // kondisi request success
// func TestGetProductsBySubcategoryIDControllerFalseParam(t *testing.T) {
// 	var testCase = TestCase{
// 		Name:       "false param to get one data product by subcategory id",
// 		Path:       "/products/subcategory/:id",
// 		ExpectCode: http.StatusBadRequest,
// 	}

// 	e := InitEcho()

// 	InsertMockDataToDB()
// 	config.DB.Save(&mock_data_user)
// 	config.DB.Save(&mock_data_product)
// 	config.DB.Save(&mock_data_product_guarantee)
// 	config.DB.Save(&mock_data_photo)

// 	req := httptest.NewRequest(http.MethodGet, "/products/subcategory/:id", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	context.SetParamNames("id")
// 	context.SetParamValues("!")

// 	if assert.NoError(t, GetProductsBySubcategoryIDController(context)) {
// 		res_body := res.Body.String()
// 		var response ProductsResponse
// 		er := json.Unmarshal([]byte(res_body), &response)
// 		if er != nil {
// 			assert.Error(t, er, "error")
// 		}
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "False Param", response.Message)
// 	}

// }

// // Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// // kondisi request failed
// func TestGetProductsBySubcategoryIDControllerWrongID(t *testing.T) {
// 	var testCase = TestCase{
// 		Name:       "product by subcategory id not found",
// 		Path:       "/products/subcategory/:id",
// 		ExpectCode: http.StatusBadRequest,
// 	}

// 	e := InitEcho()

// 	InsertMockDataToDB()
// 	config.DB.Save(&mock_data_user)
// 	config.DB.Save(&mock_data_product)
// 	config.DB.Save(&mock_data_product_guarantee)
// 	config.DB.Save(&mock_data_photo)

// 	req := httptest.NewRequest(http.MethodGet, "/products/subcategory/:id", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	context.SetParamNames("id")
// 	context.SetParamValues("2")

// 	if assert.NoError(t, GetProductsBySubcategoryIDController(context)) {
// 		res_body := res.Body.String()
// 		var response ProductsResponse
// 		er := json.Unmarshal([]byte(res_body), &response)
// 		if er != nil {
// 			assert.Error(t, er, "error")
// 		}
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "Items not found", response.Message)
// 	}

// }

// // Fungsi untuk melakukan testing fungsiGetProductsBySubcategoryIDControllers
// // kondisi request failed
// func TestGetProductsBySubcategoryIDControllerFailed(t *testing.T) {
// 	var testCase = TestCase{
// 		Name:       "failed to get product by subcategory id",
// 		Path:       "/products/subcateogry/:id",
// 		ExpectCode: http.StatusBadRequest,
// 	}

// 	e := InitEcho()

// 	// Melakukan penghapusan tabel untuk membuat request failed
// 	config.DB.Migrator().DropTable(&models.Products{})

// 	req := httptest.NewRequest(http.MethodGet, "/products/subcategory/:id", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	context.SetParamNames("id")
// 	context.SetParamValues("1")

// 	if assert.NoError(t, GetProductsBySubcategoryIDController(context)) {
// 		res_body := res.Body.String()
// 		var response ProductsResponse
// 		er := json.Unmarshal([]byte(res_body), &response)
// 		if er != nil {
// 			assert.Error(t, er, "error")
// 		}
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "Bad Request", response.Message)
// 	}
// }

// // Fungsi testing CreateProductController
// func GetProductsByUserIDControllerTesting() echo.HandlerFunc {
// 	return GetProductsByUserIDController
// }

// // Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// // kondisi request success
// func TestGetProductsByUserIDControllerSuccess(t *testing.T) {
// 	var testCase = TestCase{
// 		Name:       "success to get one data product by user id",
// 		Path:       "/products",
// 		ExpectCode: http.StatusOK,
// 	}

// 	e := InitEcho()
// 	// Mendapatkan token
// 	token, err := UsingJWT()
// 	if err != nil {
// 		panic(err)
// 	}

// 	InsertMockDataToDB()
// 	config.DB.Save(&mock_data_user)
// 	config.DB.Save(&mock_data_product)
// 	config.DB.Save(&mock_data_product_guarantee)
// 	config.DB.Save(&mock_data_photo)

// 	req := httptest.NewRequest(http.MethodGet, "/products", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	middleware.JWT([]byte(constant.SECRET_JWT))(GetProductsByUserIDControllerTesting())(context)

// 	res_body := res.Body.String()
// 	var response ProductsResponse
// 	er := json.Unmarshal([]byte(res_body), &response)
// 	if er != nil {
// 		assert.Error(t, er, "error")
// 	}
// 	t.Run("GET /jwt/products", func(t *testing.T) {
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "Successful Operation", response.Message)
// 	})

// }

// // Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// // kondisi request failed
// func TestGetProductsByUserIDControllerNotFound(t *testing.T) {
// 	var testCase = TestCase{
// 		Name:       "product by user id not found",
// 		Path:       "/products",
// 		ExpectCode: http.StatusBadRequest,
// 	}

// 	e := InitEcho()
// 	// Mendapatkan token
// 	token, err := UsingJWT()
// 	if err != nil {
// 		panic(err)
// 	}

// 	InsertMockDataToDB()
// 	config.DB.Save(&mock_data_user)

// 	req := httptest.NewRequest(http.MethodGet, "/products", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	middleware.JWT([]byte(constant.SECRET_JWT))(GetProductsByUserIDControllerTesting())(context)

// 	res_body := res.Body.String()
// 	var response ProductsResponse
// 	er := json.Unmarshal([]byte(res_body), &response)
// 	if er != nil {
// 		assert.Error(t, er, "error")
// 	}
// 	t.Run("GET /jwt/products", func(t *testing.T) {
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "Items not found", response.Message)
// 	})

// }

// // Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// // kondisi request failed
// func TestGetProductsByUserIDControllerFailed(t *testing.T) {
// 	var testCase = TestCase{
// 		Name:       "failed to get one data product by user id",
// 		Path:       "/products",
// 		ExpectCode: http.StatusBadRequest,
// 	}

// 	e := InitEcho()
// 	// Mendapatkan token
// 	token, err := UsingJWT()
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Melakukan penghapusan tabel untuk membuat request failed
// 	config.DB.Migrator().DropTable(&models.Products{})

// 	req := httptest.NewRequest(http.MethodGet, "/products", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	middleware.JWT([]byte(constant.SECRET_JWT))(GetProductsByUserIDControllerTesting())(context)

// 	res_body := res.Body.String()
// 	var response ProductsResponse
// 	er := json.Unmarshal([]byte(res_body), &response)
// 	if er != nil {
// 		assert.Error(t, er, "error")
// 	}
// 	t.Run("GET /jwt/products", func(t *testing.T) {
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "Bad Request", response.Message)
// 	})

// }

// Fungsi testing CreateProductController
func CreateProductControllerTesting() echo.HandlerFunc {
	return CreateProductControllers
}

// // Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// // kondisi request failed
// func TestCreateProductsControllerGuaranteeSuccess(t *testing.T) {
// 	var testCase = TestCase{
// 		Name:       "success to create guarantee product",
// 		Path:       "/products",
// 		ExpectCode: http.StatusOK,
// 	}

// 	e := InitEcho()
// 	// Mendapatkan token
// 	token, err := UsingJWT()
// 	if err != nil {
// 		panic(err)
// 	}

// 	InsertMockDataToDB()

// 	filePath := "download.jpg"
// 	fieldName := "photos"
// 	photo := new(bytes.Buffer)

// 	mw := multipart.NewWriter(photo)

// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	w, err := mw.CreateFormFile(fieldName, filePath)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if _, err := io.Copy(w, file); err != nil {
// 		t.Fatal(err)
// 	}

// 	// close the writer before making the request
// 	mw.Close()

// 	type BodyCreate struct {
// 		Name          string
// 		SubcategoryID int
// 		CityID        int
// 		Price         int
// 		Description   string
// 		Stock         int
// 		Guarantee     []int
// 		Photo         *os.File
// 	}

// 	mock_data := BodyCreate{
// 		Name:          "Kamera DSLR Canon",
// 		SubcategoryID: 1,
// 		CityID:        1,
// 		Price:         50000,
// 		Description:   "Murah yang terbaik",
// 		Stock:         5,
// 		Guarantee:     []int{1},
// 		Photo:         file,
// 	}

// 	body, err := json.Marshal(mock_data)
// 	if err != nil {
// 		t.Error(t, err, "error")
// 	}

// 	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Add("Content-Type", mw.FormDataContentType())
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)

// 	// router is of type http.Handler
// 	// http.Handler.ServeHTTP(res, req)

// 	context.SetPath(testCase.Path)
// 	middleware.JWT([]byte(constant.SECRET_JWT))(CreateProductControllerTesting())(context)

// 	res_body := res.Body.String()
// 	var response ProductsResponse
// 	er := json.Unmarshal([]byte(res_body), &response)
// 	if er != nil {
// 		assert.Error(t, er, "error")
// 	}
// 	t.Run("POST /jwt/products", func(t *testing.T) {
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "Successful Operation", response.Message)
// 	})

// }

// // Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// // kondisi request failed
// func TestCreateProductsControllerFailed(t *testing.T) {
// 	var testCase = TestCase{
// 		Name:       "failed to create product",
// 		Path:       "/products",
// 		ExpectCode: http.StatusBadRequest,
// 	}

// 	e := InitEcho()
// 	// Mendapatkan token
// 	token, err := UsingJWT()
// 	if err != nil {
// 		panic(err)
// 	}

// 	InsertMockDataToDB()
// 	mock_data := models.BodyCreateProducts{
// 		Name:          "Kamera DSLR Canon",
// 		SubcategoryID: 1,
// 		CityID:        1,
// 		Price:         50000,
// 		Description:   "Murah yang terbaik",
// 		Stock:         5,
// 		Guarantee:     []int{1},
// 	}
// 	body, err := json.Marshal(mock_data)
// 	if err != nil {
// 		t.Error(t, err, "error")
// 	}

// 	// Melakukan penghapusan tabel untuk membuat request failed
// 	config.DB.Migrator().DropTable(&models.Products{})

// 	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	middleware.JWT([]byte(constant.SECRET_JWT))(CreateProductControllerTesting())(context)

// 	res_body := res.Body.String()
// 	var response ProductsResponse
// 	er := json.Unmarshal([]byte(res_body), &response)
// 	if er != nil {
// 		assert.Error(t, er, "error")
// 	}
// 	t.Run("POST /jwt/products", func(t *testing.T) {
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "Bad Request", response.Message)
// 	})

// }

// // Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// // kondisi request failed
// func TestCreateProductsControllerGuaranteeFailed(t *testing.T) {
// 	var testCase = TestCase{
// 		Name:       "failed to create guarantee product",
// 		Path:       "/products",
// 		ExpectCode: http.StatusBadRequest,
// 	}

// 	e := InitEcho()
// 	// Mendapatkan token
// 	token, err := UsingJWT()
// 	if err != nil {
// 		panic(err)
// 	}

// 	InsertMockDataToDB()
// 	config.DB.Save(&mock_data_user)
// 	mock_data := models.BodyCreateProducts{
// 		Name:          "Kamera DSLR Canon",
// 		SubcategoryID: 1,
// 		CityID:        1,
// 		Price:         50000,
// 		Description:   "Murah yang terbaik",
// 		Stock:         5,
// 		Guarantee:     []int{2},
// 	}
// 	body, err := json.Marshal(mock_data)
// 	if err != nil {
// 		t.Error(t, err, "error")
// 	}

// 	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	middleware.JWT([]byte(constant.SECRET_JWT))(CreateProductControllerTesting())(context)

// 	res_body := res.Body.String()
// 	var response ProductsResponse
// 	er := json.Unmarshal([]byte(res_body), &response)
// 	if er != nil {
// 		assert.Error(t, er, "error")
// 	}
// 	t.Run("POST /jwt/products", func(t *testing.T) {
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "Bad Request", response.Message)
// 	})

// }

// // Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// // kondisi request failed
// func TestCreateProductsControllerNilPrice(t *testing.T) {
// 	var testCase = TestCase{
// 		Name:       "Price must be more than 0",
// 		Path:       "/products",
// 		ExpectCode: http.StatusBadGateway,
// 	}

// 	e := InitEcho()
// 	// Mendapatkan token
// 	token, err := UsingJWT()
// 	if err != nil {
// 		panic(err)
// 	}

// 	InsertMockDataToDB()

// 	mock_data := models.BodyCreateProducts{
// 		Name:          "Kamera DSLR Canon",
// 		SubcategoryID: 1,
// 		CityID:        1,
// 		Price:         0,
// 		Description:   "Murah yang terbaik",
// 		Stock:         5,
// 		Guarantee:     []int{1},
// 	}
// 	body, err := json.Marshal(mock_data)
// 	if err != nil {
// 		t.Error(t, err, "error")
// 	}

// 	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	middleware.JWT([]byte(constant.SECRET_JWT))(CreateProductControllerTesting())(context)

// 	res_body := res.Body.String()
// 	var response ProductsResponse
// 	er := json.Unmarshal([]byte(res_body), &response)
// 	if er != nil {
// 		assert.Error(t, er, "error")
// 	}
// 	t.Run("POST /jwt/products", func(t *testing.T) {
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "Price must be more than 0", response.Message)
// 	})

// }

// Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// kondisi request failed
func TestCreateProductsControllerNilName(t *testing.T) {
	var testCase = TestCase{
		Name:       "Must add name ",
		Path:       "/products",
		ExpectCode: http.StatusBadGateway,
	}

	e := InitEcho()
	// Mendapatkan token
	token, err := UsingJWT()
	if err != nil {
		panic(err)
	}

	InsertMockDataToDB()

	mock_data := models.BodyCreateProducts{
		SubcategoryID: 1,
		CityID:        1,
		Price:         50000,
		Description:   "Murah yang terbaik",
		Stock:         5,
	}

	body, err := json.Marshal(mock_data)
	if err != nil {
		t.Error(t, err, "error")
	}

	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCase.Path)
	middleware.JWT([]byte(constant.SECRET_JWT))(CreateProductControllerTesting())(context)

	res_body := res.Body.String()
	var response ProductsResponse
	er := json.Unmarshal([]byte(res_body), &response)
	if er != nil {
		assert.Error(t, er, "error")
	}
	t.Run("POST /jwt/products", func(t *testing.T) {
		assert.Equal(t, testCase.ExpectCode, res.Code)
		assert.Equal(t, "Must add product name", response.Message)
	})

}

// // Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// // kondisi request failed
// func TestCreateProductsControllerNilStock(t *testing.T) {
// 	var testCase = TestCase{
// 		Name:       "Stock must be more than 0",
// 		Path:       "/products",
// 		ExpectCode: http.StatusBadGateway,
// 	}

// 	e := InitEcho()
// 	// Mendapatkan token
// 	token, err := UsingJWT()
// 	if err != nil {
// 		panic(err)
// 	}

// 	InsertMockDataToDB()

// 	mock_data := models.BodyCreateProducts{
// 		Name:          "Kamera Canon",
// 		SubcategoryID: 1,
// 		CityID:        1,
// 		Price:         50000,
// 		Description:   "Murah yang terbaik",
// 		Stock:         0,
// 		Guarantee:     []int{1},
// 	}
// 	body, err := json.Marshal(mock_data)
// 	if err != nil {
// 		t.Error(t, err, "error")
// 	}

// 	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	middleware.JWT([]byte(constant.SECRET_JWT))(CreateProductControllerTesting())(context)

// 	res_body := res.Body.String()
// 	var response ProductsResponse
// 	er := json.Unmarshal([]byte(res_body), &response)
// 	if er != nil {
// 		assert.Error(t, er, "error")
// 	}
// 	t.Run("POST /jwt/products", func(t *testing.T) {
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "Stock must be more than 0", response.Message)
// 	})

// }

// // Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// // kondisi request failed
// func TestCreateProductsControllerNilDesc(t *testing.T) {
// 	var testCase = TestCase{
// 		Name:       "Must add desc",
// 		Path:       "/products",
// 		ExpectCode: http.StatusBadGateway,
// 	}

// 	e := InitEcho()
// 	// Mendapatkan token
// 	token, err := UsingJWT()
// 	if err != nil {
// 		panic(err)
// 	}

// 	InsertMockDataToDB()

// 	mock_data := models.BodyCreateProducts{
// 		Name:          "Kamera Canon",
// 		SubcategoryID: 1,
// 		CityID:        1,
// 		Price:         50000,
// 		Stock:         1,
// 		Guarantee:     []int{1},
// 	}
// 	body, err := json.Marshal(mock_data)
// 	if err != nil {
// 		t.Error(t, err, "error")
// 	}

// 	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath(testCase.Path)
// 	middleware.JWT([]byte(constant.SECRET_JWT))(CreateProductControllerTesting())(context)

// 	res_body := res.Body.String()
// 	var response ProductsResponse
// 	er := json.Unmarshal([]byte(res_body), &response)
// 	if er != nil {
// 		assert.Error(t, er, "error")
// 	}
// 	t.Run("POST /jwt/products", func(t *testing.T) {
// 		assert.Equal(t, testCase.ExpectCode, res.Code)
// 		assert.Equal(t, "Must add description", response.Message)
// 	})

// }

// Fungsi testing DeleteProductController
func DeleteProductControllerTesting() echo.HandlerFunc {
	return DeleteProductByIDController
}

// Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// kondisi request failed
func TestDeleteProductsControllerSuccess(t *testing.T) {
	var testCase = TestCase{
		Name:       "Success delete product",
		Path:       "/products",
		ExpectCode: http.StatusOK,
	}

	e := InitEcho()
	// Mendapatkan token
	token, err := UsingJWT()
	if err != nil {
		panic(err)
	}

	InsertMockDataToDB()
	config.DB.Save(&mock_data_product)

	req := httptest.NewRequest(http.MethodDelete, "/products", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCase.Path)
	context.SetParamNames("id")
	context.SetParamValues("1")
	middleware.JWT([]byte(constant.SECRET_JWT))(DeleteProductControllerTesting())(context)

	res_body := res.Body.String()
	var response ProductsResponse
	er := json.Unmarshal([]byte(res_body), &response)
	if er != nil {
		assert.Error(t, er, "error")
	}
	t.Run("DELETE /jwt/products", func(t *testing.T) {
		assert.Equal(t, testCase.ExpectCode, res.Code)
		assert.Equal(t, "Successful Operation", response.Message)
	})

}

// Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// kondisi request failed
func TestDeleteProductsControllerFailed(t *testing.T) {
	var testCase = TestCase{
		Name:       "Failed to delete product",
		Path:       "/products/:id",
		ExpectCode: http.StatusBadRequest,
	}

	e := InitEcho()
	// Mendapatkan token
	token, err := UsingJWT()
	if err != nil {
		panic(err)
	}

	// Melakukan penghapusan tabel untuk membuat request failed
	config.DB.Migrator().DropTable(&models.Products{})

	req := httptest.NewRequest(http.MethodDelete, "/products/:id", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCase.Path)
	context.SetParamNames("id")
	context.SetParamValues("1")
	middleware.JWT([]byte(constant.SECRET_JWT))(DeleteProductControllerTesting())(context)

	res_body := res.Body.String()
	var response ProductsResponse
	er := json.Unmarshal([]byte(res_body), &response)
	if er != nil {
		assert.Error(t, er, "error")
	}
	t.Run("DELETE /jwt/products/:id", func(t *testing.T) {
		assert.Equal(t, testCase.ExpectCode, res.Code)
		assert.Equal(t, "Bad Request", response.Message)
	})

}

// Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// kondisi request failed
func TestDeleteProductsControllerFalseParam(t *testing.T) {
	var testCase = TestCase{
		Name:       "False param",
		Path:       "/products",
		ExpectCode: http.StatusBadRequest,
	}

	e := InitEcho()
	// Mendapatkan token
	token, err := UsingJWT()
	if err != nil {
		panic(err)
	}

	req := httptest.NewRequest(http.MethodDelete, "/products/:id", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCase.Path)
	context.SetParamNames("id")
	context.SetParamValues("!")
	middleware.JWT([]byte(constant.SECRET_JWT))(DeleteProductControllerTesting())(context)

	res_body := res.Body.String()
	var response ProductsResponse
	er := json.Unmarshal([]byte(res_body), &response)
	if er != nil {
		assert.Error(t, er, "error")
	}
	t.Run("DELETE /jwt/products/:id", func(t *testing.T) {
		assert.Equal(t, testCase.ExpectCode, res.Code)
		assert.Equal(t, "False Param", response.Message)
	})

}

// Fungsi untuk melakukan testing fungsi GetProductsBySubcategoryIDControllers
// kondisi request failed
func TestDeleteProductsControllerNotAllowed(t *testing.T) {
	var testCase = TestCase{
		Name:       "Not allowed to delet",
		Path:       "/products/:id",
		ExpectCode: http.StatusBadRequest,
	}

	e := InitEcho()
	// Mendapatkan token
	token, err := UsingJWT()
	if err != nil {
		panic(err)
	}

	InsertMockDataToDB()
	config.DB.Save(&mock_data_user2)
	config.DB.Save(&mock_data_product)
	config.DB.Save(&mock_data_product2)

	req := httptest.NewRequest(http.MethodDelete, "/products/:id", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCase.Path)
	context.SetParamNames("id")
	context.SetParamValues("2")
	middleware.JWT([]byte(constant.SECRET_JWT))(DeleteProductControllerTesting())(context)

	res_body := res.Body.String()
	var response ProductsResponse
	er := json.Unmarshal([]byte(res_body), &response)
	if er != nil {
		assert.Error(t, er, "error")
	}
	t.Run("DELETE /jwt/products/:id", func(t *testing.T) {
		assert.Equal(t, testCase.ExpectCode, res.Code)
		assert.Equal(t, "Access Forbidden", response.Message)
	})

}
