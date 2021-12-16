package models

type Subcategory struct {
	ID               int        `gorm:"primaryKey;autoIncrement;not null" json:"id" form:"id"`
	Subcategory_Name string     `gorm:"type:varchar(50);not null" json:"subcategory" form:"subcategory"`
	CategoryID       uint       `json:"category_id" form:"category_id"`
	Products         []Products `gorm:"foreignKey:SubcategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
