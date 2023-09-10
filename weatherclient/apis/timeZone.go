package apis

import (
	"fmt"
	"os"
	"real-time-weather-app/weatherclient"
)

/*
Add description here
*/
func GetTimeZone(q string) ([]byte, error) {
	var apiKey = os.Getenv("API_KEY")
	fmt.Println("Initiated API call for GetTimeZone")
	weatherClientResource, err := weatherclient.GetWeatherClientResource()
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		var url = "timezone.json"
		response, err := weatherClientResource.GetDataFromClient(url, map[string]string{
			"q":   q,
			"key": apiKey,
		})

		if err != nil {
			fmt.Println(err)
			return nil, err
		} else {
			fmt.Println("GetTimeZone Data Generated response")
			return response, nil
		}
	}
}
