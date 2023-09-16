package appIntialization

import (
	"fmt"
	"log"
	"real-time-weather-app/server"
	"real-time-weather-app/weatherclient"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
)

type ServerConfig struct {
	Port               string `mapstructure:"Port"`
	Name               string `mapstructure:"Name"`
	Network            string `mapstructure:"Network"`
	Address            string `mapstructure:"Address"`
	ReadBufferSize     int    `mapstructure:"ReadBufferSize"`
	MaxConnsPerIP      int    `mapstructure:"MaxConnsPerIP"`
	MaxRequestsPerConn int    `mapstructure:"MaxRequestsPerConn"`
	MaxRequestBodySize int    `mapstructure:"MaxRequestBodySize"`
	Concurrency        int    `mapstructure:"Concurrency"`
}

type WebConfig struct {
	ServerConfig ServerConfig `mapstructure:"ServerConfig"`
}

var config WebConfig

func InitializeWeatherApplication() {
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

// type webServer struct {
// 	Config WebConfig
// 	Log    log.Logger
// 	ln     net.Listener
// 	router *router.Router
// 	debug  bool
// }

func initializeServer() {
	router, err := server.MuxRouter()
	if err != nil {
		log.Fatal("Failed to initialize router", err)
	}

	s := &fasthttp.Server{
		Handler:            router.Handler,
		Name:               config.ServerConfig.Name,
		ReadBufferSize:     config.ServerConfig.ReadBufferSize,
		MaxConnsPerIP:      config.ServerConfig.MaxConnsPerIP,
		MaxRequestsPerConn: config.ServerConfig.MaxRequestsPerConn,
		MaxRequestBodySize: config.ServerConfig.MaxRequestBodySize, //  100 << 20, // 100MB // 1024 * 4, // MaxRequestBodySize:
		Concurrency:        config.ServerConfig.Concurrency,
	}
	if err := s.ListenAndServe(":" + config.ServerConfig.Port); err != nil {
		log.Fatal(err)
	}
}
