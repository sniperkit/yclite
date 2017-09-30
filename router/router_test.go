package router

import (
	"fmt"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/valyala/fasthttp"
)

func TestFastHttpClient(t *testing.T) {
	c := &fasthttp.Client{}
	status, res, err := c.Get(nil, "https://news.ycombinator.com/")
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(string(res))
}

func TestGoQuery(t *testing.T) {
	doc, err := goquery.NewDocument("https://news.ycombinator.com/news?p=2")

	//
	cnt := 0
	ms := "#hnmain tr td table.itemlist tbody"
	// doc.Find("#hnmain tr td table.itemlist tbody tr:not(.spacer)").Each(func(i int, s *goquery.Selection) {
	doc.Find(ms + " tr.athing").Each(func(i int, s *goquery.Selection) {
		ns := s.NextUntil(ms + " tr.spacer")

		fmt.Printf("#######################  %d  #######################\n", cnt)
		id, _ := s.Attr("id")
		sls := s.Find("td.title a.storylink")
		link, _ := sls.Attr("href")
		fmt.Println(id)
		fmt.Println(sls.Text())
		fmt.Println(link)
		fmt.Println(s.Find("td.title span a span.sitestr").Text())
		fmt.Println(ns.Find("td.subtext span.age a").Text())
		fmt.Println(ns.Find("td.subtext span.score").Text())
		fmt.Println(ns.Find("td.subtext a").Last().Text())
		fmt.Println(ns.Find("td.subtext a.hnuser").Text())

		cnt++
	})
	fmt.Println(err)
}
