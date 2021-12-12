package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UsersID    uint         `json:"user_id" form:"user_id"`
	CartDetail []CartDetail `gorm:"foreignKey:CartID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
