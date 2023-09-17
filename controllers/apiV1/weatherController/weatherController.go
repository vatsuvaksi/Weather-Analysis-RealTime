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
		networkutils.RenderInternalFailiureResponse(ctx, err)
		return
	}
	networkutils.RenderSuccessResponse(ctx, response)
}

func (wc *weatherController) GetForecast(ctx *fasthttp.RequestCtx) {
	queryParams, _ := networkutils.FetchQueryParams(ctx)
	response, err := wc.WeatherService.GetForecast(queryParams["q"], queryParams["days"])
	//TODO : Need to change this
	if err != nil {
		networkutils.RenderInternalFailiureResponse(ctx, err)
		return
	}
	networkutils.RenderSuccessResponse(ctx, response)
}

func (wc *weatherController) GetFuture(ctx *fasthttp.RequestCtx) {
	queryParams, _ := networkutils.FetchQueryParams(ctx)
	response, err := wc.WeatherService.GetFuture(queryParams["q"], queryParams["dt"])
	//TODO : Need to change this
	if err != nil {
		networkutils.RenderInternalFailiureResponse(ctx, err)
		return
	}
	networkutils.RenderSuccessResponse(ctx, response)
}

func (wc *weatherController) GetTimeZone(ctx *fasthttp.RequestCtx) {
	queryParams, _ := networkutils.FetchQueryParams(ctx)
	response, err := wc.WeatherService.GetTimeZone(queryParams["q"])
	//TODO : Need to change this
	if err != nil {
		networkutils.RenderInternalFailiureResponse(ctx, err)
		return
	}
	networkutils.RenderSuccessResponse(ctx, response)
}
