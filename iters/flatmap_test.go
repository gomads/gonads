package iters_test

import (
	"reflect"
	"testing"

	"github.com/alsi-lawr/gonads/iters"
)

func TestFlatMapStatic(t *testing.T) {
	input := iters.Iter[int]{1, 2, 3}
	got := iters.FlatMap(input, func(x int) iters.Iter[int] {
		return []int{x, x * 10}
	})
	want := iters.Iter[int]{1, 10, 2, 20, 3, 30}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("FlatMap() = %v, want %v", got, want)
	}
}

func TestFlatMapIStatic(t *testing.T) {
	input := iters.Iter[string]{"a", "b", "c"}
	got := iters.FlatMapI(input, func(i int, s string) iters.Iter[string] {
		return []string{s, string(rune('A' + i))}
	})
	want := iters.Iter[string]{"a", "A", "b", "B", "c", "C"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FlatMapI() = %v, want %v", got, want)
	}
}

func TestFlatMap(t *testing.T) {
	input := iters.Mappable[int, int]{1, 2, 3}
	got := input.FlatMap(func(x int) iters.Iter[int] {
		return []int{x, x * 10}
	})
	want := iters.Iter[int]{1, 10, 2, 20, 3, 30}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FlatMap() = %v, want %v", got, want)
	}
}

func TestFlatMapI(t *testing.T) {
	input := iters.Mappable[string, string]{"a", "b", "c"}
	got := input.FlatMapI(func(i int, s string) iters.Iter[string] {
		return []string{s, string(rune('A' + i))}
	})
	want := iters.Iter[string]{"a", "A", "b", "B", "c", "C"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FlatMapI() = %v, want %v", got, want)
	}
}
func TestFlatMapUnsafe(t *testing.T) {
	input := iters.Iter[int]{1, 2, 3}
	got := input.FlatMapUnsafe(func(x int) iters.Iter[any] {
		return []any{x, x * 10}
	})
	want := iters.Iter[any]{1, 10, 2, 20, 3, 30}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("FlatMap() = %v, want %v", got, want)
	}
}

func TestFlatMapIUnsafe(t *testing.T) {
	input := iters.Iter[string]{"a", "b", "c"}
	got := input.FlatMapIUnsafe(func(i int, s string) iters.Iter[any] {
		return []any{s, string(rune('A' + i))}
	})
	want := iters.Iter[any]{"a", "A", "b", "B", "c", "C"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FlatMapI() = %v, want %v", got, want)
	}
}
