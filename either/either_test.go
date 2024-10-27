package either_test

import (
	"strconv"
	"testing"

	"github.com/alsi-lawr/gonads/either"
)

func TestEitherLeft(t *testing.T) {
	e := either.Left[string](42)
	if !e.IsLeft() || e.IsRight() {
		t.Errorf("expected Left result")
	}
}

func TestEitherRight(t *testing.T) {
	e := either.Right[string]("42")
	if !e.IsRight() || e.IsLeft() {
		t.Errorf("expected Right result")
	}
}

func TestMatchLeft(t *testing.T) {
	e := either.Left[string](42)
	var res *int
	e.Match(func(l int) {
		res = new(int)
		*res = l + 1
	}, func(r string) {
		t.Errorf("expected Right result")
	})
	if res == nil || *res != 43 {
		t.Errorf("expected Left result")
	}
}

func TestMatchRight(t *testing.T) {
	e := either.Right[int]("42")
	var res *string
	e.Match(func(l int) {
		t.Errorf("expected Left result")
	}, func(r string) {
		res = new(string)
		*res = r + "1"
	})
	if res == nil || *res != "421" {
		t.Errorf("expected Left result")
	}
}

func TestLBindLeft(t *testing.T) {
	e := either.Left[string](42)
	result := e.LBind(func(l int) either.Either[int, string] {
		return either.Left[string](l + 1)
	})
	if !result.IsLeft() || result.IsRight() {
		t.Errorf("expected Left result")
	}
	unwrapped := result.LeftOrNil()
	if *unwrapped != 43 {
		t.Errorf("expected Left result")
	}
}

func TestLBindRight(t *testing.T) {
	e := either.Right[int]("42")
	result := e.LBind(func(l int) either.Either[int, string] {
		return either.Left[string](l + 1)
	})
	if result.IsLeft() || !result.IsRight() {
		t.Errorf("expected Right result")
	}
	unwrapped := result.RightOrNil()
	if *unwrapped != "42" {
		t.Errorf("expected Right result")
	}
}

func TestStaticLBindLeft(t *testing.T) {
	e := either.Left[int]("42")
	result := either.LBind(e, func(l string) either.Either[string, int] {
		res, _ := strconv.Atoi(l)
		return either.Right[string](res + 1)
	})
	if result.IsLeft() || !result.IsRight() {
		t.Errorf("expected Left result")
	}
	unwrapped := result.RightOrNil()
	if *unwrapped != 43 {
		t.Errorf("expected Left result")
	}
}

func TestStaticLBindRight(t *testing.T) {
	e := either.Right[string](42)
	result := either.LBind(e, func(l string) either.Either[string, int] {
		res, _ := strconv.Atoi(l)
		return either.Right[string](res + 1)
	})
	if result.IsLeft() || !result.IsRight() {
		t.Errorf("expected Left result")
	}
	unwrapped := result.RightOrNil()
	if *unwrapped != 42 {
		t.Errorf("expected Left result")
	}
}

func TestStaticLMapLeft(t *testing.T) {
	e := either.Left[int]("42")
	result := either.LMap(e, func(l string) int {
		res, _ := strconv.Atoi(l)
		return res + 1
	})
	if !result.IsLeft() || result.IsRight() {
		t.Errorf("expected Left result")
	}
	unwrapped := result.LeftOrNil()
	if *unwrapped != 43 {
		t.Errorf("expected Left result")
	}
}

func TestStaticLMapRight(t *testing.T) {
	e := either.Right[string](42)
	result := either.LMap(e, func(l string) int {
		res, _ := strconv.Atoi(l)
		return res + 1
	})
	if result.IsLeft() || !result.IsRight() {
		t.Errorf("expected Left result")
	}
	unwrapped := result.LeftOrNil()
	if unwrapped != nil {
		t.Errorf("expected Left result")
	}
}

func TestRBindLeft(t *testing.T) {
	e := either.Left[string](42)
	result := e.RBind(func(l string) either.Either[int, string] {
		return either.Right[int](l + "1")
	})
	if !result.IsLeft() || result.IsRight() {
		t.Errorf("expected Left result")
	}
	unwrapped := result.LeftOrNil()
	if *unwrapped != 42 {
		t.Errorf("expected Left result")
	}
}

func TestRBindRight(t *testing.T) {
	e := either.Right[int]("42")
	result := e.RBind(func(r string) either.Either[int, string] {
		return either.Right[int](r + "1")
	})
	if result.IsLeft() || !result.IsRight() {
		t.Errorf("expected Right result")
	}
	unwrapped := result.RightOrNil()
	if *unwrapped != "421" {
		t.Errorf("expected Right result")
	}
}

