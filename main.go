package main

import (
	"fmt"
	// "real-time-weather-app/weatherclient"
)

func main() {
	fmt.Println("Initialized the Application")
	/*   REMOVE THIS CODE AS SOON AS POSSIBLE

	weatherClientResource, err := weatherclient.NewWeatherClientResource()
	if err != nil {
		// Fact that the client got created
		fmt.Println(err)
	} else {
		// Fact that the client got created
		requestParams := map[string]string{
			"q":   "48.8567,2.3508",
			"key": "NEEDS TO BE FETCHED FROM .env file",
		}
		response, err := weatherClientResource.GetDataFromClient("current.json", requestParams)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(response))
		}
	}
	*/
}
