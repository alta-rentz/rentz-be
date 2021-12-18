package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	ProductsID     uint      `json:"product_id" form:"product_id"`
	CartID         uint      `json:"cart_id" form:"cart_id"`
	TransactionID  *uint     `json:"transaction_id" form:"transaction_id"`
	Time_In        time.Time `gorm:"type:datetime;not null" json:"time_in" form:"time_in"`
	Time_Out       time.Time `gorm:"type:datetime;not null" json:"time_out" form:"time_out"`
	Total_Day      int       `json:"total_day" form:"total_day"`
	Qty            int       `gorm:"not null" json:"qty" form:"qty"`
	Total          int       `json:"total" form:"total"`
	Status_Payment string    `gorm:"type:varchar(100)" json:"status_payment" form:"status_payment"`
	Review         Reviews   `gorm:"foreignKey:BookingID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
type BookingBody struct {
	ProductsID uint   `json:"product_id" form:"product_id"`
	Time_In    string `gorm:"type:datetime;not null" json:"time_in" form:"time_in"`
	Time_Out   string `gorm:"type:datetime;not null" json:"time_out" form:"time_out"`
	Qty        int    `gorm:"not null" json:"qty" form:"qty"`
}

type GetBooking struct {
	ProductsID  uint
	Check_In    string
	Check_Out   string
	Qty         int
	Total_Harga int
}
