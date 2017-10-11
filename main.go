package main

import (
	"log"

	"github.com/shohi/yclite/router"
	"github.com/valyala/fasthttp"
)

func main() {
	port := ":8080"
	log.Printf("server started ===> http://localhost%s/list/1\n", port)
	log.Fatal(fasthttp.ListenAndServe(port, router.Router.Handler))
}
