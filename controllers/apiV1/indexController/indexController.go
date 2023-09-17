package indexcontroller

import (
	indexservice "real-time-weather-app/services/apiV1/indexService"

	"github.com/valyala/fasthttp"
)

type IndexController struct {
	IndexServise *indexservice.IndexService
}

func NewIndexControler() *IndexController {
	return &IndexController{
		IndexServise: indexservice.NewIndexService(),
	}
}

func (ic *IndexController) WelcomeHome(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set(fasthttp.HeaderContentType, "application/json; charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Response.SetBodyString(ic.IndexServise.RandomWelcomeMessage())
}
