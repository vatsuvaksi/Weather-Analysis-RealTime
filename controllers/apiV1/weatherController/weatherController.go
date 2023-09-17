package weathercontroller

import (
	weatherservice "real-time-weather-app/services/apiV1/weatherService"
	networkutils "real-time-weather-app/utils/networkUtils"

	"github.com/valyala/fasthttp"
)

type weatherController struct {
	WeatherService *weatherservice.WeatherService
}

func NewWeatherController() *weatherController {
	return &weatherController{
		WeatherService: weatherservice.NewWeatherService(),
	}
}

func (wc *weatherController) GetRealTime(ctx *fasthttp.RequestCtx) {
	queryParams, _ := networkutils.FetchQueryParams(ctx)
	response, err := wc.WeatherService.GetRealTimeData(queryParams["q"])
	//TODO : Need to change this
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	networkutils.RenderSuccessResponse(ctx, response)
}
