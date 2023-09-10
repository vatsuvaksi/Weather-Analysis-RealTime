package apis

import (
	"fmt"
	"os"
	"real-time-weather-app/weatherclient"

	"github.com/valyala/fasthttp"
)

/*
Function Calls the weather API to get real time Data using the paramater q which can either be
lattitude,longitude value in the format "lat,lang" or it can be a single zip code "zipCode"
*/
func GetRealTimeData(q string) (*fasthttp.Response, error) {
	var apiKey = os.Getenv("API_KEY")
	fmt.Println("Initiated API call for GetRealTimeData")
	weatherClientResource, err := weatherclient.GetWeatherClientResource()
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		var url = "current.json"
		response, err := weatherClientResource.GetDataFromClient(url, map[string]string{
			"q":   q,
			"key": apiKey,
		})

		if err != nil {
			fmt.Println(err)
			return nil, err
		} else {
			fmt.Println("Get Real Time Data Generated response")
			return response, nil
		}
	}
}
