package server

import (
	"log"

	"github.com/valyala/fasthttp"
)

func Start() {
	router, err := setRoutes()
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}
