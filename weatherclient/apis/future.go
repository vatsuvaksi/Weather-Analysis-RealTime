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
and date in
*/
func GetFuture(q, dt string) (*weatherdata.FutureData, error) {
	var apiKey = os.Getenv("API_KEY")
	loggers.Logger.Info("Initiated API call for FutureDataData")
	var FutureData *weatherdata.FutureData
	weatherClientResource, err := weatherclient.GetWeatherClientResource()
	if err != nil {
		loggers.Logger.WithField("Key", logrus.Fields{
			"Status":  500,
			"isFatal": false,
		}).Warn("Error Fetching weather client resource for Future API ")
		return FutureData, err
	} else {
		var url = "future.json"
		response, err := weatherClientResource.GetDataFromClient(url, map[string]string{
			"q":   q,
			"key": apiKey,
			"dt":  dt,
		})

		if err != nil {
			loggers.Logger.WithField("Key", logrus.Fields{
				"Status":  500,
				"isFatal": false,
			}).Warn("Error Fetching Info from Client for Future API ")
			return nil, err
		} else {

			if err := json.Unmarshal(response, &FutureData); err != nil {
				// loggers.Logger.Info("I am here")
				loggers.Logger.WithField("Key", logrus.Fields{
					"Status":  500,
					"isFatal": false,
				}).Warn("Error Unmarshalling FutureData ")
				return FutureData, err
			} else {
				loggers.Logger.Info("GetFuture Data Generated response")
				return FutureData, nil
			}
		}
	}
}
