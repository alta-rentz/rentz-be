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

	"cloud.google.com/go/storage"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
)

func CreateProductControllers(c echo.Context) error {
	var storageClient *storage.Client
	var body models.BodyCreateProducts
	c.Bind(&body)
	logged := middlewares.ExtractTokenUserId(c)

	// create input product
	var product models.Products
	product.Name = body.Name
	product.SubcategoryID = body.SubcategoryID
	product.CityID = body.CityID
	product.Price = body.Price
	product.Description = body.Description
	product.Stock = body.Stock
	product.UsersID = uint(logged)
	getCity, _ := databases.GetCity(product.CityID)
	lat, long, _ := plugins.Geocode(getCity)
	product.Latitude = lat
	product.Longitude = long

	createdProduct, err := databases.CreateProduct(&product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
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

	bucket := "rentz-id" //your bucket name

	ctx := appengine.NewContext(c.Request())

	storageClient, err = storage.NewClient(ctx, option.WithCredentialsFile("keys.json"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.UploadErrorResponse(err))
	}

	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["photos"]

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

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

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "product created and file uploaded successfully",
	})
}

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
