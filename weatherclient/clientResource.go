// weatherclient.go
package weatherclient

import (
	"fmt"
	"sync"

	"github.com/valyala/fasthttp"
)

type WeatherClientResource struct {
	Client  *fasthttp.Client
	BaseURL string
}

var singletonWeatherClient *WeatherClientResource
var once sync.Once
var baseURL = "https://api.weatherapi.com/v1/"

// NewWeatherClientResource creates a new WeatherClientResource with custom configuration.
func GetWeatherClientResource() (*WeatherClientResource, error) {
	// Creating a new WeatherClientResource
	once.Do(func() {
		fmt.Println("Creating Weather Resource Singleton Object...")
		singletonWeatherClient = &WeatherClientResource{
			Client: &fasthttp.Client{
				MaxConnsPerHost:        10,
				DisablePathNormalizing: false,
			},
			BaseURL: baseURL,
		}
	})
	return singletonWeatherClient, nil
}

func (wCR *WeatherClientResource) GetDataFromClient(url string, queryParams map[string]string) ([]byte, error) {
	// Create a new fasthttp request object
	fmt.Println("GetDataFromClient Initiated for --> ", url)
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("GET")
	defer fasthttp.ReleaseRequest(req)

	// Set the request URL
	finalQuery := wCR.BaseURL + url + "?" + buildQueryString(queryParams)
	// fmt.Println("final query: ", finalQuery)
	//TODO : Find a way to encode this URL
	// Found out that if client's DisablePathNormalization is set to false it will take care of URI encoding
	req.SetRequestURI(finalQuery)

	// Create a new fasthttp response object
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	// Making a copy of acquired response as the response will be empty necause defer statements will
	// release the response to connection pool
	// var responseCopy *fasthttp.Response
	// resp.CopyTo(responseCopy)

	// Send the GET request
	if err := wCR.Client.Do(req, resp); err != nil {
		return nil, err
	}

	// Check the response status code
	if resp.StatusCode() != fasthttp.StatusOK {
		return nil, fmt.Errorf("Received non-OK status code: %d", resp.StatusCode())
	}

	// Get the response body as bytes
	// responseBody := resp.Body()
	fmt.Println("GetDataFromClient Completed and response returned for --> ", finalQuery)

	return resp.Body(), nil
}

func buildQueryString(params map[string]string) string {
	var queryString string

	for k, v := range params {
		if queryString == "" {
			queryString = k + "=" + v
		} else {
			queryString += "&" + k + "=" + v
		}
	}
	return queryString
}
