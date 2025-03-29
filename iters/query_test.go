package iters_test

import (
	"strings"
	"testing"

	"github.com/alsi-lawr/gonads/iters"
)

func TestFindStatic(t *testing.T) {
	ints := []int{1, 2, 3, 4}
	iters.Find(ints, func(x int) bool { return x == 3 }).Match(
		func(x int) {},
		func() { t.Errorf("funcs.Find() = not found, want 3") },
	)

	iters.Find(ints, func(x int) bool { return x == 5 }).Match(
		func(x int) { t.Errorf("funcs.Find() = %d, want none", x) },
		func() {},
	)
}

func TestFindIndexStatic(t *testing.T) {
	ints := []int{1, 2, 3, 4}
	iters.FindIndex(ints, func(x int) bool { return x == 3 }).Match(
		func(x int) {
			if x != 2 {
				t.Errorf("funcs.Find() = %d, want 2", x)
			}
		},
		func() { t.Errorf("funcs.Find() = not found, want 2") },
	)

	iters.FindIndex(ints, func(x int) bool { return x == 5 }).Match(
		func(x int) { t.Errorf("funcs.Find() = %d, want none", x) },
		func() {},
	)
}

func TestFindFirstStatic(t *testing.T) {
	strs := []string{"apple", "banana", "cherry", "bana"}
	iters.FindFirst(strs, func(s string) bool { return strings.HasPrefix(s, "ban") }).Match(
		func(s string) {
			if s != "banana" {
				t.Errorf("funcs.FindFirst() = %s, want banana", s)
			}
		},
		func() { t.Errorf("funcs.FindFirst() = not found, want banana") },
	)
	iters.FindFirst(strs, func(s string) bool { return s == "durian" }).Match(
		func(s string) { t.Errorf("funcs.FindFirst() = %s, want none", s) },
		func() {},
	)
}

func TestFindLastStatic(t *testing.T) {
	strs := []string{"apple", "bana", "cherry", "banana"}
	iters.FindLast(strs, func(s string) bool { return strings.HasPrefix(s, "ban") }).Match(
		func(s string) {
			if s != "banana" {
				t.Errorf("funcs.FindLast() = %s, want banana", s)
			}
		},
		func() { t.Errorf("funcs.FindLast() = not found, want banana") },
	)
	iters.FindLast(strs, func(s string) bool { return s == "durian" }).Match(
		func(s string) { t.Errorf("funcs.FindLast() = %s, want none", s) },
		func() {},
	)
}

func TestAnyStatic(t *testing.T) {
	ints := []int{1, 2, 3, 4}
	if !iters.Any(ints, func(x int) bool { return x > 3 }) {
		t.Errorf("funcs.Any() returned false, want true")
	}
	if iters.Any(ints, func(x int) bool { return x > 10 }) {
		t.Errorf("funcs.Any() returned true, want false")
	}
}

func TestAllStatic(t *testing.T) {
	ints := []int{2, 4, 6, 8}
	if !iters.All(ints, func(x int) bool { return x%2 == 0 }) {
		t.Errorf("funcs.All() returned false, want true")
	}
	mixed := []int{2, 4, 5, 8}

	if iters.All(mixed, func(x int) bool { return x%2 == 0 }) {
		t.Errorf("funcs.All() returned true, want false")
	}
}

func TestCountStatic(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5, 6}
	count := iters.Count(ints, func(x int) bool { return x%2 == 0 })
	if count != 3 {
		t.Errorf("funcs.Count() = %d, want %d", count, 3)
	}

	count = iters.Count(ints, func(x int) bool { return x > 10 })
	if count != 0 {
		t.Errorf("funcs.Count() = %d, want %d", count, 0)
	}
}

