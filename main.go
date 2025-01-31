package main

import (
	"order_management/app/http/response"
	"order_management/config"
	"order_management/database"
	"order_management/server"
)

func main() {
	response.LogMessage("INFO", "initialize configuration", nil)
	config.Init()
	response.LogMessage("INFO", "initialize database", nil)
	database.Init()
	response.LogMessage("INFO", "migrating database", nil)
	database.Migrate()
	response.LogMessage("INFO", "seeding necessary data", nil)
	database.Seed()
	response.LogMessage("INFO", "starting server", nil)
	server.Init()
}
