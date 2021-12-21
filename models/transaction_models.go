package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Total   int       `json:"total" form:"total"`
	Booking []Booking `gorm:"foreignKey:TransactionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

//
type CheckOut struct {
	User_ID        int    `json:"user_id" form:"user_id"`
	CheckoutMethod string `json:"checkout_method" form:"checkout_method"`
	Total          int    `json:"total" form:"total"`
	Booking_ID     []int  `json:"booking_id" form:"booking_id"`
	Phone          string `json:"phone" form:"phone"`
}

//
type RequestBodyStruct struct {
	ReferenceID       string            `json:"reference_id" validate:"required"`
	Currency          string            `json:"currency" validate:"required"`
	Amount            float64           `json:"amount" validate:"required"`
	CheckoutMethod    string            `json:"checkout_method" validate:"required"`
	ChannelCode       string            `json:"channel_code" validate:"required"`
	ChannelProperties ChannelProperties `json:"channel_properties" validate:"required"`
	Metadata          Metadata          `json:"metadata" validate:"required"`
}

//
type Metadata struct {
	BranchArea string `json:"branch_area" validate:"required"`
	BranchCity string `json:"branch_city" validate:"required"`
}

//
type ChannelProperties struct {
	Success_redirect_url string `json:"success_redirect_url" validate:"required"`
}

//
type BasketItem struct {
	ReferenceID string                 `json:"reference_id" validate:"required"`
	Name        string                 `json:"name" validate:"required"`
	Category    string                 `json:"category" validate:"required"`
	Currency    string                 `json:"currency" validate:"required"`
	Price       float64                `json:"price" validate:"required"`
	Quantity    int                    `json:"quantity" validate:"required"`
	Type        string                 `json:"type" validate:"required"`
	Url         string                 `json:"url,omitempty"`
	Description string                 `json:"description,omitempty"`
	SubCategory string                 `json:"sub_category,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type ResponsePayment struct {
	ID                 string            `json:"id"`
	BusinessID         string            `json:"business_id"`
	ReferenceID        string            `json:"reference_id"`
	Status             string            `json:"status"`
	Currency           string            `json:"currency"`
	ChargeAmount       float64           `json:"charge_amount"`
	CaptureAmount      float64           `json:"capture_amount"`
	CheckoutMethod     string            `json:"checkout_method"`
	ChannelCode        string            `json:"channel_code"`
	ChannelProperties  ChannelProperties `json:"channel_properties"`
	Actions            map[string]string `json:"actions"`
	IsRedirectRequired bool              `json:"is_redirect_required"`
	CallbackURL        string            `json:"callback_url"`
	Created            string            `json:"created"`
	Updated            string            `json:"updated"`
	VoidedAt           string            `json:"voided_at,omitempty"`
	CaptureNow         bool              `json:"capture_now"`
	CustomerID         string            `json:"customer_id,omitempty"`
	PaymentMethodID    string            `json:"payment_method_id,omitempty"`
	FailureCode        string            `json:"failure_code,omitempty"`
	Basket             []BasketItem      `json:"basket,omitempty"`
	Metadata           Metadata          `json:"metadata,omitempty"`
}

type RequestBodyStructOVO struct {
	ReferenceID       string               `json:"reference_id" validate:"required"`
	Currency          string               `json:"currency" validate:"required"`
	Amount            float64              `json:"amount" validate:"required"`
	CheckoutMethod    string               `json:"checkout_method" validate:"required"`
	ChannelCode       string               `json:"channel_code" validate:"required"`
	ChannelProperties ChannelPropertiesOVO `json:"channel_properties" validate:"required"`
	Metadata          Metadata             `json:"metadata" validate:"required"`
}

type MetadataOVO struct {
	BranchArea string `json:"branch_area" validate:"required"`
	BranchCity string `json:"branch_city" validate:"required"`
}

type ChannelPropertiesOVO struct {
	MobileNumber string `json:"mobile_number" validate:"required"`
}

type ResponsePaymentOVO struct {
	ID                 string               `json:"id"`
	BusinessID         string               `json:"business_id"`
	ReferenceID        string               `json:"reference_id"`
	Status             string               `json:"status"`
	Currency           string               `json:"currency"`
	ChargeAmount       float64              `json:"charge_amount"`
	CaptureAmount      float64              `json:"capture_amount"`
	CheckoutMethod     string               `json:"checkout_method"`
	ChannelCode        string               `json:"channel_code"`
	ChannelProperties  ChannelPropertiesOVO `json:"channel_properties"`
	Actions            map[string]string    `json:"actions"`
	IsRedirectRequired bool                 `json:"is_redirect_required"`
	CallbackURL        string               `json:"callback_url"`
	Created            string               `json:"created"`
	Updated            string               `json:"updated"`
	VoidedAt           string               `json:"voided_at,omitempty"`
	CaptureNow         bool                 `json:"capture_now"`
	CustomerID         string               `json:"customer_id,omitempty"`
	PaymentMethodID    string               `json:"payment_method_id,omitempty"`
	FailureCode        string               `json:"failure_code,omitempty"`
	Basket             []BasketItemOvo      `json:"basket,omitempty"`
	Metadata           Metadata             `json:"metadata,omitempty"`
}

type BasketItemOvo struct {
	ReferenceID string                 `json:"reference_id" validate:"required"`
	Name        string                 `json:"name" validate:"required"`
	Category    string                 `json:"category" validate:"required"`
	Currency    string                 `json:"currency" validate:"required"`
	Price       float64                `json:"price" validate:"required"`
	Quantity    int                    `json:"quantity" validate:"required"`
	Type        string                 `json:"type" validate:"required"`
	Url         string                 `json:"url,omitempty"`
	Description string                 `json:"description,omitempty"`
	SubCategory string                 `json:"sub_category,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}
