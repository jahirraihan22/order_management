package _interface

import (
	"github.com/labstack/echo/v4"
	"order_management/app/http/controller"
)

type CrudService interface {
	Create(ctx echo.Context) error
	Get(ctx echo.Context) error
}
type Request interface {
	Validate() map[string][]string
	Bind() error
}

func OrderManager() *controller.OrderController {
	return &controller.OrderController{}
}

func AuthManager() *controller.AuthController {
	return &controller.AuthController{}
}
