package request

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"order_management/app/http/response"
	"order_management/app/model"
	"time"
)

type Order struct {
	CTX      echo.Context
	orderDTO *model.OrderDTO
}

func (req *Order) Bind() error {
	req.orderDTO = new(model.OrderDTO)
	err := req.CTX.Bind(&req.orderDTO)
	if err != nil {
		response.LogMessage("Error", "binding instance", err)
		return err
	}
	return nil
}

func (req *Order) Validate() map[string][]string {
	err := req.CTX.Validate(req.orderDTO)
	if err != nil {
		return ValidationMsg(err)
	}
	return nil
}

func (req *Order) OrderReqDataToObject() *model.Order {
	return &model.Order{
		StoreId:            req.orderDTO.StoreId,
		MerchantOrderId:    req.orderDTO.MerchantOrderId,
		RecipientName:      req.orderDTO.RecipientName,
		RecipientPhone:     req.orderDTO.RecipientPhone,
		RecipientAddress:   req.orderDTO.RecipientAddress,
		RecipientCity:      req.orderDTO.RecipientCity,
		RecipientZone:      req.orderDTO.RecipientZone,
		RecipientArea:      req.orderDTO.RecipientArea,
		DeliveryType:       req.orderDTO.DeliveryType,
		ItemType:           req.orderDTO.ItemType,
		SpecialInstruction: req.orderDTO.SpecialInstruction,
		ItemQuantity:       req.orderDTO.ItemQuantity,
		ItemWeight:         req.orderDTO.ItemWeight,
		AmountToCollect:    req.orderDTO.AmountToCollect,
		ItemDescription:    req.orderDTO.ItemDescription,
		ConsignmentId:      uuid.New().String(),
		OrderStatus:        "Pending",
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}
}
