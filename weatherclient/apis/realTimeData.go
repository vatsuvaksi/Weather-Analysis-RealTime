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
*/
func GetRealTimeData(q string) (*weatherdata.RealTimeData, error) {
	var apiKey = os.Getenv("API_KEY")
	var realTimeData *weatherdata.RealTimeData
	fmt.Println("Initiated API call for GetRealTimeData")
	weatherClientResource, err := weatherclient.GetWeatherClientResource()
	if err != nil {
		fmt.Println(err)
		return realTimeData, err
	} else {
		var url = "current.json"
		response, err := weatherClientResource.GetDataFromClient(url, map[string]string{
			"q":   q,
			"key": apiKey,
		})

		if err != nil {
			fmt.Println(err)
			return realTimeData, err
		} else {
			if err := json.Unmarshal(response, &realTimeData); err != nil {
				fmt.Println(err)
				return realTimeData, err
			} else {
				fmt.Println("Get Real Time Data Generated response")
				return realTimeData, nil
			}
		}
	}
}
