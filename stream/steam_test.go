package stream

import (
	"fmt"
	"strings"
	"testing"
)

func TestMap(t *testing.T) {
	vals := []int{1, 2, 3, 4}
	stream := Slice(vals)
	ss := Map[int, string](stream, func(i int) string {
		return fmt.Sprintf("%v", i)
	})

	asSlice := Collect(ss)

	if strings.Join(asSlice, ", ") != "1, 2, 3, 4" {
		t.Fatalf("Unexpected slice: %v", asSlice)
	}
}

func TestFilter(t *testing.T) {
	vals := []int{1, 2, 3, 4}
	stream := Slice(vals)
	evens := Filter[int](stream, func(i int) bool {
		return i%2 == 0
	})

	ss := Map[int, string](evens, func(i int) string {
		return fmt.Sprintf("%v", i)
	})

	asSlice := Collect(ss)

	if strings.Join(asSlice, ", ") != "2, 4" {
		t.Fatalf("Unexpected slice: %v", asSlice)
	}
}
