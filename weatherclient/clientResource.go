// weatherclient.go
package weatherclient

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

type WeatherClientResource struct {
	Client *fasthttp.Client
}

// NewWeatherClientResource creates a new WeatherClientResource with custom configuration.
func NewWeatherClientResource() (*WeatherClientResource, error) {
	// Creating a new WeatherClientResource
	fmt.Println("Creating Weather Resource...")
	client := &fasthttp.Client{
		MaxConnsPerHost: 10,
	}
	return &WeatherClientResource{Client: client}, nil
}
