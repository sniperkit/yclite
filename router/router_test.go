package router

import (
	"fmt"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/valyala/fasthttp"
	. "github.com/shohi/yclite/util"
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
	ms := "#hnmain tr td table.itemlist tbody"
	// doc.Find("#hnmain tr td table.itemlist tbody tr:not(.spacer)").Each(func(i int, s *goquery.Selection) {
	doc.Find(ms + " tr.athing").Each(func(i int, s *goquery.Selection) {
		fmt.Printf("#######################  %d  #######################\n", i)
		ExtractHackerNews(s, ms)
	})
	fmt.Println(err)
}
