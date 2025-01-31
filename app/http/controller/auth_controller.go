package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"order_management/app/http/request"
	"order_management/app/http/response"
	"order_management/app/model"
	"order_management/app/service"
	"order_management/config"
	"order_management/database"
)

type AuthController struct {
	jwtService service.JWTService
}

func (a *AuthController) Login(ctx echo.Context) error {
	req := request.Auth{CTX: ctx}
	err := req.Bind()
	if err != nil {
		return response.ErrorResponse(req.CTX, "Invalid input", err, http.StatusUnprocessableEntity)
	}

	validationErr := req.Validate()
	if validationErr != nil {
		return response.ErrorResponse(req.CTX, "Invalid input", validationErr, http.StatusUnprocessableEntity)
	}
	var user = new(model.Users)
	result := database.Client.Model(&model.Users{}).Where("email = ?", req.GetLoginDTO().Username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			response.LogMessage("ERROR", "user not found", result.Error)
			return response.ErrorResponse(req.CTX, "Login error", result.Error, http.StatusUnauthorized)
		} else {
			response.LogMessage("ERROR", "something wrong", result.Error)
		}
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.GetLoginDTO().Password))
	if err != nil {
		response.LogMessage("ERROR", "wrong password", err)
		return response.ErrorResponse(req.CTX, "password does not matched", err, http.StatusUnauthorized)
	}
	accessToken, refreshToken, err := a.jwtService.CreateJwtTokens(model.JwtPayload{
		Username: user.Email,
		Phone:    user.Phone,
	})
	if err != nil {
		return err
	}
	token := response.Token{
		TokenType:    "Bearer",
		ExpiresIn:    config.TokenTTL,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return response.SuccessResponse(req.CTX, "login successful", token)
}

func NewAuthController(jwtService service.JWTService) *AuthController {
	return &AuthController{jwtService}
}
