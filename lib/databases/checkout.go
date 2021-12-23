package databases

import (
	"project3/config"
	"project3/models"
)

func GetBooking(booking_id int, cart_id int) (*models.Booking, error) {
	cart := models.Booking{}
	if err := config.DB.Where("id=? AND cart_id=? AND status_payment= \"waiting\"", booking_id, cart_id).Find(&cart).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}

func GetBookings(id []int, cart_id int) ([]models.Booking, error) {
	totCheckout := []models.Booking{}
	for i := 0; i < len(id); i++ {
		bookingId := id[i]
		respon, _ := GetBooking(bookingId, cart_id)
		totCheckout = append(totCheckout, *respon)
	}
	if len(totCheckout) > 0 {
		return totCheckout, nil
	}
	return nil, nil
}

func CreateCheckout(CheckOut models.CheckOut) (*models.Transaction, error) {
	newCheckout := models.Transaction{
		Total: CheckOut.Total,
	}
	// Save Order Detail
	if err := config.DB.Save(&newCheckout).Error; err != nil {
		return nil, err
	}
	// Update Status
	for i := 0; i < len(CheckOut.Booking_ID); i++ {
		tx := config.DB.Model(models.Booking{}).Where("id=? AND status_payment='waiting'", CheckOut.Booking_ID[i]).Updates(models.Booking{Status_Payment: "succes", TransactionID: &newCheckout.ID})
		if err := tx.Error; err != nil {
			return nil, tx.Error
		}
	}
	// Return Transaction ID for Purchase the Order
	config.DB.Find(&newCheckout, newCheckout.ID)
	return &newCheckout, nil
}
