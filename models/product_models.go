package models

import "gorm.io/gorm"

// struct user
type Products struct {
	gorm.Model
	Name              string              `gorm:"type:varchar(255)" json:"name" form:"name"`
	UsersID           uint                `json:"user_id" form:"user_id"`
	SubcategoryID     int                 `json:"subcategory_id" form:"subcategory_id"`
	CityID            int                 `json:"city_id" form:"city_id"`
	Price             int                 `gorm:"not null" json:"price" form:"price"`
	Description       string              `gorm:"type:longtext;not null" json:"description" form:"description"`
	Stock             int                 `gorm:"type:int;default:1" json:"stock" form:"stock"`
	Latitude          float64             `gorm:"type:varchar(30);not null" json:"lat" form:"lat"`
	Longitude         float64             `gorm:"type:varchar(30);not null" json:"lon" form:"lon"`
	Photos            []Photos            `gorm:"foreignKey:ProductsID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductsGuarantee []ProductsGuarantee `gorm:"foreignKey:ProductsID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Booking           []Booking           `gorm:"foreignKey:ProductsID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Reviews           []Reviews           `gorm:"foreignKey:ProductsID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// struct body create product
type BodyCreateProducts struct {
	Name          string `gorm:"type:varchar(255)" json:"name" form:"name"`
	SubcategoryID int    `json:"subcategory_id" form:"subcategory_id"`
	CityID        int    `json:"city_id" form:"city_id"`
	Price         int    `gorm:"not null" json:"price" form:"price"`
	Description   string `gorm:"type:longtext;not null" json:"description" form:"description"`
	Stock         int    `gorm:"type:int;default:1" json:"stock" form:"stock"`
	Guarantee     []int  `json:"guarantee" form:"guarantee"`
}

// struct get product
type GetAllProduct struct {
	ID               uint
	UsersID          uint
	Name             string
	Subcategory_Name string
	SubcategoryID    int
	CityID           int
	Price            int
	Description      string
	Stock            int
	Url              string
}

// struct get product
type GetProduct struct {
	ID               uint
	UsersID          uint
	Name             string
	Subcategory_Name string
	SubcategoryID    int
	CityID           int
	Price            int
	Description      string
	Stock            int
	Latitude         float64
	Longitude        float64
	Url              []string
	Guarantee        []string
}
