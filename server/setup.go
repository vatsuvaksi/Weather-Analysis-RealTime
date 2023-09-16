package server

import (
	"log"

	"github.com/valyala/fasthttp"
)

func Start(port string) {
	router, err := setRoutes()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(fasthttp.ListenAndServe(":"+port, router.Handler))
}
