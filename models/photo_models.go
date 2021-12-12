package models

import (
	"gorm.io/gorm"
)

type Photos struct {
	gorm.Model
	Photo_Name string     `gorm:"type:longtext;not null" json:"photo_name" form:"photo_name"`
	Url        string     `gorm:"type:longtext" json:"url" form:"url"`
	Products   []Products `gorm:"foreignKey:PhotosID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
