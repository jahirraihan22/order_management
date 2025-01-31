package model

type JwtPayload struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
}
