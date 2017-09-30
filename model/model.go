package model

// HackerNews represent on record
type HackerNews struct {
	Title    []byte
	Points   int
	Comments int
	Author   []byte
	Link     []byte
	Domain   []byte
	Discuss  []byte
	Time     []byte
}
