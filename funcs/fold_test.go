package funcs_test

import (
	"reflect"
	"testing"

	"github.com/alsi-lawr/gonads/funcs"
)

func TestFold(t *testing.T) {
	input := []int{1, 2, 3, 4}
	sum := funcs.Fold(input, 0, func(acc, x int) int {
		return acc + x
	})
	if sum != 10 {
		t.Errorf("funcs.Fold() = %d, want %d", sum, 10)
	}
}

func TestFoldI(t *testing.T) {
	input := []int{1, 2, 3, 4}
	// Calculation: 0*1 + 1*2 + 2*3 + 3*4 = 0 + 2 + 6 + 12 = 20.
	result := funcs.FoldI(input, 0, func(i, acc, x int) int {
		return acc + i*x
	})
	if result != 20 {
		t.Errorf("funcs.FoldI() = %d, want %d", result, 20)
	}
}

func TestFoldMap(t *testing.T) {
	input := map[string]int{"a": 1, "b": 2, "c": 3}
	sum := funcs.FoldMap(input, 0, func(acc int, k string, v int) int {
		return acc + v
	})
	if sum != 6 {
		t.Errorf("funcs.FoldMap() = %d, want %d", sum, 6)
	}
}

func TestFoldChan(t *testing.T) {
	ch := make(chan int, 4)
	nums := []int{1, 2, 3, 4}
	for _, n := range nums {
		ch <- n
	}
	close(ch)
	sum := funcs.FoldChan(ch, 0, func(acc, x int) int {
		return acc + x
	})

	if sum != 10 {
		t.Errorf("funcs.FoldChan() = %d, want %d", sum, 10)
	}
}

func TestFoldString(t *testing.T) {
	s := "ABC" // 'A'=65, 'B'=66, 'C'=67
	sum := funcs.FoldString(s, 0, func(acc int, r rune) int {
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
	result := funcs.FoldStringI(s, 0, func(i int, acc int, r rune) int {
		return acc + i + int(r)
	})
	if result != 201 {
		t.Errorf("funcs.FoldStringI() = %d, want %d", result, 201)

	}
}

func TestFoldEmptySlice(t *testing.T) {
	input := []int{}
	sum := funcs.Fold(input, 100, func(acc, x int) int {
		return acc + x
	})
	if sum != 100 {
		t.Errorf("funcs.Fold() on empty slice = %d, want %d", sum, 100)
	}
}

func TestFoldMapEmpty(t *testing.T) {
	input := map[string]int{}
	sum := funcs.FoldMap(input, 50, func(acc int, k string, v int) int {
		return acc + v
	})
	if sum != 50 {
		t.Errorf("funcs.FoldMap() on empty map = %d, want %d", sum, 50)
	}
}

func TestFoldChanEmpty(t *testing.T) {
	ch := make(chan int)
	close(ch)
	sum := funcs.FoldChan(ch, 200, func(acc, x int) int {
		return acc + x
	})
	if sum != 200 {
		t.Errorf("funcs.FoldChan() on empty channel = %d, want %d", sum, 200)
	}
}

func TestFoldBuildSlice(t *testing.T) {
	input := []string{"a", "b", "c"}
	result := funcs.Fold(input, []string{}, func(acc []string, s string) []string {
		return append(acc, s+"1")
	})

	want := []string{"a1", "b1", "c1"}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("funcs.Fold() building slice = %v, want %v", result, want)
	}
}
