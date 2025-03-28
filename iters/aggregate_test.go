package iters_test

import (
	"reflect"
	"testing"

	"github.com/alsi-lawr/gonads/iters"
)

func TestAggregateGStatic(t *testing.T) {
	input := iters.Grouping[bool, int]{false: {1, 2, 3, 4}, true: {5, 6, 7, 8}}
	got := iters.Aggregate(input, func(y iters.Collection[int]) int { return 1 })
	want := map[bool]int{false: 1, true: 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("AggregateG() = %v, want %v", got, want)
	}
}

func TestAggregateIGStatic(t *testing.T) {
	input := iters.Grouping[bool, int]{false: {1, 2, 3, 4}, true: {5, 6, 7, 8}}
	got := iters.AggregateI(input, func(k bool, y iters.Collection[int]) int {
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

func TestAggregateGUnsafe(t *testing.T) {
	input := iters.Grouping[bool, int]{false: {1, 2, 3, 4}, true: {5, 6, 7, 8}}
	got := input.AggregateUnsafe(func(y iters.Collection[int]) any { return 1 })
	want := map[bool]any{false: 1, true: 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("AggregateG() = %v, want %v", got, want)
	}
}

func TestAggregateIGUnsafe(t *testing.T) {
	input := iters.Grouping[bool, int]{false: {1, 2, 3, 4}, true: {5, 6, 7, 8}}
	got := input.AggregateIUnsafe(func(k bool, y iters.Collection[int]) any {
		if k {
			return 1
		}
		return 0
	})

	want := map[bool]any{false: 0, true: 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("AggregateIG() = %v, want %v", got, want)
	}
}

func TestAggregateG(t *testing.T) {
	input := iters.Aggregable[bool, int, int]{false: {1, 2, 3, 4}, true: {5, 6, 7, 8}}
	got := input.Aggregate(func(y iters.Collection[int]) int { return 1 })
	want := map[bool]int{false: 1, true: 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("AggregateG() = %v, want %v", got, want)
	}
}

func TestAggregateIG(t *testing.T) {
	input := iters.Aggregable[bool, int, int]{false: {1, 2, 3, 4}, true: {5, 6, 7, 8}}
	got := input.AggregateI(func(k bool, y iters.Collection[int]) int {
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
