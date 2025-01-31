package route

import (
	"github.com/labstack/echo/v4"
	"net/http"
	_interface "order_management/app/interface"
)

func Init(e *echo.Echo) {
	e.GET("/health", func(ctx echo.Context) error { return ctx.JSON(http.StatusOK, "I am alive") })
	orderRoute(e.Group("/api/v1/orders"))
	e.POST("/api/v1/login", _interface.AuthManager().Login)
}

func orderRoute(g *echo.Group) {
	g.GET("/all", _interface.OrderManager().Get)
	g.POST("", _interface.OrderManager().Create)
}
