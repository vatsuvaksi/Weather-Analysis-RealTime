package server

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func setRoutes() (*router.Router, error) {
	r := router.New()
	r.GET("/", index)
	return r, nil
}

func index(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Welcome!")
}
