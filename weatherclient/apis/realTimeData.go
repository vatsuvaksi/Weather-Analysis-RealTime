package apis

import (
	"encoding/json"
	"os"
	weatherdata "real-time-weather-app/models/network/weatherData"
	"real-time-weather-app/utils/loggers"
	"real-time-weather-app/weatherclient"

	"github.com/sirupsen/logrus"
)

type ConcurrentFetchData struct {
	Data *weatherdata.RealTimeData
	Err  error
}

/*
Function Calls the weather API to get real time Data using the paramater q which can either be
lattitude,longitude value in the format "lat,lang" or it can be a single zip code "zipCode"
*/
func GetRealTimeData(q string) (*weatherdata.RealTimeData, error) {

	// Create a channel for results
	resultCh := make(chan ConcurrentFetchData)

	// Launch a goroutine to fetch data
	go func() {
		var apiKey = os.Getenv("API_KEY")
		var realTimeData *weatherdata.RealTimeData
		loggers.Logger.Info("Initiated API call for GetRealTimeData")
		weatherClientResource, err := weatherclient.GetWeatherClientResource()
		if err != nil {
			loggers.Logger.Info(err)
			resultCh <- ConcurrentFetchData{Data: nil, Err: err}
			return
		}
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
			resultCh <- ConcurrentFetchData{Data: nil, Err: err}
			return
		}
		if err := json.Unmarshal(response, &realTimeData); err != nil {
			loggers.Logger.WithField("Key", logrus.Fields{
				"Status":  500,
				"isFatal": false,
			}).Warn("Error Unmarshalling RealTime Data ")
			resultCh <- ConcurrentFetchData{Data: nil, Err: err}
			return
		}
		loggers.Logger.Info("Get Real Time Data Generated response")
		resultCh <- ConcurrentFetchData{Data: realTimeData, Err: nil}
	}()
	// Wait for the result from the goroutine
	result := <-resultCh
	return result.Data, result.Err
}
