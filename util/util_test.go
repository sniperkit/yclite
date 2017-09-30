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
	}

	for _, c := range cases {
		got := ExtractInt(c.input)

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
