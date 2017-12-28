package router

import (
	"log"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestFastHttpClient(t *testing.T) {
	c := &fasthttp.Client{}
	status, res, err := c.Get(nil, "https://news.ycombinator.com/")
	log.Println(status)
	log.Println(err)
	log.Println(string(res))
}
