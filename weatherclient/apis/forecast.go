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
Function Calls the weather API to get Forecast using the paramater q  which can either be
lattitude,longitude value in the format "lat,lang" or it can be a single zip code "zipCode"
and days in string format
*/
func GetForecast(q, days string) (*weatherdata.ForeCastData, error) {
	var apiKey = os.Getenv("API_KEY")
	var ForecastData *weatherdata.ForeCastData
	loggers.Logger.Info("Initiated API call for GetForecast")
	weatherClientResource, err := weatherclient.GetWeatherClientResource()
	if err != nil {
		loggers.Logger.WithField("Key", logrus.Fields{
			"Status":  500,
			"isFatal": true,
		}).Warn("Error Loading Weather Client Resource")
		return ForecastData, err
	} else {
		var url = "forecast.json"
		response, err := weatherClientResource.GetDataFromClient(url, map[string]string{
			"q":    q,
			"key":  apiKey,
			"days": days,
		})

		if err != nil {
			loggers.Logger.WithField("Key", logrus.Fields{
				"Status":  500,
				"isFatal": false,
			}).Warn("Error Fetching Info from client")
			return ForecastData, err
		} else {
			if err := json.Unmarshal(response, &ForecastData); err != nil {
				loggers.Logger.WithField("Key", logrus.Fields{
					"Status":  500,
					"isFatal": false,
				}).Warn("Error Unmarshalling ForecastData ")
				return ForecastData, err
			}
			loggers.Logger.Info("GetForecast Data Generated response")
			return ForecastData, nil
		}
	}
}
