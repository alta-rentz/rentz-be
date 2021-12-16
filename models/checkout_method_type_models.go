package models

type CheckoutMethodType struct {
	ID             int              `gorm:"primaryKey;autoIncrement;not null" json:"id" form:"id"`
	Checkout_Type  string           `gorm:"type:varchar(100)" json:"checkout_type" form:"checkout_type"`
	CheckoutMethod []CheckoutMethod `gorm:"foreignKey:CheckoutMethodTypeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
