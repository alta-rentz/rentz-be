package controllers

import (
	"io"
	"net/http"
	"net/url"
	"project3/lib/databases"
	"project3/middlewares"
	"project3/models"
	"project3/plugins"
	"project3/response"
	"strconv"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
)

const MAX_UPLOAD_SIZE = 1024 * 1024 // 1MB

// Controller untuk membuat product baru
func CreateProductControllers(c echo.Context) error {
	var storageClient *storage.Client
	var body models.BodyCreateProducts
	c.Bind(&body)
	logged := middlewares.ExtractTokenUserId(c)

	// create input product
	var product models.Products
	if body.Name == "" {
		return c.JSON(http.StatusBadGateway, response.ProductsBadGatewayResponse("Must add product name"))
	}
	spaceEmpty := strings.TrimSpace(body.Name)
	product.Name = spaceEmpty
	product.SubcategoryID = body.SubcategoryID
	product.CityID = body.CityID
	if body.Price <= 0 {
		return c.JSON(http.StatusBadGateway, response.ProductsBadGatewayResponse("Price must be more than 0"))
	}
	product.Price = body.Price
	product.Description = body.Description
	if body.Description == "" {
		return c.JSON(http.StatusBadGateway, response.ProductsBadGatewayResponse("Must add description"))
	}
	product.Stock = body.Stock
	if body.Stock <= 0 {
		return c.JSON(http.StatusBadGateway, response.ProductsBadGatewayResponse("Stock must be more than 0"))
	}
	product.UsersID = uint(logged)
	getCity, _ := databases.GetCity(product.CityID)
	lat, long, _ := plugins.Geocode(getCity)
	product.Latitude = lat
	product.Longitude = long

	bucket := "rentz-id" //your bucket name

	ctx := appengine.NewContext(c.Request())

	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile("keys.json"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.UploadErrorResponse(err))
	}

	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File["photos"]
	if files == nil {
		return c.JSON(http.StatusBadGateway, response.ProductsBadGatewayResponse("must add photo"))
	}

	createdProduct, err := databases.CreateProduct(&product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}

	for _, file := range files {
		if file.Size > MAX_UPLOAD_SIZE {
			databases.DeleteProduct(int(createdProduct.ID))
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "The uploaded image is too big. Please use an image less than 1MB in size",
			})
		}

		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		if file.Filename[len(file.Filename)-3:] != "jpg" && file.Filename[len(file.Filename)-3:] != "png" {
			if file.Filename[len(file.Filename)-4:] != "jpeg" {
				databases.DeleteProduct(int(createdProduct.ID))
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "The provided file format is not allowed. Please upload a JPG or JPEG or PNG image",
				})
			}
		}

		sw := storageClient.Bucket(bucket).Object(file.Filename).NewWriter(ctx)

		if _, err := io.Copy(sw, src); err != nil {
			return c.JSON(http.StatusInternalServerError, response.UploadErrorResponse(err))
		}

		if err := sw.Close(); err != nil {
			return c.JSON(http.StatusInternalServerError, response.UploadErrorResponse(err))
		}

		u, err := url.Parse("https://storage.googleapis.com/" + bucket + "/" + sw.Attrs().Name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.UploadErrorResponse(err))
		}
		photo := models.Photos{
			Photo_Name: sw.Attrs().Name,
			Url:        u.String(),
			ProductsID: uint(createdProduct.ID),
		}
		_, err = databases.InsertPhoto(&photo)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
		}

	}

	// add product guarantee
	for _, guarantee := range body.Guarantee {
		var input = models.ProductsGuarantee{
			ProductsID:  createdProduct.ID,
			GuaranteeID: uint(guarantee),
		}
		_, err := databases.InsertGuarantee(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "product created and file uploaded successfully",
	})
}

// Controller untuk mendapatkan seluruh product
func GetAllProductsController(c echo.Context) error {
	product, err := databases.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	if product == nil {
		return c.JSON(http.StatusBadRequest, response.ItemsNotFoundResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData(product))
}

// Controller untuk mendapatkan product berdasarkan product id
func GetProductByIDController(c echo.Context) error {
	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	product, err := databases.GetProductByID(uint(productId))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	if product == nil {
		return c.JSON(http.StatusBadRequest, response.ItemsNotFoundResponse())
	}
	product.Url, _ = databases.GetUrl(uint(productId))
	product.Guarantee, _ = databases.GetGuarantee(productId)
	return c.JSON(http.StatusOK, response.SuccessResponseData(product))
}

// Controller untuk mendapatkan seluruh product berdasarkan subcategory id
func GetProductsBySubcategoryIDController(c echo.Context) error {
	subcategoryId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	product, err := databases.GetProductsBySubcategoryID(subcategoryId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	if product == nil {
		return c.JSON(http.StatusBadRequest, response.ItemsNotFoundResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData(product))
}

// Controller untuk mendapatkan seluruh product berdasarkan subcategory id
func GetProductsByUserIDController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	product, err := databases.GetProductsBySubcategoryID(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	if product == nil {
		return c.JSON(http.StatusBadRequest, response.ItemsNotFoundResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData(product))
}
