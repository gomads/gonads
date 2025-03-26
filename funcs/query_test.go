package funcs_test

import (
	"strings"
	"testing"

	"github.com/alsi-lawr/gonads/funcs"
)

func TestFindStatic(t *testing.T) {
	ints := []int{1, 2, 3, 4}
	if v, ok := funcs.Find(ints, func(x int) bool { return x == 3 }); !ok || v != 3 {
		t.Errorf("funcs.Find() = (%d, %v), want (3, true)", v, ok)
	}

	if v, ok := funcs.Find(ints, func(x int) bool { return x == 5 }); ok {
		t.Errorf("funcs.Find() = (%d, %v), want (zero value, false)", v, ok)
	}
}

func TestFindIndexStatic(t *testing.T) {
	ints := []int{1, 2, 3, 4}
	index := funcs.FindIndex(ints, func(x int) bool { return x == 3 })
	if index != 2 {
		t.Errorf("funcs.FindIndex() = %d, want %d", index, 2)
	}

	index = funcs.FindIndex(ints, func(x int) bool { return x == 5 })
	if index != -1 {
		t.Errorf("funcs.FindIndex() = %d, want %d", index, -1)
	}
}

func TestFindFirstStatic(t *testing.T) {
	strs := []string{"apple", "banana", "cherry", "bana"}
	if v, ok := funcs.FindFirst(strs, func(s string) bool { return strings.HasPrefix(s, "ban") }); !ok ||
		v != "banana" {
		t.Errorf("funcs.FindFirst() = (%s, %v), want (banana, true)", v, ok)
	}

	if v, ok := funcs.FindFirst(strs, func(s string) bool { return s == "durian" }); ok {
		t.Errorf("funcs.FindFirst() = (%s, %v), want (zero value, false)", v, ok)
	}
}

func TestFindLastStatic(t *testing.T) {
	strs := []string{"apple", "bana", "cherry", "banana"}
	if v, ok := funcs.FindLast(strs, func(s string) bool { return strings.HasPrefix(s, "ban") }); !ok ||
		v != "banana" {
		t.Errorf("funcs.FindLast() = (%s, %v), want (banana, true)", v, ok)
	}

	if v, ok := funcs.FindLast(strs, func(s string) bool { return s == "durian" }); ok {
		t.Errorf("funcs.FindLast() = (%s, %v), want (zero value, false)", v, ok)
	}
}

func TestSomeStatic(t *testing.T) {
	ints := []int{1, 2, 3, 4}
	if !funcs.Some(ints, func(x int) bool { return x > 3 }) {
		t.Errorf("funcs.Some() returned false, want true")
	}
	if funcs.Some(ints, func(x int) bool { return x > 10 }) {
		t.Errorf("funcs.Some() returned true, want false")
	}
}

func TestAllStatic(t *testing.T) {
	ints := []int{2, 4, 6, 8}
	if !funcs.All(ints, func(x int) bool { return x%2 == 0 }) {
		t.Errorf("funcs.All() returned false, want true")
	}
	mixed := []int{2, 4, 5, 8}

	if funcs.All(mixed, func(x int) bool { return x%2 == 0 }) {
		t.Errorf("funcs.All() returned true, want false")
	}
}

func TestNoneStatic(t *testing.T) {
	ints := []int{1, 3, 5, 7}
	if !funcs.None(ints, func(x int) bool { return x%2 == 0 }) {
		t.Errorf("funcs.None() returned false, want true")
	}
	mixed := []int{1, 3, 4, 7}
	if funcs.None(mixed, func(x int) bool { return x%2 == 0 }) {
		t.Errorf("funcs.None() returned true, want false")
	}
}

func TestCountStatic(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5, 6}
	count := funcs.Count(ints, func(x int) bool { return x%2 == 0 })
	if count != 3 {
		t.Errorf("funcs.Count() = %d, want %d", count, 3)
	}

	count = funcs.Count(ints, func(x int) bool { return x > 10 })
	if count != 0 {
		t.Errorf("funcs.Count() = %d, want %d", count, 0)
	}
}

