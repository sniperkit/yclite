package router

import (
	"fmt"
	"testing"
	"github.com/valyala/fasthttp"
)

func TestFastHttpClient(t *testing.T) {
	c := &fasthttp.Client{}
	status, res, err := c.Get(nil, "https://news.ycombinator.com/")
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(string(res))
}

