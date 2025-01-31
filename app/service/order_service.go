package service

import "order_management/app/model"

type OrderService struct{}

func (os *OrderService) CalculatePrice(order *model.Order) *model.Order {
	if order.RecipientCity <= 0 {
		return order
	}
	basePrice := 60.0
	feeThreshold := 0.
	weight := 1.0
	if order.RecipientCity != 1 {
		basePrice = 100
	}

	if order.ItemWeight <= 0.5 {
		feeThreshold = 0
	} else if order.ItemWeight > 0.5 && order.ItemWeight <= 1 {
		feeThreshold = 10
	} else {
		weight = order.ItemWeight - 1
		feeThreshold = 15
	}
	order.DeliveryFee = basePrice + (weight * feeThreshold)
	if order.AmountToCollect > 0 {
		order.CODFee = order.AmountToCollect * 0.01
	}
	return order
}

//func (os *OrderService) Paginate(order *model.Order) *model.Order {
//
//}
