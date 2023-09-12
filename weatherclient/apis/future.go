package apis

import (
	"encoding/json"
	"fmt"
	"os"
	weatherdata "real-time-weather-app/models/network/weatherData"
	"real-time-weather-app/weatherclient"
)

/*
Function Calls the weather API to get real time Data using the paramater q which can either be
lattitude,longitude value in the format "lat,lang" or it can be a single zip code "zipCode"
and date in
*/
func GetFuture(q, dt string) (*weatherdata.ForeCastData, error) {
	var apiKey = os.Getenv("API_KEY")
	fmt.Println("Initiated API call for ForeCastDataData")
	var ForeCastData *weatherdata.ForeCastData
	weatherClientResource, err := weatherclient.GetWeatherClientResource()
	if err != nil {
		fmt.Println(err)
		return ForeCastData, err
	} else {
		var url = "future.json"
		response, err := weatherClientResource.GetDataFromClient(url, map[string]string{
			"q":   q,
			"key": apiKey,
			"dt":  dt,
		})

		if err != nil {
			fmt.Println(err)
			return nil, err
		} else {

			if err := json.Unmarshal(response, &ForeCastData); err != nil {
				// fmt.Println("I am here")
				return ForeCastData, err
			} else {
				fmt.Println("GetFuture Data Generated response")
				return ForeCastData, nil
			}
		}
	}
}
