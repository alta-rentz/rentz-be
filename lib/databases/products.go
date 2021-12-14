package databases

import (
	"project3/config"
	"project3/models"
)

// Fungsi untuk membuat menyewakan produk baru
func CreateProduct(product *models.Products) (*models.Products, error) {
	if err := config.DB.Create(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

// Fungsi untuk mendapatkan nama kota dari city_id
func GetCity(id int) (string, error) {
	var town models.City
	if err := config.DB.Where("id = ?", id).Find(&town); err.Error != nil {
		return "", err.Error
	}
	return town.City_Name, nil
}

// Fungsi untuk memasukkan foto product
func InsertPhoto(photo *models.Photos) (interface{}, error) {
	if err := config.DB.Create(&photo).Error; err != nil {
		return nil, err
	}
	return photo, nil
}

// Fungsi untuk memasukkan guarantee product
func InsertGuarantee(guarantee *models.ProductsGuarantee) (interface{}, error) {
	if err := config.DB.Create(&guarantee).Error; err != nil {
		return nil, err
	}
	return guarantee, nil
}

// Fungsi untuk mendapatkan seluruh product
func GetAllProducts() (interface{}, error) {
	var results []models.GetAllProduct
	tx := config.DB.Table("products").Select("products.id, products.users_id, products.name, subcategories.subcategory_name, products.subcategory_id, products.city_id, products.price, products.description, products.stock, photos.url").Group("products.id").Joins(
		"join subcategories on subcategories.id = products.subcategory_id").Joins(
		"join photos on photos.products_id = products.id").Where("products.deleted_at IS NULL").Find(&results)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return results, nil
}
