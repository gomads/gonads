package funcs_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/alsi-lawr/gonads/funcs"
)

func TestFilter(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	got := funcs.Filter(input, func(x int) bool {
		return x%2 == 0
	})

	want := []int{2, 4, 6}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Filter() = %v, want %v", got, want)
	}
}

func TestFilterI(t *testing.T) {
	input := []string{"a", "b", "c", "d"}
	got := funcs.FilterI(input, func(i int, s string) bool {
		return i%2 != 0
	})
	want := []string{"b", "d"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FilterI() = %v, want %v", got, want)
	}
}

func TestFilterMap(t *testing.T) {
	input := map[string]int{"a": 1, "b": 2, "c": 3}
	got := funcs.FilterMap(input, func(k string, v int) bool {
		return v%2 == 0
	})
	want := map[string]int{"b": 2}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FilterMap() = %v, want %v", got, want)
	}
}

func TestFilterChan(t *testing.T) {
	in := make(chan int, 5)
	nums := []int{1, 2, 3, 4, 5}
	for _, n := range nums {
		in <- n
	}
	close(in)

	out := funcs.FilterChan(in, func(x int) bool {
		return x > 3
	})
	var got []int
	for x := range out {
		got = append(got, x)
	}
	want := []int{4, 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FilterChan() = %v, want %v", got, want)
	}
}

func TestFilterString(t *testing.T) {
	input := "hello, world!"
	got := funcs.FilterString(input, func(r rune) bool {
		return !strings.ContainsRune("aeiouAEIOU", r)
	})
	want := "hll, wrld!"
	if got != want {
		t.Errorf("FilterString() = %v, want %v", got, want)
	}
}

func TestFilterStringI(t *testing.T) {
	input := "abcdef"
	got := funcs.FilterStringI(input, func(i int, r rune) bool {
		return i%2 == 0
	})
	want := "ace"
	if got != want {
		t.Errorf("FilterStringI() = %v, want %v", got, want)
	}
}
