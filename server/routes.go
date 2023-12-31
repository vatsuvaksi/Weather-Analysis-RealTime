package server

import (
	// "os"
	"os"
	indexcontroller "real-time-weather-app/controllers/apiV1/indexController"
	weathercontroller "real-time-weather-app/controllers/apiV1/weatherController"
	"real-time-weather-app/middleware"
	"real-time-weather-app/utils/loggers"

	// "real-time-weather-app/utils/loggers"

	"github.com/fasthttp/router"
	httpSwagger "github.com/swaggo/http-swagger"

	// "github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func MuxRouter() (*router.Router, error) {

	//Initialize Controllers
	indexController := indexcontroller.NewIndexControler()
	weathercontroller := weathercontroller.NewWeatherController()
	//Initializing Routers
	r := router.New()
	r.GET("/", indexController.WelcomeHome)

	// Versioning APIs
	// The following router works for v1
	r.GET("/v1/realTimeData", middleware.RateLimitMiddleware(weathercontroller.GetRealTime))
	r.GET("/v1/forecast", middleware.RateLimitMiddleware(weathercontroller.GetForecast))
	r.GET("/v1/future", middleware.RateLimitMiddleware(weathercontroller.GetFuture))
	r.GET("/v1/timeZone", middleware.RateLimitMiddleware(weathercontroller.GetTimeZone))

	// Serve Swagger UI at the /swagger URL
	r.GET("/swagger/{*filepath}", fasthttpadaptor.NewFastHTTPHandlerFunc(httpSwagger.WrapHandler))
	// Serve Swagger JSON file
	r.GET("/swagger/doc.json", func(ctx *fasthttp.RequestCtx) {
		// Read and serve the Swagger JSON file here
		swaggerJSON, err := os.ReadFile("C:/Users/vatsal/Desktop/go-workspace/src/Weather-Analysis-RealTime/docs/doc.json")
		if err != nil {
			loggers.Logger.Info(err)
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			return
		}
		ctx.SetContentType("application/json")
		ctx.Write(swaggerJSON)
	})

	return r, nil
}
