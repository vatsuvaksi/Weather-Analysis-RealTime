package main

import (
	"fmt"
	"real-time-weather-app/weatherclient"
)

func main() {
	fmt.Println("Initialized the Application")
	weatherClient, err := weatherclient.NewWeatherClientResource()
	if err != nil {
		// Fact that the client got created
		fmt.Println(err)
	} else {
		// Fact that the client got created
		fmt.Println(&weatherClient)
	}
}
