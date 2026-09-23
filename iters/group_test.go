package iters_test

import (
	"reflect"
	"testing"

	"github.com/alsi-lawr/gonads/iters"
)

func TestGroupByStatic(t *testing.T) {
	input := iters.Iter[int]{1, 2, 3, 4, 5, 6}
	got := iters.GroupBy(input, func(x int) bool {
		return x%2 == 0
	})
	want := iters.Grouping[bool, int]{
		false: {1, 3, 5},
		true:  {2, 4, 6},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GroupBy() = %v, want %v", got, want)
	}
}

func TestGroupByIStatic(t *testing.T) {
	input := iters.Iter[string]{"apple", "banana", "apricot", "blueberry", "avocado"}
	got := iters.GroupByI(input, func(i int, s string) string {
		if i%2 == 0 {
			return "even"
		}
		return "odd"
	})
	want := iters.Grouping[string, string]{
		"even": {"apple", "apricot", "avocado"},
		"odd":  {"banana", "blueberry"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GroupByI() = %v, want %v", got, want)
	}
}

func TestGroupBy(t *testing.T) {
	input := iters.Iter[int]{1, 2, 3, 4, 5, 6}
	got := input.GroupBy[string](func(x int) string {
		if x%2 == 0 {
			return "even"
		}
		return "odd"
	})
	want := iters.Grouping[string, int]{
		"odd":  {1, 3, 5},
		"even": {2, 4, 6},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GroupBy() = %v, want %v", got, want)
	}
}

func TestGroupByBool(t *testing.T) {
	input := iters.Iter[int]{1, 2, 3, 4, 5, 6}
	got := input.GroupBy[bool](func(x int) bool {
		return x%2 == 0
	})
	want := iters.Grouping[bool, int]{
		false: {1, 3, 5},
		true:  {2, 4, 6},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GroupBy() = %v, want %v", got, want)
	}
}

func TestGroupByI(t *testing.T) {
	input := iters.Iter[string]{"apple", "banana", "apricot", "blueberry", "avocado"}
	got := input.GroupByI[string](func(i int, s string) string {
		if i%2 == 0 {
			return "even"
		}
		return "odd"
	})
	want := iters.Grouping[string, string]{
		"even": {"apple", "apricot", "avocado"},
		"odd":  {"banana", "blueberry"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GroupByI() = %v, want %v", got, want)
	}
}