func TestStaticRBindLeft(t *testing.T) {
	e := either.Left[string](42)
	result := either.RBind(e, func(l string) either.Either[int, int] {
		res, _ := strconv.Atoi(l)
		return either.Right[int](res + 1)
	})
	if !result.IsLeft() || result.IsRight() {
		t.Errorf("expected Left result")
	}
	unwrapped := result.LeftOrNil()
	if *unwrapped != 42 {
		t.Errorf("expected Left result")
	}
}

func TestStaticRBindRight(t *testing.T) {
	e := either.Right[int]("42")
	result := either.RBind(e, func(r string) either.Either[int, int] {
		res, _ := strconv.Atoi(r)
		return either.Right[int](res + 1)
	})
	if result.IsLeft() || !result.IsRight() {
		t.Errorf("expected Right result")
	}
	unwrapped := result.RightOrNil()
	if *unwrapped != 43 {
		t.Errorf("expected Right result")
	}
}

func TestStaticRMapRight(t *testing.T) {
	e := either.Right[int]("42")
	result := either.RMap(e, func(r string) int {
		res, _ := strconv.Atoi(r)
		return res + 1
	})
	if result.IsLeft() || !result.IsRight() {
		t.Errorf("expected right result")
	}
	unwrapped := result.RightOrNil()
	if *unwrapped != 43 {
		t.Errorf("expected right result")
	}
}

func TestStaticRMapLeft(t *testing.T) {
	e := either.Left[int]("42")
	result := either.RMap(e, func(r int) int {
		return r + 1
	})
	if !result.IsLeft() || result.IsRight() {
		t.Errorf("expected right result")
	}
	unwrapped := result.RightOrNil()
	if unwrapped != nil {
		t.Errorf("expected right result")
	}
}

func TestBiBindLeft(t *testing.T) {
	e := either.Left[string](42)
	result := either.BiBind(e, func(l int) either.Either[int, string] {
		return either.Left[string](l + 1)
	}, func(r string) either.Either[int, string] {
		return either.Right[int](r + "1")
	})
	if !result.IsLeft() || result.IsRight() {
		t.Errorf("expected Left result")
	}
	unwrapped := result.LeftOrNil()
	if *unwrapped != 43 {
		t.Errorf("expected Left result")
	}
}

func TestBiBindRight(t *testing.T) {
	e := either.Right[int]("42")
	result := either.BiBind(e, func(l int) either.Either[int, string] {
		return either.Left[string](l + 1)
	}, func(r string) either.Either[int, string] {
		return either.Right[int](r + "1")
	})
	if result.IsLeft() || !result.IsRight() {
		t.Errorf("expected Right result")
	}
	unwrapped := result.RightOrNil()
	if *unwrapped != "421" {
		t.Errorf("expected Right result")
	}
}

func TestBiMapLeft(t *testing.T) {
	e := either.Left[string](42)
	result := either.BiMap(e, func(l int) int {
		return l + 1
	}, func(r string) int {
		res, _ := strconv.Atoi(r)
		return res + 1
	})
	if result != 43 {
		t.Errorf("expected Left result")
	}
}

func TestBiMapRight(t *testing.T) {
	e := either.Right[int]("42")
	result := either.BiMap(e, func(l int) int {
		return l + 1
	}, func(r string) int {
		res, _ := strconv.Atoi(r)
		return res + 1
	})
	if result != 43 {
		t.Errorf("expected right result")
	}
}

func TestBiBindMapLeft(t *testing.T) {
	e := either.Left[string](42)
	result := either.BiBindMap(e, func(l int) int {
		return l + 1
	}, func(r string) string {
		return r + "1"
	})
	if !result.IsLeft() || result.IsRight() {
		t.Errorf("expected Left result")
	}
	unwrapped := result.LeftOrNil()
	if *unwrapped != 43 {
		t.Errorf("expected Left result")
	}
}

func TestBiBindMapRight(t *testing.T) {
	e := either.Right[int]("42")
	result := either.BiBindMap(e, func(l int) int {
		return l + 1
	}, func(r string) string {
		return r + "1"
	})
	if result.IsLeft() || !result.IsRight() {
		t.Errorf("expected Right result")
	}
	unwrapped := result.RightOrNil()
	if *unwrapped != "421" {
		t.Errorf("expected Right result")
	}
}

func TestLeftOrNilNil(t *testing.T) {
	e := either.Right[string](42)
	unwrapped := e.LeftOrNil()
	if unwrapped != nil {
		t.Errorf("expected nil result")
	}
}

func TestRightOrNilNil(t *testing.T) {
	e := either.Left[string](42)
	unwrapped := e.RightOrNil()
	if unwrapped != nil {
		t.Errorf("expected nil result")
	}
}
