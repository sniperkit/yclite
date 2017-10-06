package router

import (
	"fmt"
	"html/template"
	"sort"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"github.com/shohi/yclite/util"
)

// Router is the main router
var Router = fasthttprouter.New()
var listTemplate *template.Template

func init() {
	Router.GET("/", index)
	Router.GET("/list/*path", list)
	listTemplate, _ = template.ParseFiles("template/list.html")
}

func index(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome!\n")
}

func list(ctx *fasthttp.RequestCtx) {
	// path := ctx.UserValue("path")
	// fmt.Fprintf(ctx, "path ==> %v, %T\n", path, path)
	hn := util.ExtractHackerNews(1)
	sort.Sort(hn)
	ctx.Response.Header.SetContentType("text/html; charset=utf-8")
	listTemplate.Execute(ctx, hn)
}
