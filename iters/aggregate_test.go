package iters_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/alsi-lawr/gonads/iters"
)

func TestAggregateGStatic(t *testing.T) {
	input := iters.Grouping[bool, int]{false: {1, 2, 3, 4}, true: {5, 6, 7, 8}}
	got := iters.Aggregate(input, func(y iters.Iter[int]) int { return 1 })
	want := map[bool]int{false: 1, true: 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("AggregateG() = %v, want %v", got, want)
	}
}

func TestAggregateIGStatic(t *testing.T) {
	input := iters.Grouping[bool, int]{false: {1, 2, 3, 4}, true: {5, 6, 7, 8}}
	got := iters.AggregateI(input, func(k bool, y iters.Iter[int]) int {
		if k {
			return 1
		}
		return 0
	})

	want := map[bool]int{false: 0, true: 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("AggregateIG() = %v, want %v", got, want)
	}
}

func TestAggregateG(t *testing.T) {
	input := iters.Grouping[bool, int]{false: {1, 2, 3, 4}, true: {5, 6, 7, 8}}
	got := input.Aggregate[int](func(y iters.Iter[int]) int { return len(y) })
	want := map[bool]int{false: 4, true: 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("AggregateG() = %v, want %v", got, want)
	}
}

func TestAggregateGChangeType(t *testing.T) {
	input := iters.Grouping[bool, int]{false: {1, 2, 3, 4}, true: {5, 6, 7, 8}}
	got := input.Aggregate[string](func(y iters.Iter[int]) string { return fmt.Sprint(len(y)) })
	want := map[bool]string{false: "4", true: "4"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("AggregateG() = %v, want %v", got, want)
	}
}

func TestAggregateIG(t *testing.T) {
	input := iters.Grouping[bool, int]{false: {1, 2, 3, 4}, true: {5, 6, 7, 8}}
	got := input.AggregateI[int](func(k bool, y iters.Iter[int]) int {
		if k {
			return 1
		}
		return 0
	})

	want := map[bool]int{false: 0, true: 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("AggregateIG() = %v, want %v", got, want)
	}
}
