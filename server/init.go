package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"order_management/app/http/request"
	"order_management/config"
	"order_management/route"
	_middleware "order_management/route/middleware"
)

func Init() {
	e := echo.New()
	e = _middleware.Init(e)
	e.Validator = &request.Validator{Validator: validator.New()}
	route.Init(e)
	e.Logger.Fatal(e.Start(config.ServerPort))
}
