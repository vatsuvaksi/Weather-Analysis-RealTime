package networkutils

import (
	"encoding/json"

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
