package models

import "gorm.io/gorm"

// struct user
type Products struct {
	gorm.Model
	Name          string  `gorm:"type:varchar(255)" json:"name" form:"name"`
	UsersID       uint    `json:"user_id" form:"user_id"`
	SubcategoryID int     `json:"subcategory_id" form:"subcategory_id"`
	PhotosID      uint    `json:"photo_id" form:"photo_id"`
	CityID        int     `json:"city_id" form:"city_id"`
	Price         int     `gorm:"not null" json:"price" form:"price"`
	Description   string  `gorm:"type:longtext;not null" json:"description" form:"description"`
	Stock         int     `gorm:"type:int;default:1" json:"stock" form:"stock"`
	Address       string  `gorm:"type:varchar(250);not null" json:"address" form:"address"`
	Longitude     float64 `gorm:"type:varchar(30);not null" json:"lon" form:"lon"`
	Latitude      float64 `gorm:"type:varchar(30);not null" json:"lat" form:"lat"`
}
