package networkutils

import (
	"encoding/json"
	"real-time-weather-app/models/network"

	"github.com/valyala/fasthttp"
)

func RenderSuccessResponse(ctx *fasthttp.RequestCtx, responseBody interface{}) {
	// Set the Content-Type header to indicate JSON content
	ctx.Response.Header.Set(fasthttp.HeaderContentType, "application/json; charset=utf-8")
	// Write the JSON response to the client
	ctx.Response.SetStatusCode(fasthttp.StatusOK)
	jsonResponse, err := json.Marshal(&responseBody)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
	ctx.Response.SetBody(jsonResponse)
}

func RenderInternalFailiureResponseForLimitExceed(ctx *fasthttp.RequestCtx, err error) {
	ctx.Response.Header.Set(fasthttp.HeaderContentType, "application/json; charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusTooManyRequests)
	errorStruct := &network.FailiureResponse{
		ErrorMessage: err.Error(),
		Status:       "Rate Limit Exceeded",
	}
	jsonResponse, err := json.Marshal(errorStruct)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
	ctx.Response.SetBody(jsonResponse)
}
func RenderInternalFailiureResponse(ctx *fasthttp.RequestCtx, err error) {
	ctx.Response.Header.Set(fasthttp.HeaderContentType, "application/json; charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	errorStruct := &network.FailiureResponse{
		ErrorMessage: err.Error(),
		Status:       "Internal Server Error",
	}
	jsonResponse, err := json.Marshal(errorStruct)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
	ctx.Response.SetBody(jsonResponse)
}

func FetchQueryParams(ctx *fasthttp.RequestCtx) (map[string]string, error) {
	queryParams := make(map[string]string)
	// Get all the key-value pairs from the request param
	queryArgs := ctx.QueryArgs()
	queryArgs.VisitAll(func(key, value []byte) {
		queryParams[string(key)] = string(value)
	})
	return queryParams, nil
}
func fillQueryParams(k, v []byte) ([]byte, []byte) {
	return nil, nil
}
