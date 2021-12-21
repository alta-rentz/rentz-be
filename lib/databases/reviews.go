package databases

import (
	"project3/config"
	"project3/models"
)

func AddReviews(review *models.Reviews) (interface{}, error) {
	if err := config.DB.Create(&review).Error; err != nil {
		return nil, err
	}
	return review, nil
}

func AddRatingToProduct(id int) {
	config.DB.Exec("UPDATE products SET rating = (SELECT AVG(rating) FROM reviews WHERE products_id = ?) WHERE id = ?", id, id)
}

func GetProductID(id int) (int, error) {
	var productID int
	tx := config.DB.Raw("SELECT products_id FROM bookings WHERE id = ?", id).Scan(&productID)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return 0, tx.Error
	}
	return productID, nil
}

func GetBookingOwner(id int) (int, error) {
	var userID int
	tx := config.DB.Raw("SELECT carts.users_id FROM bookings JOIN carts ON bookings.cart_id = carts.id WHERE bookings.id = ?", id).Scan(&userID)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return 0, tx.Error
	}
	return userID, nil
}

func GetBookingStatus(id int) (string, error) {
	var statusPayment string
	tx := config.DB.Raw("SELECT status_payment FROM bookings WHERE id = ?", id).Scan(&statusPayment)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return "", tx.Error
	}
	return statusPayment, nil
}
