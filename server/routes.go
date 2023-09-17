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
	// r.GET("/forecast", forecast)
	// r.GET("/future", index)
	// r.GET("/timeZone", index)

	return r, nil
}

// func index(ctx *fasthttp.RequestCtx) {
// 	fmt.Println("testing")
// 	ctx.WriteString("Welcome!")
// }