func TestFindOnStringsStatic(t *testing.T) {
	strs := []string{"alpha", "beta", "gamma", "delta"}
	if v, ok := funcs.Find(strs, func(s string) bool { return len(s) == 5 }); !ok || v != "alpha" {
		t.Errorf("funcs.Find() on strings = (%s, %v), want (alpha, true)", v, ok)
	}

	index := funcs.FindIndex(strs, func(s string) bool { return s == "delta" })
	if index != 3 {
		t.Errorf("funcs.FindIndex() on strings = %d, want %d", index, 3)
	}

	if v, ok := funcs.FindLast(strs, func(s string) bool { return s[0] == 'b' }); !ok ||
		v != "beta" {
		t.Errorf("funcs.FindLast() on strings = (%s, %v), want (beta, true)", v, ok)
	}
}

func TestEmptySliceQueriesStatic(t *testing.T) {
	var ints []int
	if _, ok := funcs.Find(ints, func(x int) bool { return true }); ok {
		t.Errorf("funcs.Find() on empty slice returned true, want false")
	}

	if idx := funcs.FindIndex(ints, func(x int) bool { return true }); idx != -1 {
		t.Errorf("funcs.FindIndex() on empty slice = %d, want -1", idx)
	}

	if _, ok := funcs.FindLast(ints, func(x int) bool { return true }); ok {
		t.Errorf("funcs.FindLast() on empty slice returned true, want false")
	}

	if funcs.Some(ints, func(x int) bool { return true }) {
		t.Errorf("funcs.Some() on empty slice returned true, want false")
	}

	if !funcs.All(ints, func(x int) bool { return false }) {
		t.Errorf("funcs.All() on empty slice returned false, want true (vacuously)")
	}

	if !funcs.None(ints, func(x int) bool { return true }) {
		t.Errorf("funcs.None() on empty slice returned false, want true")
	}

	if cnt := funcs.Count(ints, func(x int) bool { return true }); cnt != 0 {
		t.Errorf("funcs.Count() on empty slice = %d, want 0", cnt)
	}
}

func TestFind(t *testing.T) {
	ints := funcs.LiftSlice([]int{1, 2, 3, 4})
	if v, ok := ints.Find(func(x int) bool { return x == 3 }); !ok || v != 3 {
		t.Errorf("funcs.Find() = (%d, %v), want (3, true)", v, ok)
	}

	if v, ok := ints.Find(func(x int) bool { return x == 5 }); ok {
		t.Errorf("funcs.Find() = (%d, %v), want (zero value, false)", v, ok)
	}
}

func TestFindIndex(t *testing.T) {
	ints := funcs.LiftSlice([]int{1, 2, 3, 4})
	index := ints.FindIndex(func(x int) bool { return x == 3 })
	if index != 2 {
		t.Errorf("funcs.FindIndex() = %d, want %d", index, 2)
	}

	index = ints.FindIndex(func(x int) bool { return x == 5 })
	if index != -1 {
		t.Errorf("funcs.FindIndex() = %d, want %d", index, -1)
	}
}

func TestFindFirst(t *testing.T) {
	strs := funcs.LiftSlice([]string{"apple", "banana", "cherry", "bana"})
	if v, ok := strs.FindFirst(func(s string) bool { return strings.HasPrefix(s, "ban") }); !ok ||
		v != "banana" {
		t.Errorf("funcs.FindFirst() = (%s, %v), want (banana, true)", v, ok)
	}

	if v, ok := strs.FindFirst(func(s string) bool { return s == "durian" }); ok {
		t.Errorf("funcs.FindFirst() = (%s, %v), want (zero value, false)", v, ok)
	}
}

func TestFindLast(t *testing.T) {
	strs := funcs.LiftSlice([]string{"apple", "bana", "cherry", "banana"})
	if v, ok := strs.FindLast(func(s string) bool { return strings.HasPrefix(s, "ban") }); !ok ||
		v != "banana" {
		t.Errorf("funcs.FindLast() = (%s, %v), want (banana, true)", v, ok)
	}

	if v, ok := strs.FindLast(func(s string) bool { return s == "durian" }); ok {
		t.Errorf("funcs.FindLast() = (%s, %v), want (zero value, false)", v, ok)
	}
}

