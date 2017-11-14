package util

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/shohi/yclite/model"
)

var defaultSize = 64

func ExtractInts(s string) []int {
	re, err := regexp.Compile("\\d+")
	if err != nil {
		return nil
	}

	var ret []int
	for _, v := range re.FindAllStringSubmatch(s, -1) {
		for _, m := range v {
			tmp, err := strconv.Atoi(m)
			if err == nil {
				ret = append(ret, tmp)
			}
		}
	}

	return ret
}

func ExtractHackerNews(p int) model.HackerNewsSlice {
	doc, err := goquery.NewDocument(model.BaseUrl + strconv.Itoa(p))

	if err != nil {
		return nil
	}

	var hn model.HackerNewsSlice

	//
	ms := "#hnmain tr td table.itemlist tbody"
	doc.Find(ms + " tr.athing").Each(func(i int, s *goquery.Selection) {
		ns := s.NextUntil(ms + " tr.spacer")
		hack := model.HackerNews{}

		id, _ := s.Attr("id")
		sls := s.Find("td.title a.storylink")
		link, _ := sls.Attr("href")

		hack.Id = id
		hack.Page = p
		hack.Sequence = i
		hack.Link = link
		hack.Discuss = model.DiscussUrl + id
		hack.Title = sls.Text()

		hack.Domain = s.Find("td.title span a span.sitestr").Text()
		hack.Time = ns.Find("td.subtext span.age a").Text()
		hack.Author = ns.Find("td.subtext a.hnuser").Text()

		if scores := ExtractInts(ns.Find("td.subtext span.score").Text()); len(scores) > 0 {
			hack.Points = scores[0]
		}

		if comments := ExtractInts(ns.Find("td.subtext a").Last().Text()); len(comments) > 0 {
			hack.Comments = comments[0]
		}
		hn = append(hn, hack)
	})

	return hn
}

// ParseIntRange - parse int range in format "a,b" where a and b are integer
func ParseIntRange(v string) (bounds []int, err error) {
	if v == "" {
		err = errors.New("empty string")
		return
	}

	vSlice := strings.Split(v, ",")
	if len(vSlice) < 1 {
		err = errors.New("No interval information")
		return
	}

	startRange, er := strconv.ParseUint(strings.TrimSpace(vSlice[0]), 10, 32)
	if er != nil {
		err = er
		return
	}
	bounds = append(bounds, int(startRange))

	if len(vSlice) >= 2 {
		endRange, er := strconv.ParseUint(strings.TrimSpace(vSlice[1]), 10, 32)
		if er != nil {
			err = er
			return
		}
		bounds = append(bounds, int(endRange))
	}

	return
}