func TestFindOnStringsStatic(t *testing.T) {
	strs := []string{"alpha", "beta", "gamma", "delta"}
	iters.Find(strs, func(s string) bool { return s == "alpha" }).Match(
		func(s string) {
			if s != "alpha" {
				t.Errorf("funcs.Find() on strings = %s, want alpha", s)
			}
		},
		func() { t.Errorf("funcs.Find() on strings = not found, want alpha") },
	)

	iters.FindIndex(strs, func(s string) bool { return s == "delta" }).Match(
		func(i int) {
			if i != 3 {
				t.Errorf("funcs.FindIndex() on strings = %d, want %d", i, 3)
			}
		},
		func() { t.Errorf("funcs.FindIndex() on strings = not found, want %d", 3) },
	)

	iters.FindLast(strs, func(s string) bool { return s == "beta" }).Match(
		func(s string) {
			if s != "beta" {
				t.Errorf("funcs.FindLast() on strings = %s, want beta", s)
			}
		},
		func() { t.Errorf("funcs.FindLast() on strings = not found, want beta") },
	)
}

func TestEmptySliceQueriesStatic(t *testing.T) {
	var ints []int

	iters.Find(ints, func(x int) bool { return true }).Match(
		func(x int) { t.Errorf("funcs.Find() on empty slice = %d, want none", x) },
		func() {},
	)

	iters.FindIndex(ints, func(x int) bool { return true }).Match(
		func(x int) { t.Errorf("funcs.FindIndex() on empty slice = %d, want -1", x) },
		func() {},
	)

	iters.FindLast(ints, func(x int) bool { return true }).Match(
		func(x int) { t.Errorf("funcs.FindLast() on empty slice = %d, want none", x) },
		func() {},
	)

	if iters.Any(ints, func(x int) bool { return true }) {
		t.Errorf("funcs.Any() on empty slice returned true, want false")
	}

	if cnt := iters.Count(ints, func(x int) bool { return true }); cnt != 0 {
		t.Errorf("funcs.Count() on empty slice = %d, want 0", cnt)
	}
}

func TestFind(t *testing.T) {
	ints := iters.Iter[int]{1, 2, 3, 4}
	ints.Find(func(x int) bool { return x == 3 }).Match(
		func(x int) {
			if x != 3 {
				t.Errorf("funcs.Find() = %d, want 3", x)
			}
		},
		func() { t.Errorf("funcs.Find() = not found, want 3") },
	)
	ints.Find(func(x int) bool { return x == 5 }).Match(
		func(x int) { t.Errorf("funcs.Find() = %d, want none", x) },
		func() {},
	)
}

func TestFindIndex(t *testing.T) {
	ints := iters.Iter[int]{1, 2, 3, 4}

	ints.FindIndex(func(x int) bool { return x == 3 }).Match(
		func(i int) {
			if i != 2 {
				t.Errorf("funcs.FindIndex() = %d, want %d", i, 2)
			}
		},
		func() { t.Errorf("funcs.FindIndex() = not found, want %d", 2) },
	)

	ints.FindIndex(func(x int) bool { return x == 5 }).Match(
		func(i int) { t.Errorf("funcs.FindIndex() = %d, want None", i) },
		func() {},
	)
}

func TestFindFirst(t *testing.T) {
	strs := iters.Iter[string]{"apple", "banana", "cherry", "bana"}
	strs.FindFirst(func(s string) bool { return strings.HasPrefix(s, "ban") }).Match(
		func(s string) {
			if s != "banana" {
				t.Errorf("funcs.FindFirst() = %s, want banana", s)
			}
		},
		func() { t.Errorf("funcs.FindFirst() = not found, want banana") },
	)
	strs.FindFirst(func(s string) bool { return s == "durian" }).Match(
		func(s string) { t.Errorf("funcs.FindFirst() = %s, want none", s) },
		func() {},
	)
}

