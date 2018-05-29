package main

import (
	"flag"
	"log"

	"github.com/shohi/yclite/config"
	"github.com/shohi/yclite/router"
	"github.com/valyala/fasthttp"
)

func main() {
	flag.Parse()
	log.Printf("server started: %[1]s, list: http://localhost%[1]s/list/1", config.DefaultGlobalConfig.Port)
	log.Fatal(fasthttp.ListenAndServe(config.DefaultGlobalConfig.Port, router.Router.Handler))
}
