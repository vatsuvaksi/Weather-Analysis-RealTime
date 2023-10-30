package apis

import (
	"encoding/json"
	"os"
	weatherdata "real-time-weather-app/models/network/weatherData"
	"real-time-weather-app/utils/loggers"
	"real-time-weather-app/weatherclient"

	"github.com/sirupsen/logrus"
)

type ConcurrentFetchForecastData struct {
	Data *weatherdata.ForeCastData
	Err  error
}

/*
Function Calls the weather API to get Forecast using the paramater q  which can either be
lattitude,longitude value in the format "lat,lang" or it can be a single zip code "zipCode"
and days in string format
*/
func GetForecast(q, days string) (*weatherdata.ForeCastData, error) {

	var resultCh = make(chan ConcurrentFetchForecastData)
	// Go Routine for execution
	// Launch a goroutine to fetch data
	go func() {
		var apiKey = os.Getenv("API_KEY")
		var forecastData *weatherdata.ForeCastData
		loggers.Logger.Info("Initiated API call for Get ForeCast Data")
		weatherClientResource, err := weatherclient.GetWeatherClientResource()
		if err != nil {
			loggers.Logger.Info(err)
			resultCh <- ConcurrentFetchForecastData{Data: nil, Err: err}
			return
		}
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
			}).Warn("Error Fetching Info from client for forecast API ")
			resultCh <- ConcurrentFetchForecastData{Data: nil, Err: err}
			return
		}
		if err := json.Unmarshal(response, &forecastData); err != nil {
			loggers.Logger.WithField("Key", logrus.Fields{
				"Status":  500,
				"isFatal": false,
			}).Warn("Error Unmarshalling forecastData Data ")
			resultCh <- ConcurrentFetchForecastData{Data: nil, Err: err}
			return
		}
		loggers.Logger.Info("GetForecast Data Generated response")
		resultCh <- ConcurrentFetchForecastData{Data: forecastData, Err: nil}
	}()
	// Wait for the result from the goroutine
	result := <-resultCh
	return result.Data, result.Err
}
