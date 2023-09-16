package server

import (
	"fmt"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func MuxRouter() (*router.Router, error) {
	r := router.New()
	r.GET("/", index)
	return r, nil
}

func index(ctx *fasthttp.RequestCtx) {
	fmt.Println("testing")
	ctx.WriteString("Welcome!")
}
