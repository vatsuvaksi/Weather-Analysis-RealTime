package apis

import (
	"encoding/json"
	"fmt"
	"os"
	weatherdata "real-time-weather-app/models/network/weatherData"
	"real-time-weather-app/weatherclient"
)

/*
Add description here
*/
func GetTimeZone(q string) (*weatherdata.TimeZoneData, error) {
	var apiKey = os.Getenv("API_KEY")
	var location *weatherdata.TimeZoneData
	fmt.Println("Initiated API call for GetTimeZone")
	weatherClientResource, err := weatherclient.GetWeatherClientResource()
	if err != nil {
		fmt.Println(err)
		return location, err
	} else {
		var url = "timezone.json"
		response, err := weatherClientResource.GetDataFromClient(url, map[string]string{
			"q":   q,
			"key": apiKey,
		})

		if err != nil {
			fmt.Println(err)
			return location, err
		} else {
			if err := json.Unmarshal(response, &location); err != nil {
				return location, err
			} else {
				fmt.Println("GetTimeZone Data Generated response")
				return location, nil
			}
		}
	}
}
