package server

import (
	indexcontroller "real-time-weather-app/controllers/apiV1/indexController"
	weathercontroller "real-time-weather-app/controllers/apiV1/weatherController"

	"github.com/fasthttp/router"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/valyala/fasthttp/fasthttpadaptor"
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

	// Serve Swagger UI at the /swagger URL
	r.GET("/swagger/{*filepath}", fasthttpadaptor.NewFastHTTPHandlerFunc(httpSwagger.WrapHandler))

	return r, nil
}
