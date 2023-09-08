package main

import (
	"fmt"
	"real-time-weather-app/weatherclient"
)

func main() {
	fmt.Println("Initialized the Application")
	weatherClientResource, err := weatherclient.NewWeatherClientResource()
	if err != nil {
		// Fact that the client got created
		fmt.Println(err)
	} else {
		// Fact that the client got created
		requestParams := map[string]string{
			"q":   "48.8567,2.3508",
			"key": "252ecab989574b60a0d194244230809",
		}
		response, err := weatherClientResource.GetDataFromClient("current.json", requestParams)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(response))
		}
	}
}
