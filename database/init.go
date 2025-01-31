package database

import (
	"log"
	"order_management/config"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Client *gorm.DB
	once   sync.Once
)

func Init() *gorm.DB {
	once.Do(func() {
		conn, err := gorm.Open(mysql.Open(config.ConnectionString), &gorm.Config{})
		if err != nil {
			log.Fatalf("DB CONNECTION FAILED: %v", err)
		}

		Client = conn
	})

	return Client
}
