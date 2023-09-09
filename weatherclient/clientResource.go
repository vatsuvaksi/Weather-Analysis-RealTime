// weatherclient.go
package weatherclient

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

type WeatherClientResource struct {
	Client  *fasthttp.Client
	BaseURL string
}

// NewWeatherClientResource creates a new WeatherClientResource with custom configuration.
func NewWeatherClientResource() (*WeatherClientResource, error) {
	// Creating a new WeatherClientResource
	fmt.Println("Creating Weather Resource...")
	client := &fasthttp.Client{
		MaxConnsPerHost:        10,
		DisablePathNormalizing: false,
	}
	return &WeatherClientResource{Client: client, BaseURL: "https://api.weatherapi.com/v1/"}, nil
}

func (wCR *WeatherClientResource) GetDataFromClient(url string, queryParams map[string]string) ([]byte, error) {
	// Create a new fasthttp request object
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("GET")
	defer fasthttp.ReleaseRequest(req)

	// Set the request URL
	finalQuery := wCR.BaseURL + url + "?" + buildQueryString(queryParams)
	//TODO : Find a way to encode this URL
	// Found out that if client's DisablePathNormalization is set to false it will take care of URI encoding
	req.SetRequestURI(finalQuery)

	// Create a new fasthttp response object
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	// Send the GET request
	if err := wCR.Client.Do(req, resp); err != nil {
		return nil, err
	}

	// Check the response status code
	if resp.StatusCode() != fasthttp.StatusOK {
		return nil, fmt.Errorf("Received non-OK status code: %d", resp.StatusCode())
	}

	// Get the response body as bytes
	responseBody := resp.Body()
	return responseBody, nil
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
