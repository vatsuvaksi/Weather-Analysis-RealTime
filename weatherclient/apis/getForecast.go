package apis

import (
	"fmt"
	"os"
	"real-time-weather-app/weatherclient"
)

/*
Function Calls the weather API to get real time Data using the paramater q which can either be
lattitude,longitude value in the format "lat,lang" or it can be a single zip code "zipCode"
*/
func GetForecast(q, days string) ([]byte, error) {
	var apiKey = os.Getenv("API_KEY")
	fmt.Println("Initiated API call for GetForecast")
	weatherClientResource, err := weatherclient.GetWeatherClientResource()
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		var url = "forecast.json"
		response, err := weatherClientResource.GetDataFromClient(url, map[string]string{
			"q":    q,
			"key":  apiKey,
			"days": days,
		})

		if err != nil {
			fmt.Println(err)
			return nil, err
		} else {
			fmt.Println("GetForecast Data Generated response")
			return response, nil
		}
	}
}
