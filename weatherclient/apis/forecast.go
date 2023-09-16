package apis

import (
	"encoding/json"
	"fmt"
	"os"
	weatherdata "real-time-weather-app/models/network/weatherData"
	"real-time-weather-app/weatherclient"
)

/*
Function Calls the weather API to get Forecast using the paramater q  which can either be
lattitude,longitude value in the format "lat,lang" or it can be a single zip code "zipCode"
and days in string format
*/
func GetForecast(q, days string) (*weatherdata.ForeCastData, error) {
	var apiKey = os.Getenv("API_KEY")
	var ForecastData *weatherdata.ForeCastData
	fmt.Println("Initiated API call for GetForecast")
	weatherClientResource, err := weatherclient.GetWeatherClientResource()
	if err != nil {
		fmt.Println(err)
		return ForecastData, err
	} else {
		var url = "forecast.json"
		response, err := weatherClientResource.GetDataFromClient(url, map[string]string{
			"q":    q,
			"key":  apiKey,
			"days": days,
		})

		if err != nil {
			return ForecastData, err
		} else {
			if err := json.Unmarshal(response, &ForecastData); err != nil {
				return ForecastData, err
			}
			fmt.Println("GetForecast Data Generated response")
			return ForecastData, nil
		}
	}
}
