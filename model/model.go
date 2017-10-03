package model

// HackerNews represent on record
type HackerNews struct {
	Id 		string
	Title    string
	Points   int
	Comments int
	Author   string
	Link     string
	Domain   string
	Discuss  string
	Time     string
}

var BaseUrl = "https://news.ycombinator.com/news?p="
var DiscussUrl = "https://news.ycombinator.com/item?id="