func TestSome(t *testing.T) {
	ints := funcs.LiftSlice([]int{1, 2, 3, 4})
	if !ints.Some(func(x int) bool { return x > 3 }) {
		t.Errorf("funcs.Some() returned false, want true")
	}
	if ints.Some(func(x int) bool { return x > 10 }) {
		t.Errorf("funcs.Some() returned true, want false")
	}
}

func TestAll(t *testing.T) {
	ints := funcs.LiftSlice([]int{2, 4, 6, 8})
	if !ints.All(func(x int) bool { return x%2 == 0 }) {
		t.Errorf("funcs.All() returned false, want true")
	}
	mixed := funcs.LiftSlice([]int{2, 4, 5, 8})
	if mixed.All(func(x int) bool { return x%2 == 0 }) {
		t.Errorf("funcs.All() returned true, want false")
	}
}

func TestNone(t *testing.T) {
	ints := funcs.LiftSlice([]int{1, 3, 5, 7})
	if !ints.None(func(x int) bool { return x%2 == 0 }) {
		t.Errorf("funcs.None() returned false, want true")
	}
	mixed := funcs.LiftSlice([]int{1, 3, 4, 7})
	if mixed.None(func(x int) bool { return x%2 == 0 }) {
		t.Errorf("funcs.None() returned true, want false")
	}
}

func TestCount(t *testing.T) {
	ints := funcs.LiftSlice([]int{1, 2, 3, 4, 5, 6})
	count := ints.Count(func(x int) bool { return x%2 == 0 })
	if count != 3 {
		t.Errorf("funcs.Count() = %d, want %d", count, 3)
	}

	count = ints.Count(func(x int) bool { return x > 10 })
	if count != 0 {
		t.Errorf("funcs.Count() = %d, want %d", count, 0)
	}
}

func TestFindOnStrings(t *testing.T) {
	strs := funcs.LiftSlice([]string{"alpha", "beta", "gamma", "delta"})
	if v, ok := strs.Find(func(s string) bool { return len(s) == 5 }); !ok || v != "alpha" {
		t.Errorf("funcs.Find() on strings = (%s, %v), want (alpha, true)", v, ok)
	}

	index := strs.FindIndex(func(s string) bool { return s == "delta" })
	if index != 3 {
		t.Errorf("funcs.FindIndex() on strings = %d, want %d", index, 3)
	}

	if v, ok := strs.FindLast(func(s string) bool { return s[0] == 'b' }); !ok ||
		v != "beta" {
		t.Errorf("funcs.FindLast() on strings = (%s, %v), want (beta, true)", v, ok)
	}
}

func TestEmptySliceQueries(t *testing.T) {
	var ints funcs.Collection[int]
	if _, ok := ints.Find(func(x int) bool { return true }); ok {
		t.Errorf("funcs.Find() on empty slice returned true, want false")
	}

	if idx := ints.FindIndex(func(x int) bool { return true }); idx != -1 {
		t.Errorf("funcs.FindIndex() on empty slice = %d, want -1", idx)
	}

	if _, ok := ints.FindLast(func(x int) bool { return true }); ok {
		t.Errorf("funcs.FindLast() on empty slice returned true, want false")
	}

	if ints.Some(func(x int) bool { return true }) {
		t.Errorf("funcs.Some() on empty slice returned true, want false")
	}

	if !ints.All(func(x int) bool { return false }) {
		t.Errorf("funcs.All() on empty slice returned false, want true (vacuously)")
	}

	if !ints.None(func(x int) bool { return true }) {
		t.Errorf("funcs.None() on empty slice returned false, want true")
	}

	if cnt := ints.Count(func(x int) bool { return true }); cnt != 0 {
		t.Errorf("funcs.Count() on empty slice = %d, want 0", cnt)
	}
}
