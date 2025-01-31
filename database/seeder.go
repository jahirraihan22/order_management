package database

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"order_management/app/http/response"
	"order_management/app/model"
	"time"
)

type Seeder interface{}

func Seed() {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("321dsa"), bcrypt.DefaultCost) // default password "321dsa"

	users := []*model.Users{
		{
			Name:      "Mr ABC",
			Email:     "abc@mailinator.com",
			Phone:     "01901901902",
			Password:  string(hashedPassword),
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
		{
			Name:      "Mr XYZ",
			Email:     "01901901901@mailinator.com",
			Phone:     "01901901901",
			Password:  string(hashedPassword),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	err := userSeeder(users)
	if err != nil {
		response.LogMessage("WARNING", "user seeding ignored", err)
	}
}

func userSeeder(users []*model.Users) error {
	for _, user := range users {
		if Client.First(&model.Users{}, "email = ?", user.Email).Error == nil {
			response.LogMessage("WARNING", "user seeding ignored", errors.New("user email already exists"))
		}
		if Client.First(&model.Users{}, "phone = ?", user.Phone).Error == nil {
			response.LogMessage("WARNING", "user seeding ignored", errors.New("user phone already exists"))
		}
	}
	result := Client.Create(users)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
