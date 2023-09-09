package main

import (
	"fmt"
	"real-time-weather-app/weatherclient"

	"github.com/joho/godotenv"
)

func InitializeWeatherApplication() {
	fmt.Println("InitializeWeatherApplication Started ...")

	// Load environment variables from the .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		return
	} else {
		fmt.Println("Environment Variables Loaded ...")
	}

	// Initialize WeatherClient
	_, err := weatherclient.NewWeatherClientResource()
	if err != nil {
		// Fact that the client got created = No Error Logs
		fmt.Println(err)
	}
	fmt.Println("InitializeWeatherApplication Completed ...")
}
