package main

import (
	"log"

	"github.com/shohi/yclite/router"
	"github.com/valyala/fasthttp"
)

func main() {
	log.Fatal(fasthttp.ListenAndServe(":8080", router.Router.Handler))
}
