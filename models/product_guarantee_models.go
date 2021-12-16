package models

import (
	"gorm.io/gorm"
)

type ProductsGuarantee struct {
	gorm.Model
	ProductsID  uint `json:"product_id" form:"product_id"`
	GuaranteeID uint `json:"guarantee_id" form:"guarantee_id"`
}
