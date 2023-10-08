package apis

import (
	"encoding/json"
	"os"
	weatherdata "real-time-weather-app/models/network/weatherData"
	"real-time-weather-app/utils/loggers"
	"real-time-weather-app/weatherclient"

	"github.com/sirupsen/logrus"
)

/*
Add description here
*/
func GetTimeZone(q string) (*weatherdata.TimeZoneData, error) {
	var apiKey = os.Getenv("API_KEY")
	var location *weatherdata.TimeZoneData
	loggers.Logger.Info("Initiated API call for GetTimeZone")
	weatherClientResource, err := weatherclient.GetWeatherClientResource()
	if err != nil {
		loggers.Logger.WithField("Key", logrus.Fields{
			"Status":  500,
			"isFatal": false,
		}).Warn("Error Fetching weather client resource ")
		return location, err
	} else {
		var url = "timezone.json"
		response, err := weatherClientResource.GetDataFromClient(url, map[string]string{
			"q":   q,
			"key": apiKey,
		})

		if err != nil {
			loggers.Logger.WithField("Key", logrus.Fields{
				"Status":  500,
				"isFatal": false,
			}).Warn("Error - Fetching Info from client")
			return location, err
		} else {
			if err := json.Unmarshal(response, &location); err != nil {
				loggers.Logger.WithField("Key", logrus.Fields{
					"Status":  500,
					"isFatal": false,
				}).Warn("Error Unmarshalling Location Struct ")
				return location, err
			} else {
				loggers.Logger.Info("GetTimeZone Data Generated response")
				return location, nil
			}
		}
	}
}
