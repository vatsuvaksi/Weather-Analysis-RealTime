package indexcontroller

import (
	indexservice "real-time-weather-app/services/apiV1/indexService"
	networkutils "real-time-weather-app/utils/networkUtils"

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
	networkutils.RenderSuccessResponse(ctx, ic.IndexServise.ResponseWelcomeMessage())
}
