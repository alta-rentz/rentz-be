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
func GetBookingById(id int) (*models.GetBookingDetail, error) {
	var rent models.Booking
	var rentdetail models.GetBookingDetail
	tx := config.DB.Model(rent).Where("id = ?", id).Find(&rentdetail)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return nil, tx.Error
	}
	return &rentdetail, nil
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

func GetBookingByCartID(id int) (interface{}, error) {
	var booking []models.GetBookingDetail
	tx := config.DB.Model(models.Booking{}).Where("cart_id=? AND status_payment = \"waiting\"", id).Find(&booking)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return nil, tx.Error
	}
	return booking, nil
}

func GetHistoryByCartID(id int) (interface{}, error) {
	var booking []models.GetBookingDetail
	tx := config.DB.Model(models.Booking{}).Where("cart_id=? AND status_payment = \"succes\"", id).Find(&booking)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return nil, tx.Error
	}
	return booking, nil
}

type RentDate struct {
	Time_In  time.Time
	Time_Out time.Time
}

// Fungsi untuk mendapatkan tanggal check_in dan time_out suatu reservasi
func ProductRentList(id int) ([]RentDate, error) {
	var dates []RentDate
	tx := config.DB.Table("bookings").Select("bookings.time_in, bookings.time_out").Where("bookings.products_id = ? AND bookings.time_out > ?", id, time.Now()).Find(&dates)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return nil, tx.Error
	}
	return dates, nil
}

// Fungsi untuk menambahkan harga pada reservasi
func GetHargaRoom(idRoom int) (int, error) {
	var harga int
	tx := config.DB.Raw("SELECT harga FROM rooms WHERE id = ?", idRoom).Scan(&harga)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return 0, tx.Error
	}
	return harga, nil
}
