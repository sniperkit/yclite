package model

var BaseUrl = "https://news.ycombinator.com/news?p="
var DiscussUrl = "https://news.ycombinator.com/item?id="

// HackerNews represent one record
type HackerNews struct {
	Id 		string
	Title    string
	Points   int
	Comments int
	Sequence	int
	Page 	int
	Author   string
	Link     string
	Domain   string
	Discuss  string
	Time     string
}

type HackerNewsSlice []HackerNews

func (p HackerNewsSlice) Len() int {
	return len(p)
}

func (p HackerNewsSlice) Less(i, j int) bool {
	return p[i].Comments > p[j].Comments
}

func (p HackerNewsSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}


