package models

import (
	"gorm.io/gorm"
)

type Photos struct {
	gorm.Model
	Photo_Name string `gorm:"type:longtext;not null" json:"photo_name" form:"photo_name"`
	Url        string `gorm:"type:longtext" json:"url" form:"url"`
	ProductsID uint   `json:"product_id" form:"product_id"`
}
