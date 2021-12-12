package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Total            int          `json:"total" form:"total"`
	CheckoutMethodID int          `json:"checkout_method_id" form:"checkout_method_id"`
	CartDetail       []CartDetail `gorm:"foreignKey:TransactionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
