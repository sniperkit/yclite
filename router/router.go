package router

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"

	"github.com/buaazp/fasthttprouter"
	"github.com/shohi/yclite/util"
	"github.com/valyala/fasthttp"
)

// Router is the main router
var Router = fasthttprouter.New()
var listTemplate *template.Template
var funcMap template.FuncMap

func init() {
	Router.GET("/", index)
	Router.GET("/list/*path", list)
	funcMap = template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"equal": func(a, b int) bool {
			return a == b
		},
	}
	listTemplatePath := "template/list.html"
	bytes, _ := ioutil.ReadFile(listTemplatePath)
	listTemplate, _ = template.New(listTemplatePath).Funcs(funcMap).Parse(string(bytes))
	// listTemplate, _ = listTemplate.ParseFiles(listTemplatePath)
}

func index(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome!\n")
}

func list(ctx *fasthttp.RequestCtx) {
	path := ctx.UserValue("path").(string)
	page := 1
	if m, err := regexp.MatchString(`/\d+$`, path); err == nil && m {
		tmp, e := strconv.Atoi(path[1:])
		if e == nil && tmp > page {
			page = tmp
		}
	}
	hn := util.ExtractHackerNews(page)
	sort.Sort(hn)
	ctx.Response.Header.SetContentType("text/html; charset=utf-8")
	listTemplate.Execute(ctx, hn)
}
