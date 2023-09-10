package apis

import (
	"fmt"
	"os"
	"real-time-weather-app/weatherclient"
)

/*
Function Calls the weather API to get real time Data using the paramater q which can either be
lattitude,longitude value in the format "lat,lang" or it can be a single zip code "zipCode"
and date in
*/
func GetFuture(q, dt string) ([]byte, error) {
	var apiKey = os.Getenv("API_KEY")
	fmt.Println("Initiated API call for GetFuture")
	weatherClientResource, err := weatherclient.GetWeatherClientResource()
	if err != nil {
		fmt.Println(err)
		return nil, err
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
			fmt.Println("GetFuture Data Generated response")
			return response, nil
		}
	}
}
