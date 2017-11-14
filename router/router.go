package router

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"math"
	"regexp"
	"sort"
	"strconv"

	"github.com/buaazp/fasthttprouter"
	"github.com/shohi/yclite/model"
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
	info := "Welcome!\n"
	fmt.Fprint(ctx, info)
}

// ToDo: add filter, e.g. http://some-path/page?points=1,3
func list(ctx *fasthttp.RequestCtx) {
	path := ctx.UserValue("path").(string)

	page := 1
	if m, err := regexp.MatchString(`/\d+$`, path); err == nil && m {
		tmp, e := strconv.Atoi(path[1:])
		if e == nil && tmp > page {
			page = tmp
		}
	}

	// sort hacker news by
	hn := util.ExtractHackerNews(page)
	sort.Sort(hn)

	// filter result by points
	var hackers model.HackerNewsSlice
	filter := ctx.QueryArgs()
	points := string(filter.Peek("points"))
	vs, err := util.ParseIntRange(points)
	if err == nil {
		low, upper := math.MinInt32, math.MaxInt32
		low = vs[0]
		if len(vs) > 1 {
			upper = vs[1]
		}
		for _, v := range hn {
			if v.Points < low || v.Points > upper {
				continue
			} else {
				hackers = append(hackers, v)
			}
		}
	} else {
		hackers = hn
	}

	ctx.Response.Header.SetContentType("text/html; charset=utf-8")
	ctxData := model.ContextData{Hackers: hackers, Filter: filter.String()}
	listTemplate.Execute(ctx, ctxData)
}
