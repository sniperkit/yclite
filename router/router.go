package router

import (
	"fmt"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

// Router is main router
var Router = fasthttprouter.New()

func init() {
	Router.GET("/", index)
	Router.GET("/list/*path", list)
}

func index(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome!\n")
}

func list(ctx *fasthttp.RequestCtx) {
	path := ctx.UserValue("path")
	fmt.Fprintf(ctx, "path ==> %v, %T\n", path, path)
}
