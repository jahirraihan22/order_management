package utility

import (
	"order_management/app/http/response"
	"order_management/config"
	"time"
)

func GetCurrentTimeInDefaultTimezone() time.Time {
	location, err := time.LoadLocation(config.TimeZone)
	if err != nil {
		response.LogMessage("Error", "while getting current time from timezone", err)
	}
	return time.Now().In(location)
}
