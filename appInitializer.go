package main

import (
	"fmt"
	"log"
	"real-time-weather-app/server"
	"real-time-weather-app/weatherclient"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type cfg struct {
	ServerConfig ServerConfig `mapstructure:"ServerConfig"`
}

type ServerConfig struct {
	Port string `mapstructure:"Port"`
}

var config cfg

func initializeWeatherApplication() {
	fmt.Println("InitializeWeatherApplication Started ...")

	// Load environment variables from the .env file
	loadEnvVariables()
	//Initialize go configurations
	loadAppConfig()

	// Initialize WeatherClient
	loadWeatherClientResource()

	initializeServer()
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
func initializeServer() {
	server.Start(config.ServerConfig.Port)
}
func loadAppConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("unmarhsalling err", err)
	}
	fmt.Println("Configurations loaded successfully")
}
