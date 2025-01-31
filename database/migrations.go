package database

import (
	"order_management/app/http/response"
	"order_management/app/model"
)

// tables is the collections  of DB tables,
var tables = []interface{}{
	&model.Order{},
	&model.Users{},
}

// Migrate DB while booting up the app
func Migrate() {
	err := Client.AutoMigrate(tables...)
	if err != nil {
		response.LogMessage("ERROR", "Migration failed ", err)
		return
	}
	response.LogMessage("INFO", "Migration completed", nil)
}
