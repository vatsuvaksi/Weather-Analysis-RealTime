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
Function Calls the weather API to get real time Data using the paramater q which can either be
lattitude,longitude value in the format "lat,lang" or it can be a single zip code "zipCode"
*/
func GetRealTimeData(q string) (*weatherdata.RealTimeData, error) {
	var apiKey = os.Getenv("API_KEY")
	var realTimeData *weatherdata.RealTimeData
	loggers.Logger.Info("Initiated API call for GetRealTimeData")
	weatherClientResource, err := weatherclient.GetWeatherClientResource()
	if err != nil {
		loggers.Logger.Info(err)
		return realTimeData, err
	} else {
		var url = "current.json"
		response, err := weatherClientResource.GetDataFromClient(url, map[string]string{
			"q":   q,
			"key": apiKey,
		})

		if err != nil {
			loggers.Logger.WithField("Key", logrus.Fields{
				"Status":  500,
				"isFatal": false,
			}).Warn("Error Fetching Info from client for realTimeData API ")
			return realTimeData, err
		} else {
			if err := json.Unmarshal(response, &realTimeData); err != nil {
				loggers.Logger.WithField("Key", logrus.Fields{
					"Status":  500,
					"isFatal": false,
				}).Warn("Error Unmarshalling RealTime Data ")
				return realTimeData, err
			} else {
				loggers.Logger.Info("Get Real Time Data Generated response")
				return realTimeData, nil
			}
		}
	}
}
