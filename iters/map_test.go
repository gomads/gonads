package iters_test

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/alsi-lawr/gonads/iters"
)

func TestMapStatic(t *testing.T) {
	input := iters.Iter[int]{1, 2, 3}
	want := iters.Iter[int]{2, 4, 6}
	got := iters.Map(input, func(x int) int { return x * 2 })
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Map() = %v, want %v", got, want)
	}
}

func TestMapIStatic(t *testing.T) {
	input := iters.Iter[string]{"a", "b", "c"}
	want := iters.Iter[string]{"0:a", "1:b", "2:c"}
	got := iters.MapI(
		input,
		func(i int, s string) string { return strings.Join([]string{fmt.Sprint(i), ":", s}, "") },
	)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MapI() = %v, want %v", got, want)
	}
}

func TestMapErr_SuccessStatic(t *testing.T) {
	input := iters.Iter[int]{1, 2, 3}
	want := iters.Iter[int]{2, 4, 6}
	got, err := iters.MapErr(input, func(x int) (int, error) { return x * 2, nil })
	if err != nil {
		t.Fatalf("MapErr() unexpected error: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MapErr() = %v, want %v", got, want)
	}
}

func TestMapErr_ErrorStatic(t *testing.T) {
	input := []int{1, 2, 3}
	testErr := errors.New("error at 2")
	_, err := iters.MapErr(input, func(x int) (int, error) {
		if x == 2 {
			return 0, testErr
		}
		return x * 2, nil
	})

	if err == nil {
		t.Fatalf("MapErr() expected error, got nil")
	}

	if err.Error() != testErr.Error() {
		t.Errorf("MapErr() error = %v, want %v", err, testErr)
	}
}

func TestMapMapStatic(t *testing.T) {
	input := map[string]int{"a": 1, "b": 2}
	got := iters.MapMap(input, func(k string, v int) (string, int) {
		return strings.ToUpper(k), v * 10
	})
	want := map[string]int{"A": 10, "B": 20}
	if !reflect.DeepEqual(got, want) {

		t.Errorf("MapMap() = %v, want %v", got, want)
	}
}

func TestMapMapKeysStatic(t *testing.T) {
	input := map[string]int{"a": 1, "b": 2}
	got := iters.MapMapKeys(input, func(k string) string { return strings.ToUpper(k) })
	want := map[string]int{"A": 1, "B": 2}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MapMapKeys() = %v, want %v", got, want)
	}
}

func TestMapMapValuesStatic(t *testing.T) {
	input := map[string]int{"a": 1, "b": 2}
	got := iters.MapMapValues(input, func(v int) int { return v * 100 })
	want := map[string]int{"a": 100, "b": 200}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MapMapValues() = %v, want %v", got, want)
	}
}

func TestMapChanStatic(t *testing.T) {
	in := make(chan int, 3)
	in <- 1
	in <- 2
	in <- 3
	close(in)

	out := iters.MapChan(in, func(x int) int { return x * 3 })
	var got []int
	for v := range out {
		got = append(got, v)
	}
	want := []int{3, 6, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MapChan() = %v, want %v", got, want)
	}
}

func TestMapStringStatic(t *testing.T) {
	input := "abc"
	got := iters.MapString(input, func(r rune) rune {
		if r >= 'a' && r <= 'z' {
			return r - ('a' - 'A')
		}
		return r

	})

	want := "ABC"
	if got != want {
		t.Errorf("MapString() = %v, want %v", got, want)
	}
}

func TestMapStringIStatic(t *testing.T) {
	input := "abc"
	got := iters.MapStringI(input, func(i int, r rune) rune {
		return r + rune(i)
	})
	// 'a' + 0 = 'a', 'b' + 1 = 'c', 'c' + 2 = 'e'
	want := "ace"
	if got != want {
		t.Errorf("MapStringI() = %v, want %v", got, want)
	}
}

func TestMap(t *testing.T) {
	input := iters.LiftMap[int, int]([]int{1, 2, 3})
	want := iters.Iter[int]{2, 4, 6}
	got := input.Map(func(x int) int { return x * 2 })
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Map() = %v, want %v", got, want)
	}
}

func TestMapI(t *testing.T) {
	input := iters.LiftMap[string, string]([]string{"a", "b", "c"})
	want := iters.Iter[string]{"0:a", "1:b", "2:c"}
	got := input.MapI(
		func(i int, s string) string { return strings.Join([]string{fmt.Sprint(i), ":", s}, "") },
	)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MapI() = %v, want %v", got, want)
	}
}

func TestMapErr_Success(t *testing.T) {
	input := iters.LiftMap[int, int]([]int{1, 2, 3})
	want := iters.Iter[int]{2, 4, 6}
	got, err := input.MapErr(func(x int) (int, error) { return x * 2, nil })
	if err != nil {
		t.Fatalf("MapErr() unexpected error: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MapErr() = %v, want %v", got, want)
	}
}

func TestMapErr_Error(t *testing.T) {
	input := iters.LiftMap[int, int]([]int{1, 2, 3})
	testErr := errors.New("error at 2")
	_, err := input.MapErr(func(x int) (int, error) {
		if x == 2 {
			return 0, testErr
		}
		return x * 2, nil
	})

	if err == nil {
		t.Fatalf("MapErr() expected error, got nil")
	}

	if err.Error() != testErr.Error() {
		t.Errorf("MapErr() error = %v, want %v", err, testErr)
	}
}

func TestMapUnsafe(t *testing.T) {
	input := iters.Iter[int]{1, 2, 3}
	want := iters.Iter[int]{2, 4, 6}
	got := input.MapUnsafe(func(x int) any { return x * 2 })
	if !AnyDeepEqual(got, want) {
		t.Errorf("Map() = %v, want %v", got, want)
	}
}

func TestMapIUnsafe(t *testing.T) {
	input := iters.Iter[string]{"a", "b", "c"}
	want := iters.Iter[string]{"0:a", "1:b", "2:c"}
	got := input.MapIUnsafe(
		func(i int, s string) any { return strings.Join([]string{fmt.Sprint(i), ":", s}, "") },
	)
	if !AnyDeepEqual(got, want) {
		t.Errorf("MapI() = %v, want %v", got, want)
	}
}

func TestMapErr_SuccessUnsafe(t *testing.T) {
	input := iters.Iter[int]{1, 2, 3}
	want := iters.Iter[int]{2, 4, 6}
	got, err := input.MapErrUnsafe(func(x int) (any, error) { return x * 2, nil })
	if err != nil {
		t.Fatalf("MapErr() unexpected error: %v", err)
	}
	if !AnyDeepEqual(got, want) {
		t.Errorf("MapErr() = %v, want %v", got, want)
	}
}

func TestMapErr_ErrorUnsafe(t *testing.T) {
	input := iters.Iter[int]{1, 2, 3}
	testErr := errors.New("error at 2")
	_, err := input.MapErrUnsafe(func(x int) (any, error) {
		if x == 2 {
			return 0, testErr
		}
		return x * 2, nil
	})

	if err == nil {
		t.Fatalf("MapErr() expected error, got nil")
	}

	if err.Error() != testErr.Error() {
		t.Errorf("MapErr() error = %v, want %v", err, testErr)
	}
}
