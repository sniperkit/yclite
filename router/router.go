package router

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"

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

// return selected hacker news based on `points` and `filter` pattern

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

	query := ctx.QueryArgs()

	// points filter, e.g. http://localhost:8080/list/1?points=100
	pointsFilterFn := func() func(model.HackerNews) bool {
		points := string(query.Peek("points"))
		vs, err := util.ParseIntRange(points)
		if err != nil {
			return func(news model.HackerNews) bool {
				return true
			}
		}
		low, upper := math.MinInt32, math.MaxInt32
		low = vs[0]
		if len(vs) > 1 {
			upper = vs[1]
		}
		return func(news model.HackerNews) bool {
			return news.Points >= low && news.Points <= upper
		}
	}()

	// keyword filter, e.g. http://localhost:8080/list/1?filter=go
	keywordFilterFn := func() func(model.HackerNews) bool {
		ptn := strings.TrimSpace(string(query.Peek("filter")))
		if ptn == "" {
			return func(_ model.HackerNews) bool {
				return true
			}
		}

		// ignore case
		re, err := regexp.Compile("(?i).*" + ptn + ".*")
		if err != nil {
			return func(_ model.HackerNews) bool {
				return true
			}
		}
		return func(news model.HackerNews) bool {
			return re.Match([]byte(news.Title))
		}
	}()

	hackers := filter(filter(hn, pointsFilterFn), keywordFilterFn)
	ctx.Response.Header.SetContentType("text/html; charset=utf-8")
	ctxData := model.ContextData{Hackers: hackers, Filter: ctx.QueryArgs().String()}
	listTemplate.Execute(ctx, ctxData)
}

func filter(hn model.HackerNewsSlice, fn func(model.HackerNews) bool) model.HackerNewsSlice {
	var hackers model.HackerNewsSlice
	for _, v := range hn {
		if fn(v) {
			hackers = append(hackers, v)
		}
	}
	return hackers
}
