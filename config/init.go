package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

var (
	DbServer         = ""
	DbPort           = ""
	DbName           = ""
	DbUsername       = ""
	DbPassword       = ""
	ConnectionString = ""
	AppMode          = ""
)

func Init() {
	AppMode = strings.ToLower(os.Getenv("APP_MODE"))
	if AppMode == "" {
		AppMode = "develop"
	}
	if AppMode != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Println("[ERROR]: getting env variable", err.Error())
			return
		}
	}
	DbServer = os.Getenv("DB_SERVER")
	DbPort = os.Getenv("DB_PORT")
	DbUsername = os.Getenv("DB_USERNAME")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbName = os.Getenv("DB_NAME")
	ConnectionString = DbUsername + ":" + DbPassword + "@tcp(" + DbServer + ":" + DbPort + ")/" + DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
}
