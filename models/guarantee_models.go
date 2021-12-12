package models

import (
	"gorm.io/gorm"
)

type Guarantee struct {
	gorm.Model
	Guarantee_Name    string              `gorm:"type:varchar(255);not null" json:"guarantee" form:"guarantee"`
	ProductsGuarantee []ProductsGuarantee `gorm:"foreignKey:GuaranteeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