func TestFindLast(t *testing.T) {
	strs := iters.Iter[string]{"apple", "bana", "cherry", "banana"}
	strs.FindLast(func(s string) bool { return strings.HasPrefix(s, "ban") }).Match(
		func(s string) {
			if s != "banana" {
				t.Errorf("funcs.FindLast() = %s, want banana", s)
			}
		},
		func() { t.Errorf("funcs.FindLast() = not found, want banana") },
	)
	strs.FindLast(func(s string) bool { return s == "durian" }).Match(
		func(s string) { t.Errorf("funcs.FindLast() = %s, want none", s) },
		func() {},
	)
}

func TestAny(t *testing.T) {
	ints := iters.Iter[int]{1, 2, 3, 4}
	if !ints.Any(func(x int) bool { return x > 3 }) {
		t.Errorf("funcs.Any() returned false, want true")
	}
	if ints.Any(func(x int) bool { return x > 10 }) {
		t.Errorf("funcs.Any() returned true, want false")
	}
}

func TestAll(t *testing.T) {
	ints := iters.Iter[int]{2, 4, 6, 8}
	if !ints.All(func(x int) bool { return x%2 == 0 }) {
		t.Errorf("funcs.All() returned false, want true")
	}
	mixed := iters.Iter[int]{2, 4, 5, 8}
	if mixed.All(func(x int) bool { return x%2 == 0 }) {
		t.Errorf("funcs.All() returned true, want false")
	}
}

func TestCount(t *testing.T) {
	ints := iters.Iter[int]{1, 2, 3, 4, 5, 6}
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
	strs := iters.Iter[string]{"alpha", "beta", "gamma", "delta"}
	strs.Find(func(s string) bool { return s == "alpha" }).Match(
		func(s string) {
			if s != "alpha" {
				t.Errorf("funcs.Find() on strings = %s, want alpha", s)
			}
		},
		func() { t.Errorf("funcs.Find() on strings = not found, want alpha") },
	)

	strs.FindIndex(func(s string) bool { return s == "alpha" }).Match(
		func(i int) {
			if i != 0 {
				t.Errorf("funcs.FindIndex() on strings = %d, want 0", i)
			}
		},
		func() { t.Errorf("funcs.FindIndex() on strings = not found, want 0") },
	)

	strs.FindFirst(func(s string) bool { return strings.Contains(s, "t") }).Match(
		func(s string) {
			if s != "beta" {
				t.Errorf("funcs.FindFirst() on strings = %s, want beta", s)
			}
		},
		func() { t.Errorf("funcs.FindFirst() on strings = not found, want beta") },
	)
	strs.FindLast(func(s string) bool { return strings.Contains(s, "t") }).Match(
		func(s string) {
			if s != "delta" {
				t.Errorf("funcs.FindFirst() on strings = %s, want beta", s)
			}
		},
		func() { t.Errorf("funcs.FindFirst() on strings = not found, want beta") },
	)
}

func TestEmptySliceQueries(t *testing.T) {
	var ints iters.Iter[int]

	ints.Find(func(x int) bool { return true }).Match(
		func(x int) { t.Errorf("funcs.Find() on empty slice = %d, want none", x) },
		func() {},
	)

	ints.FindIndex(func(x int) bool { return true }).Match(
		func(x int) { t.Errorf("funcs.FindIndex() on empty slice = %d, want -1", x) },
		func() {},
	)

	ints.FindFirst(func(x int) bool { return true }).Match(
		func(x int) { t.Errorf("funcs.FindFirst() on empty slice = %d, want none", x) },
		func() {},
	)

	ints.FindLast(func(x int) bool { return true }).Match(
		func(x int) { t.Errorf("funcs.FindLast() on empty slice = %d, want none", x) },
		func() {},
	)

	if ints.Any(func(x int) bool { return true }) {
		t.Errorf("funcs.Any() on empty slice returned true, want false")
	}

	if !ints.All(func(x int) bool { return false }) {
		t.Errorf("funcs.All() on empty slice returned false, want true (vacuously)")
	}

	if cnt := ints.Count(func(x int) bool { return true }); cnt != 0 {
		t.Errorf("funcs.Count() on empty slice = %d, want 0", cnt)
	}
}
