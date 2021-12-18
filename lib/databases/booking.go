package databases

import (
	"project3/config"
	"project3/models"
	"time"
)

// Fungsi untuk membuat booking rental baru
func CreateBooking(rent models.Booking, cart_id int) (*models.Booking, error) {
	tx := config.DB.Where("products_id=? AND cart_id=? and transaction_id=0", rent.ProductsID, cart_id).Find(&models.Booking{})
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected < 1 {
		if err := config.DB.Save(&rent).Error; err != nil {
			return nil, err
		} else {
			return &rent, nil
		}
	}
	return nil, nil
}

// Fungsi untuk mendapatkan akumulasi hari rental
func AccumulatedDays(timeIn time.Time, timeOut time.Time, idBooking uint) {
	config.DB.Exec("UPDATE bookings SET total_day = (SELECT DATEDIFF(?, ?)) WHERE id = ?", timeOut, timeIn, idBooking)
}

// Fungsi untuk mentotal harga booking
func AddPriceBooking(idProduct, idBooking uint) int {
	var price int
	config.DB.Exec("UPDATE bookings SET total = (SELECT price FROM products WHERE id = ?)*total_day*qty WHERE id = ?", idProduct, idBooking)
	tx := config.DB.Raw("SELECT total FROM bookings WHERE id = ?", idBooking).Scan(&price)
	if tx.Error != nil {
		return 0
	}
	return price
}

// Fungsi untuk mendapatkan informasi booking by id
func GetBookingById(id int) (*models.Booking, error) {
	var rent models.Booking
	tx := config.DB.Where("id = ?", id).Find(&rent)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return nil, tx.Error
	}
	return &rent, nil
}

// Fungsi untuk menghapus booking
func CancelBooking(id int) (interface{}, error) {
	var rent models.Booking
	if err := config.DB.Where("id = ?", id).Delete(&rent).Error; err != nil {
		return nil, err
	}
	return "deleted", nil
}

func GetCartId(user_id int) (models.Cart, error) {
	var cart models.Cart
	if err := config.DB.Where("users_id=?", user_id).Find(&cart).Error; err != nil {
		return cart, err
	}
	return cart, nil
}
