package models

type CheckoutMethod struct {
	ID                   int         `gorm:"primaryKey;autoIncrement;not null" json:"id" form:"id"`
	Checkout_Name        string      `gorm:"type:varchar(100)" json:"checkout_name" form:"checkout_name"`
	CheckoutMethodTypeID int         `json:"checkout_method_type_id" form:"checkout_method_type_id"`
	Transaction          Transaction `gorm:"foreignKey:CheckoutMethodID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
