package response

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Response struct {
	Type    string      `json:"type"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func SuccessResponse(ctx echo.Context, message string, data interface{}) error {
	return ctx.JSON(http.StatusOK, Response{
		Type:    "success",
		Message: message,
		Data:    data,
		Code:    http.StatusOK,
	})
}

func ErrorResponse(ctx echo.Context, message string, error interface{}, code int) error {
	return ctx.JSON(http.StatusOK, Response{
		Type:    "error",
		Message: message,
		Code:    code,
		Error:   error,
	})
}

func LogMessage(logType string, message string, err error) {
	if err != nil {
		fmt.Printf("[%s]: %s => %v\n", logType, message, err.Error())
	} else {
		fmt.Printf("[%s]: %s \n", logType, message)
	}
}
