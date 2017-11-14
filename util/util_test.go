package util

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestExtractInt(t *testing.T) {
	cases := []struct {
		input string
		want  []int
	}{
		{"aa", nil},
		{"123, 456", []int{123, 456}},
	}

	for _, c := range cases {
		got := ExtractInts(c.input)

		if len(got) == 0 && len(c.want) == 0 {
			continue
		}

		sort.Ints(got)
		sort.Ints(c.want)

		if len(got) != len(c.want) ||
			!reflect.DeepEqual(got, c.want) {
			t.Errorf("ExtractInt(%q) == %v, want %v", c.input, got, c.want)
		}
	}

}

func TestLenOnNil(t *testing.T) {
	var arr []int

	if arr == nil {
		fmt.Println("slice init with nil")
		fmt.Printf("len of nil slice returns %d\n", len(arr))
	}
}

func TestExtractHackerNews(t *testing.T) {
	p := 1
	fmt.Println(ExtractHackerNews(p))
}

func TestParseIntRange(t *testing.T) {
	v := "a,b"
	_, err := ParseIntRange(v)
	if err == nil {
		t.Errorf("ParseIntRange(%q) should return non-nil error", err)
	}

	v = "1, 20"
	fmt.Println(ParseIntRange(v))
}
