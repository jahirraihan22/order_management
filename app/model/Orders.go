package model

import "time"

type Order struct {
	ID                 uint      `gorm:"primaryKey"`
	StoreId            int       `json:"store_id" gorm:"not null"`
	MerchantOrderId    string    `json:"merchant_order_id"`
	RecipientName      string    `json:"recipient_name"  gorm:"not null"`
	RecipientPhone     string    `json:"recipient_phone"  gorm:"not null"`
	RecipientAddress   string    `json:"recipient_address"  gorm:"not null"`
	RecipientCity      int       `json:"recipient_city"  gorm:"not null"`
	RecipientZone      int       `json:"recipient_zone"  gorm:"not null"`
	RecipientArea      int       `json:"recipient_area"  gorm:"not null"`
	DeliveryType       int       `json:"delivery_type"  gorm:"not null"`
	ItemType           int       `json:"item_type"  gorm:"not null"`
	SpecialInstruction string    `json:"special_instruction"`
	ItemQuantity       int       `json:"item_quantity"  gorm:"not null"`
	ItemWeight         float64   `json:"item_weight"  gorm:"not null"`
	AmountToCollect    float64   `json:"amount_to_collect" gorm:"not null"`
	ItemDescription    string    `json:"item_description" `
	ConsignmentId      string    `json:"consignment_id"`
	OrderStatus        string    `json:"order_status" gorm:"not null"`
	DeliveryFee        float64   `json:"delivery_fee"  gorm:"not null"`
	CODFee             float64   `json:"cod_fee"  gorm:"not null;default:0"`
	CreatedAt          time.Time `json:"created_at"  gorm:"not null"`
	UpdatedAt          time.Time `json:"updated_at"  gorm:"not null"`
	UserID             uint
	User               Users `gorm:"foreignKey:UserID"`
}

type OrderDTO struct {
	StoreId            int     `json:"store_id" validate:"required,valid_store"`        // fixed
	MerchantOrderId    string  `json:"merchant_order_id"`                               // optional
	RecipientName      string  `json:"recipient_name" validate:"required"`              // required
	RecipientPhone     string  `json:"recipient_phone" validate:"required,valid_phone"` // required
	RecipientAddress   string  `json:"recipient_address" validate:"required"`           // required
	RecipientCity      int     `json:"recipient_city" validate:"required,city"`         // fixed
	RecipientZone      int     `json:"recipient_zone"  validate:"required,zone"`        // fixed
	RecipientArea      int     `json:"recipient_area" validate:"required,area"`         // fixed
	DeliveryType       int     `json:"delivery_type" validate:"required,ge=48"`         // fixed
	ItemType           int     `json:"item_type" validate:"required,eq=2"`              // fixed
	SpecialInstruction string  `json:"special_instruction"`                             // optional
	ItemQuantity       int     `json:"item_quantity" validate:"required,ge=1"`
	ItemWeight         float64 `json:"item_weight" validate:"required,ge=0"`
	AmountToCollect    float64 `json:"amount_to_collect" validate:"required,gt=0"`
	ItemDescription    string  `json:"item_description"`
}
