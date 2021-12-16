package models

type City struct {
	ID         int        `gorm:"primaryKey;autoIncrement;not null" json:"id" form:"id"`
	City_Name  string     `gorm:"type:varchar(50);not null" json:"city" form:"city"`
	ProvinceID uint       `json:"province_id" form:"province_id"`
	Products   []Products `gorm:"foreignKey:CityID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
