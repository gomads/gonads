package iters_test

import (
	"reflect"
	"testing"

	"github.com/alsi-lawr/gonads/iters"
)

func TestFoldStatic(t *testing.T) {
	input := []int{1, 2, 3, 4}
	sum := iters.Fold(input, 0, func(acc, x int) int {
		return acc + x
	})
	if sum != 10 {
		t.Errorf("funcs.Fold() = %d, want %d", sum, 10)
	}
}

func TestFoldIStatic(t *testing.T) {
	input := []int{1, 2, 3, 4}
	// Calculation: 0*1 + 1*2 + 2*3 + 3*4 = 0 + 2 + 6 + 12 = 20.
	result := iters.FoldI(input, 0, func(i, acc, x int) int {
		return acc + i*x
	})
	if result != 20 {
		t.Errorf("funcs.FoldI() = %d, want %d", result, 20)
	}
}

func TestFold(t *testing.T) {
	input := iters.Mappable[int, int]{1, 2, 3, 4}
	sum := input.Fold(0, func(acc, x int) int {
		return acc + x
	})
	if sum != 10 {
		t.Errorf("funcs.Fold() = %d, want %d", sum, 10)
	}
}

func TestFoldI(t *testing.T) {
	input := iters.Mappable[int, int]{1, 2, 3, 4}
	// Calculation: 0*1 + 1*2 + 2*3 + 3*4 = 0 + 2 + 6 + 12 = 20.
	result := input.FoldI(0, func(i, acc, x int) int {
		return acc + i*x
	})
	if result != 20 {
		t.Errorf("funcs.FoldI() = %d, want %d", result, 20)
	}
}

func TestFoldUnsafe(t *testing.T) {
	input := iters.Collection[int]{1, 2, 3, 4}
	sum := input.FoldUnsafe(0, func(acc any, x int) any {
		return acc.(int) + x
	})
	if sum != 10 {
		t.Errorf("funcs.Fold() = %d, want %d", sum, 10)
	}
}

func TestFoldIUnsafe(t *testing.T) {
	input := iters.Collection[int]{1, 2, 3, 4}
	// Calculation: 0*1 + 1*2 + 2*3 + 3*4 = 0 + 2 + 6 + 12 = 20.
	result := input.FoldIUnsafe(0, func(i int, acc any, x int) any {
		return acc.(int) + i*x
	})
	if result != 20 {
		t.Errorf("funcs.FoldI() = %d, want %d", result, 20)
	}
}

func TestFoldMapStatic(t *testing.T) {
	input := map[string]int{"a": 1, "b": 2, "c": 3}
	sum := iters.FoldMap(input, 0, func(acc int, k string, v int) int {
		return acc + v
	})
	if sum != 6 {
		t.Errorf("funcs.FoldMap() = %d, want %d", sum, 6)
	}
}

func TestFoldChanStatic(t *testing.T) {
	ch := make(chan int, 4)
	nums := []int{1, 2, 3, 4}
	for _, n := range nums {
		ch <- n
	}
	close(ch)
	sum := iters.FoldChan(ch, 0, func(acc, x int) int {
		return acc + x
	})

	if sum != 10 {
		t.Errorf("funcs.FoldChan() = %d, want %d", sum, 10)
	}
}

func TestFoldStringStatic(t *testing.T) {
	s := "ABC" // 'A'=65, 'B'=66, 'C'=67
	sum := iters.FoldString(s, 0, func(acc int, r rune) int {
		return acc + int(r)
	})
	expected := 65 + 66 + 67
	if sum != expected {
		t.Errorf("funcs.FoldString() = %d, want %d", sum, expected)
	}
}

func TestFoldStringI(t *testing.T) {
	s := "ABC"
	// Calculation: (0+65) + (1+66) + (2+67) = 65 + 67 + 69 = 201.
	result := iters.FoldStringI(s, 0, func(i int, acc int, r rune) int {
		return acc + i + int(r)
	})
	if result != 201 {
		t.Errorf("funcs.FoldStringI() = %d, want %d", result, 201)
	}
}

func TestFoldEmptySliceStatic(t *testing.T) {
	input := []int{}
	sum := iters.Fold(input, 100, func(acc, x int) int {
		return acc + x
	})
	if sum != 100 {
		t.Errorf("funcs.Fold() on empty slice = %d, want %d", sum, 100)
	}
}

func TestFoldMapEmptyStatic(t *testing.T) {
	input := map[string]int{}
	sum := iters.FoldMap(input, 50, func(acc int, k string, v int) int {
		return acc + v
	})
	if sum != 50 {
		t.Errorf("funcs.FoldMap() on empty map = %d, want %d", sum, 50)
	}
}

func TestFoldChanEmptyStatic(t *testing.T) {
	ch := make(chan int)
	close(ch)
	sum := iters.FoldChan(ch, 200, func(acc, x int) int {
		return acc + x
	})
	if sum != 200 {
		t.Errorf("funcs.FoldChan() on empty channel = %d, want %d", sum, 200)
	}
}

func TestFoldBuildSliceStatic(t *testing.T) {
	input := []string{"a", "b", "c"}
	result := iters.Fold(input, []string{}, func(acc []string, s string) []string {
		return append(acc, s+"1")
	})

	want := []string{"a1", "b1", "c1"}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("funcs.Fold() building slice = %v, want %v", result, want)
	}
}
