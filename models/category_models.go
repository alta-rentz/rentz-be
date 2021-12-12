package models

type Category struct {
	ID            int           `gorm:"primaryKey;autoIncrement;not null" json:"id" form:"id"`
	Category_Name string        `gorm:"type:varchar(255);not null" json:"category" form:"category"`
	Subcategory   []Subcategory `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
