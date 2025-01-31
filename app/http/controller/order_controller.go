package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"math"
	"net/http"
	"order_management/app/http/request"
	"order_management/app/http/response"
	"order_management/app/model"
	"order_management/app/service"
	"order_management/database"
	"strconv"
)

type OrderController struct {
	orderService service.OrderService
	js           service.JWTService
}

func (oc *OrderController) Create(ctx echo.Context) error {
	req := request.Order{CTX: ctx}
	err := req.Bind()
	if err != nil {
		return response.ErrorResponse(req.CTX, "Invalid input", err, http.StatusUnprocessableEntity)
	}

	validationErr := req.Validate()
	if validationErr != nil {
		return response.ErrorResponse(req.CTX, "Order request is invalid", validationErr, http.StatusUnprocessableEntity)
	}
	order := req.OrderReqDataToObject()
	oc.orderService.CalculatePrice(order)
	userInfo := oc.js.GetPayloadFromClaims(ctx)
	var user = new(model.Users)
	userResult := database.Client.First(&user, "email = ?", userInfo.Username)
	if userResult.Error != nil {
		if errors.Is(userResult.Error, gorm.ErrRecordNotFound) {
			return response.ErrorResponse(req.CTX, "User not found", nil, http.StatusUnprocessableEntity)
		}
	}
	order.UserID = user.ID
	result := database.Client.Create(&order)
	if result.Error != nil {
		response.LogMessage("ERROR", "Order placed failed", result.Error)
		return response.ErrorResponse(req.CTX, "Order placed failed", errors.New("something went wrong"), http.StatusUnprocessableEntity)
	}

	data := response.Order{
		ConsignmentId:   order.ConsignmentId,
		MerchantOrderId: order.MerchantOrderId,
		OrderStatus:     order.OrderStatus,
		DeliveryFee:     order.DeliveryFee,
		OrderCreatedAt:  order.CreatedAt,
	}

	return response.SuccessResponse(req.CTX, "Order successfully placed", data)
}

func (oc *OrderController) Get(ctx echo.Context) error {
	userInfo := oc.js.GetPayloadFromClaims(ctx)
	var user = new(model.Users)
	userResult := database.Client.First(&user, "email = ?", userInfo.Username)
	if userResult.Error != nil {
		if errors.Is(userResult.Error, gorm.ErrRecordNotFound) {
			return response.ErrorResponse(ctx, "User not found", nil, http.StatusUnprocessableEntity)
		}
	}
	var orders []model.Order
	var orderResponse response.PaginatedOrderResponse
	var result *gorm.DB
	// if paginate enable
	if ctx.QueryParam("paginate") == "true" {
		var totalRecords int64
		database.Client.Model(&model.Order{}).Where("user_id = ?", user.ID).Count(&totalRecords)
		page, _ := strconv.Atoi(ctx.QueryParam("page"))
		if page <= 0 {
			page = 1
		}
		limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
		if limit <= 0 {
			limit = 10
		}
		offset := (page - 1) * limit

		result = database.Client.Where("user_id = ?", user.ID).Offset(offset).Limit(limit).Find(&orders)
		lastPage := int64(math.Ceil(float64(totalRecords) / float64(limit)))
		orderResponse.Paginate = response.Paginate{
			Total:       totalRecords,
			CurrentPage: int64(page),
			PerPage:     int64(limit),
			TotalInPage: int64(len(orders)),
			LastPage:    lastPage,
		}
	} else {
		result = database.Client.Where("user_id = ?", user.ID).Find(&orders)
	}
	if result.Error != nil {
		response.LogMessage("ERROR", "something wrong", result.Error)
	}
	for _, order := range orders {
		orderResponse.Orders = append(orderResponse.Orders, response.OrderDTO(&order))
	}
	return response.SuccessResponse(ctx, "successfully found", orderResponse)
}

func NewOrderController(orderService service.OrderService, js service.JWTService) *OrderController {
	return &OrderController{
		orderService,
		js,
	}
}
