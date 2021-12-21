package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Total   int       `json:"total" form:"total"`
	Booking []Booking `gorm:"foreignKey:TransactionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
