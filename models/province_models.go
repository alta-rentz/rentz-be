package models

type Province struct {
	ID            int    `gorm:"primaryKey;autoIncrement;not null" json:"id" form:"id"`
	Province_Name string `gorm:"type:varchar(255);not null" json:"province" form:"province"`
	City          []City `gorm:"foreignKey:ProvinceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
