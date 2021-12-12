package models

import (
	"gorm.io/gorm"
)

type Reviews struct {
	gorm.Model
	CartDetailID uint   `json:"cart_detail_id" form:"cart_detail_id"`
	Rating       int    `gorm:"type:enum('1','2','3','4','5');not null" json:"rating" form:"rating"`
	Comment      string `gorm:"type:longtext;not null" json:"comment" form:"comment"`
	ProductsID   uint   `json:"product_id" form:"product_id"`
}
