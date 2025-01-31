package response

import (
	"order_management/app/model"
	"time"
)

type PaginatedOrderResponse struct {
	Orders   []*Order `json:"orders"`
	Paginate Paginate `json:"paginate"`
}
type Order struct {
	ConsignmentId    string    `json:"consignment_id,omitempty"`
	OrderCreatedAt   time.Time `json:"order_created_at,omitempty"`
	OrderDescription string    `json:"order_description,omitempty"`
	MerchantOrderId  string    `json:"merchant_order_id,omitempty"`
	RecipientName    string    `json:"recipient_name,omitempty"`
	RecipientAddress string    `json:"recipient_address,omitempty"`
	RecipientPhone   string    `json:"recipient_phone,omitempty"`
	OrderAmount      float64   `json:"order_amount,omitempty"`
	TotalFee         float64   `json:"total_fee,omitempty"`
	Instruction      string    `json:"instruction,omitempty"`
	OrderTypeId      int       `json:"order_type_id,omitempty"`
	CodFee           float64   `json:"cod_fee,omitempty"`
	PromoDiscount    int       `json:"promo_discount,omitempty"`
	Discount         int       `json:"discount,omitempty"`
	DeliveryFee      float64   `json:"delivery_fee,omitempty"`
	OrderStatus      string    `json:"order_status,omitempty"`
	OrderType        string    `json:"order_type,omitempty"`
	ItemType         string    `json:"item_type,omitempty"`
}
type Paginate struct {
	Total       int64 `json:"total,omitempty"`
	CurrentPage int64 `json:"current_page,omitempty"`
	PerPage     int64 `json:"per_page,omitempty"`
	TotalInPage int64 `json:"total_in_page,omitempty"`
	LastPage    int64 `json:"last_page,omitempty"`
}

func OrderDTO(order *model.Order) *Order {
	return &Order{
		ConsignmentId:    order.ConsignmentId,
		OrderCreatedAt:   order.CreatedAt,
		OrderDescription: order.ItemDescription,
		MerchantOrderId:  order.MerchantOrderId,
		RecipientName:    order.RecipientName,
		RecipientAddress: order.RecipientAddress,
		RecipientPhone:   order.RecipientPhone,
		OrderAmount:      order.AmountToCollect,
		TotalFee:         order.DeliveryFee,
		Instruction:      order.SpecialInstruction,
		OrderTypeId:      order.ItemType,
		CodFee:           order.CODFee,
		PromoDiscount:    0,
		Discount:         0,
		DeliveryFee:      order.DeliveryFee,
		OrderStatus:      order.OrderStatus,
		OrderType:        "Delivery",
		ItemType:         "Parcel",
	}
}
