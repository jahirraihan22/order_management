package _middleware

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"order_management/app/http/response"
	"order_management/app/service"
	"order_management/config"
	"strings"
)

func Init(server *echo.Echo) *echo.Echo {
	server.Use(middleware.Logger())
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"*"},
	}))
	server.Use(echojwt.WithConfig(echojwt.Config{
		Skipper: func(c echo.Context) bool {
			method := c.Request().Method
			path := c.Request().RequestURI
			return pathSkipChecker(path, method)
		},
		SigningKey: []byte(config.JwtSecretKey),
		ErrorHandler: func(c echo.Context, err error) error {
			return response.ErrorResponse(c, "Unauthorized", nil, http.StatusUnauthorized)
		},
		SuccessHandler: func(c echo.Context) {
			service.JWTService{}.ParseJWTAndSetupInfoInHttpRequest(c)
		},
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwt.MapClaims)
		},
	}))
	return server
}

func pathSkipChecker(path string, method string) bool {
	if strings.Contains(path, "login") || strings.Contains(path, "health") {
		return true
	} else {
		return false
	}
}
