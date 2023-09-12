package main

import (
	"fmt"
	"real-time-weather-app/weatherclient"

	"github.com/joho/godotenv"
)

func initializeWeatherApplication() {
	fmt.Println("InitializeWeatherApplication Started ...")

	// Load environment variables from the .env file
	loadEnvVariables()
	// Initialize WeatherClient
	loadWeatherClientResource()

	fmt.Println("InitializeWeatherApplication Completed ...")
}

func loadEnvVariables() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		return
	} else {
		fmt.Println("Environment Variables Loaded ...")
	}
}
func loadWeatherClientResource() {
	_, err := weatherclient.GetWeatherClientResource()
	if err != nil {
		// Fact that the client got created = No Error Logs
		fmt.Println(err)
	}
}
