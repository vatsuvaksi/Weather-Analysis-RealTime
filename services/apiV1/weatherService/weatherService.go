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

// @Summary Get real-time weather data
// @Description Retrieve real-time weather data based on the provided query.
// @Tags Weather
// @Accept json
// @Produce json
// @Param q query string true "Location query"
// @Success 200 {object} RealTimeData "Successful response"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Failure 404 {object} ErrorResponse "Not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /weather/real-time [get]
func (ws *WeatherService) GetRealTimeData(q string) (*weatherData.RealTimeData, error) {
	// Calling Internal Services to Get RealTime Data
	response, err := apis.GetRealTimeData(q)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// @Summary Get weather forecast data
// @Description Retrieve weather forecast data based on the provided query and number of days.
// @Tags Weather
// @Accept json
// @Produce json
// @Param q query string true "Location query"
// @Param days query string true "Number of days for the forecast"
// @Success 200 {object} ForeCastData "Successful response"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Failure 404 {object} ErrorResponse "Not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /weather/forecast [get]
func (ws *WeatherService) GetForecast(q, days string) (*weatherData.ForeCastData, error) {
	// Calling Internal Services to Get Forecast Data
	response, err := apis.GetForecast(q, days)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// @Summary Get future weather data
// @Description Retrieve future weather data based on the provided query and date.
// @Tags Weather
// @Accept json
// @Produce json
// @Param q query string true "Location query"
// @Param dt query string true "Date for the forecast (YYYY-MM-DD format)"
// @Success 200 {object} FutureData "Successful response"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Failure 404 {object} ErrorResponse "Not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /weather/future [get]
func (ws *WeatherService) GetFuture(q, dt string) (*weatherData.FutureData, error) {
	// Calling Internal Services to Get Future Data
	response, err := apis.GetFuture(q, dt)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// @Summary Get time zone data
// @Description Retrieve time zone data based on the provided location query.
// @Tags Weather
// @Accept json
// @Produce json
// @Param q query string true "Location query"
// @Success 200 {object} TimeZoneData "Successful response"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Failure 404 {object} ErrorResponse "Not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /weather/time-zone [get]
func (ws *WeatherService) GetTimeZone(q string) (*weatherData.TimeZoneData, error) {
	// Calling Internal Services to GetTimeZone Data
	response, err := apis.GetTimeZone(q)
	if err != nil {
		return nil, err
	}
	return response, nil
}
