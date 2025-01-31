package test

import (
	"github.com/stretchr/testify/assert"
	"order_management/app/model"
	"order_management/app/service"
	"testing"
)

func TestCalculatePrice(t *testing.T) {
	orderService := &service.OrderService{}

	testCases := []struct {
		name          string
		order         *model.Order
		expectedPrice float64
	}{
		{
			name:          "City is 1, ItemWeight <= 0.5",
			order:         &model.Order{RecipientCity: 1, ItemWeight: 0.4},
			expectedPrice: 60.0,
		},
		{
			name:          "City is 1, ItemWeight between 0.5 and 1",
			order:         &model.Order{RecipientCity: 1, ItemWeight: 0.8},
			expectedPrice: 70.0,
		},
		{
			name:          "City is 1, ItemWeight > 1",
			order:         &model.Order{RecipientCity: 1, ItemWeight: 2},
			expectedPrice: 75.0,
		},
		{
			name:          "City is not 1, ItemWeight <= 0.5",
			order:         &model.Order{RecipientCity: 2, ItemWeight: 0.4},
			expectedPrice: 100.0,
		},
		{
			name:          "City is not 1, ItemWeight between 0.5 and 1",
			order:         &model.Order{RecipientCity: 2, ItemWeight: 0.8},
			expectedPrice: 110.0,
		},
		{
			name:          "City is not 1, ItemWeight > 1",
			order:         &model.Order{RecipientCity: 2, ItemWeight: 2},
			expectedPrice: 115.0,
		},
		{
			name:          "City is not 1, ItemWeight > 1",
			order:         &model.Order{RecipientCity: 2, ItemWeight: 5},
			expectedPrice: 160.0,
		},
		{
			name:          "City is not 1, ItemWeight > 1",
			order:         &model.Order{RecipientCity: 1, ItemWeight: 20},
			expectedPrice: 345.0,
		},
		{
			name:          "City is invalid",
			order:         &model.Order{RecipientCity: 0, ItemWeight: 20},
			expectedPrice: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := orderService.CalculatePrice(tc.order)
			assert.Equal(t, tc.expectedPrice, result.DeliveryFee)
		})
	}
}
