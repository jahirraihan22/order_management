package model

import "time"

type Users struct {
	ID        uint      `json:"id,omitempty" gorm:"primaryKey" `
	Name      string    `json:"name" gorm:"not null" `
	Email     string    `json:"email" gorm:"unique" `
	Phone     string    `json:"phone" gorm:"unique" `
	Password  string    `json:"password,omitempty" gorm:"not null" `
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}

type UserLoginDTO struct {
	Username string `json:"username" validate:"required" `
	Password string `json:"password" validate:"required" `
}
