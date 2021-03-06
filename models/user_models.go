package models

import "gorm.io/gorm"

// struct user
type Users struct {
	gorm.Model
	Nama         string     `gorm:"type:varchar(255)" json:"nama" form:"nama"`
	Email        string     `gorm:"type:varchar(100);unique;not null" json:"email" form:"email"`
	Password     string     `gorm:"type:longtext;not null" json:"password" form:"password"`
	Phone_Number string     `gorm:"type:varchar(100);unique;not null" json:"phone" form:"phone"`
	Token        string     `gorm:"type:longtext" json:"token" form:"token"`
	Products     []Products `gorm:"foreignKey:UsersID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Cart         Cart       `gorm:"foreignKey:UsersID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UserLogin struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type Get_User struct {
	ID           uint
	Nama         string
	Email        string
	Phone_Number string
	Token        string
}
