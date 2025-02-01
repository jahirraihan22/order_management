package service

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"order_management/app/http/response"
	"order_management/app/model"
	"order_management/config"
	"order_management/utility"
	"strings"
	"time"
)

type JWTService struct{}

func (js JWTService) CreateJwtTokens(jwtPayload model.JwtPayload) (string, string, error) {
	var accessTokenPayload = jwt.MapClaims{
		"Username":  jwtPayload.Username,
		"Phone":     jwtPayload.Phone,
		"ExpiresAt": utility.GetCurrentTimeInDefaultTimezone().Add(config.TokenTTL * time.Second),
	}
	jwtAccessToken, jwtError := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenPayload).SignedString([]byte(config.JwtSecretKey))
	if jwtError != nil {
		response.LogMessage("ERROR", "jwt generation failed", jwtError)
	}

	var refreshTokenPayload = jwt.MapClaims{
		"Username":  jwtPayload.Username,
		"Phone":     jwtPayload.Phone,
		"ExpiresAt": utility.GetCurrentTimeInDefaultTimezone().Add(config.RefreshTokenTTL * time.Second),
	}
	jwtRefreshToken, jwtError := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenPayload).SignedString([]byte(config.JwtSecretKey))
	if jwtError != nil {
		response.LogMessage("ERROR", "jwt generation failed", jwtError)
	}
	return jwtAccessToken, jwtRefreshToken, nil
}

func (js JWTService) ParseJWTAndSetupInfoInHttpRequest(echoContext echo.Context) bool {
	var token string
	if len(strings.Split(echoContext.Request().Header.Get(echo.HeaderAuthorization), " ")) == 2 {
		token = strings.Split(echoContext.Request().Header.Get(echo.HeaderAuthorization), " ")[1]
	} else {
		response.LogMessage("ERROR", "invalid token", nil)
		return false
	}
	claims, err := js.ParseTokenAndRetrieveClaims(token)
	if err != nil {
		response.LogMessage("ERROR", "parsing JWT token", err)
		return false
	}
	jwtPayload := js.ParseJwtPayloadFromClaims(claims)
	var expiredTime = claims["ExpiresAt"].(string)

	var currentTime = utility.GetCurrentTimeInDefaultTimezone()
	var expirationTime, _ = time.Parse(time.RFC3339Nano, expiredTime)

	expirationTimeIsExpired := currentTime.Before(expirationTime)
	if !expirationTimeIsExpired {
		return false
	}

	echoContext.Request().Header.Add("phone", jwtPayload.Phone)
	echoContext.Request().Header.Add("username", jwtPayload.Username)
	return true
}

func (js JWTService) ParseJwtPayloadFromClaims(claims jwt.MapClaims) model.JwtPayload {
	return model.JwtPayload{
		Username: claims["Username"].(string),
		Phone:    claims["Phone"].(string),
	}
}

func (js JWTService) ParseTokenAndRetrieveClaims(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JwtSecretKey), nil
	})
	if err != nil {
		response.LogMessage("WARNING", "token is not valid", err)
		return nil, echo.NewHTTPError(http.StatusForbidden, "token not valid")
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	return claims, nil
}
func (js JWTService) GetPayloadFromClaims(ctx echo.Context) *model.JwtPayload {
	return &model.JwtPayload{
		Username: ctx.Request().Header.Get("username"),
		Phone:    ctx.Request().Header.Get("phone"),
	}
}
