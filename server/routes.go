package server

import (
	indexcontroller "real-time-weather-app/controllers/apiV1/indexController"
	weathercontroller "real-time-weather-app/controllers/apiV1/weatherController"

	"github.com/fasthttp/router"
)

func MuxRouter() (*router.Router, error) {

	//Initialize Controllers
	indexController := indexcontroller.NewIndexControler()
	weathercontroller := weathercontroller.NewWeatherController()
	//Initializing Routers
	r := router.New()
	r.GET("/", indexController.WelcomeHome)
	r.GET("/realTimeData", weathercontroller.GetRealTime)
	r.GET("/forecast", weathercontroller.GetForecast)
	r.GET("/future", weathercontroller.GetFuture)
	r.GET("/timeZone", weathercontroller.GetTimeZone)

	return r, nil
}
