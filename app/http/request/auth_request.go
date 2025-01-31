package request

import (
	"github.com/labstack/echo/v4"
	"order_management/app/http/response"
	"order_management/app/model"
)

type Auth struct {
	CTX          echo.Context
	userLoginDTO *model.UserLoginDTO
}

func (a *Auth) Bind() error {
	a.userLoginDTO = new(model.UserLoginDTO)
	err := a.CTX.Bind(&a.userLoginDTO)
	if err != nil {
		response.LogMessage("Error", "binding instance", err)
		return err
	}
	return nil
}

func (a *Auth) Validate() map[string][]string {
	err := a.CTX.Validate(a.userLoginDTO)
	if err != nil {
		return ValidationMsg(err)
	}
	return nil
}
func (a *Auth) GetLoginDTO() *model.UserLoginDTO {
	return a.userLoginDTO
}
