package weatherservice

import (
	"real-time-weather-app/models/network/weatherData"
	"real-time-weather-app/weatherclient/apis"
)

type WeatherService struct {
	// Can be used in future if needed to connect with database
}

func NewWeatherService() *WeatherService {
	return &WeatherService{}
}

func (ws *WeatherService) GetRealTimeData(q string) (*weatherData.RealTimeData, error) {
	// Calling Internal Services to Get RealTime Data
	response, err := apis.GetRealTimeData(q)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (ws *WeatherService) GetForecast(q, days string) (*weatherData.ForeCastData, error) {
	// Calling Internal Services to Get Forecast Data
	response, err := apis.GetForecast(q, days)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (ws *WeatherService) GetFuture(q, dt string) (*weatherData.FutureData, error) {
	// Calling Internal Services to Get Future Data
	response, err := apis.GetFuture(q, dt)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (ws *WeatherService) GetTimeZone(q, dt string) (*weatherData.TimeZoneData, error) {
	// Calling Internal Services to GetTimeZone Data
	response, err := apis.GetTimeZone(q)
	if err != nil {
		return nil, err
	}
	return response, nil
}
