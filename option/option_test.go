package option_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alsi-lawr/gonads/option"
)

func TestMakeSome(t *testing.T) {
	opt := option.Some(42)
	if !opt.IsSome() {
		t.Errorf("expected Some result")
	}
}

func TestMakeNone(t *testing.T) {
	opt := option.None[int]()
	if opt.IsSome() {
		t.Errorf("expected None result")
	}
}

func TestWrapSome(t *testing.T) {
	item := 42
	opt := option.Wrap(&item)
	if !opt.IsSome() {
		t.Errorf("expected Some result")
	}
}

func TestWrapNone(t *testing.T) {
	opt := option.Wrap[int](nil)
	if opt.IsSome() {
		t.Errorf("expected None result")
	}
}

func TestIsSome(t *testing.T) {
	opt := option.Some(42)
	if !opt.IsSome() {
		t.Errorf("expected Some result")
	}
	opt = option.None[int]()
	if opt.IsSome() {
		t.Errorf("expected None result")
	}
}

func TestIsNone(t *testing.T) {
	opt := option.Some(42)
	if opt.IsNone() {
		t.Errorf("expected Some result")
	}
	opt = option.None[int]()
	if !opt.IsNone() {
		t.Errorf("expected None result")
	}
}

func TestGetOrNilSome(t *testing.T) {
	opt := option.Some(42)
	result := opt.GetOrNil()
	if *result != 42 {
		t.Errorf("expected Some result")
	}
}

func TestGetOrNilNone(t *testing.T) {
	opt := option.None[int]()
	result := opt.GetOrNil()
	if result != nil {
		t.Errorf("expected nil result")
	}
}

func TestGetOrElseSome(t *testing.T) {
	opt := option.Some(42)
	result := opt.GetOrElse(func() int { return 0 })
	if result != 42 {
		t.Errorf("expected Some result")
	}
}

func TestGetOrElseNone(t *testing.T) {
	opt := option.None[int]()
	result := opt.GetOrElse(func() int { return 0 })
	if result != 0 {
		t.Errorf("expected else result")
	}
}

func TestOrElseSome(t *testing.T) {
	opt := option.Some(42)
	wrap := 0
	result := opt.OrElse(func() option.Option[int] { return option.Wrap(&wrap) })
	if result.IsNone() {
		t.Errorf("expected Some result")
	}
	unwrapped := result.GetOrNil()
	if *unwrapped != 42 {
		t.Errorf("expected Some result")
	}
}

func TestOrElseNone(t *testing.T) {
	opt := option.None[int]()
	wrap := 0
	result := opt.OrElse(func() option.Option[int] { return option.Wrap(&wrap) })

	if result.IsNone() {
		t.Errorf("expected Some result")
	}
	unwrapped := result.GetOrNil()
	if *unwrapped != 0 {
		t.Errorf("expected else result")
	}
}

func TestFlattenSome(t *testing.T) {
	opt := option.Some(option.Some(42))
	result := option.Flatten(opt) // option.Some(42)
	if result.IsNone() {
		t.Errorf("expected Some result")
	}
	unwrapped := result.GetOrNil()
	if *unwrapped != 42 {
		t.Errorf("expected Some result")
	}
}

func TestFlattenSomeNone(t *testing.T) {
	opt := option.Some(option.None[int]())
	result := option.Flatten(opt) // option.None[int]()
	if !result.IsNone() {
		t.Errorf("expected None result")
	}
}

func TestFlattenNone(t *testing.T) {
	opt := option.None[option.Option[int]]()
	result := option.Flatten(opt) // option.None[int]()
	if !result.IsNone() {
		t.Errorf("expected None result")
	}
}
func TestToSliceSome(t *testing.T) {
	opt := option.Some(42)
	result := opt.ToSlice()
	if len(result) != 1 || result[0] != 42 {
		t.Errorf("expected Some result")
	}
}

func TestToSliceNone(t *testing.T) {
	opt := option.None[int]()
	result := opt.ToSlice()
	if len(result) != 0 {
		t.Errorf("expected empty result")
	}
}

func TestZipWithSome(t *testing.T) {
	opt1 := option.Some(42)
	opt2 := option.Some("bar")
	result := option.Zip(opt1, opt2)
	if result.IsNone() {
		t.Errorf("expected Some result")
	}
	unwrapped := result.GetOrNil()
	if unwrapped.First != 42 || unwrapped.Second != "bar" {
		t.Errorf("expected Some result")
	}
}

func TestZipWithNone(t *testing.T) {
	opt1 := option.Some(42)
	opt2 := option.None[string]()
	result := option.Zip(opt1, opt2)
	if !result.IsNone() {
		t.Errorf("expected None result")
	}
}

func TestStaticBindSome(t *testing.T) {
	opt := option.Some(42)

	result := option.Bind(opt, func(x int) option.Option[string] { return option.Some(fmt.Sprintf("%d", x)) })
	if result.IsNone() {
		t.Errorf("expected Some result")
	}
	unwrapped := result.GetOrNil()
	if *unwrapped != "42" {
		t.Errorf("expected Some result")
	}
}

func TestStaticBindNone(t *testing.T) {
	opt := option.None[int]()

	result := option.Bind(opt, func(x int) option.Option[string] { return option.Some(fmt.Sprintf("%d", x)) })
	if !result.IsNone() {
		t.Errorf("expected Some result")
	}
}

func TestBindSome(t *testing.T) {
	opt := option.Some(42)
	result := opt.Bind(func(x int) option.Option[int] { return option.Some(x + 1) })

	if result.IsNone() {
		t.Errorf("expected Some result")
	}
	unwrapped := result.GetOrNil()
	if *unwrapped != 43 {
		t.Errorf("expected Some result")
	}
}

func TestBindNone(t *testing.T) {
	opt := option.None[int]()
	result := opt.Bind(func(x int) option.Option[int] { return option.Some(x + 1) })
	if !result.IsNone() {
		t.Errorf("expected Some result")
	}
}

func TestMapSome(t *testing.T) {
	opt := option.Some(42)
	result := option.Map(opt, func(x int) string { return fmt.Sprintf("%d", x+1) })
	if result.IsNone() {
		t.Errorf("expected Some result")
	}
	unwrapped := result.GetOrNil()
	if *unwrapped != "43" {
		t.Errorf("expected Some result")
	}
}

func TestMapNone(t *testing.T) {
	opt := option.None[int]()
	result := option.Map(opt, func(x int) string { return fmt.Sprintf("%d", x+1) })
	if !result.IsNone() {
		t.Errorf("expected Some result")
	}
}

func TestFilterSomeTrue(t *testing.T) {
	opt := option.Some(42)
	result := option.Filter(opt, func(x int) bool {
		if x == 42 {
			return true
		} else {
			return false
		}
	})
	if result.IsNone() {
		t.Errorf("expected Some result")
	}
	unwrapped := result.GetOrNil()
	if *unwrapped != 42 {
		t.Errorf("expected Some result")
	}
}

func TestFilterSomeFalse(t *testing.T) {
	opt := option.Some(42)
	result := option.Filter(opt, func(x int) bool {
		if x == 42 {
			return false
		} else {
			return true
		}
	})
	if !result.IsNone() {
		t.Errorf("expected Some result")
	}
}

func TestFilterNone(t *testing.T) {
	opt := option.None[int]()
	result := option.Filter(opt, func(x int) bool {
		return true
	})
	if !result.IsNone() {
		t.Errorf("expected Some result")
	}
}

func TestStaticBiMapSome(t *testing.T) {
	opt := option.Some(42)
	result := option.BiMap(opt, func(x int) string { return fmt.Sprintf("%d", x+1) }, func() string { return "none" })

	if result != "43" {
		t.Errorf("expected Some result")
	}
}

func TestBiMapNone(t *testing.T) {
	opt := option.None[int]()
	result := option.BiMap(opt, func(x int) string { return fmt.Sprintf("%d", x+1) }, func() string { return "none" })

	if result != "none" {
		t.Errorf("expected Some result")
	}
}

func TestMatchSome(t *testing.T) {
	opt := option.Some(42)
	var res1, res2 *int
	opt.Match(func(x int) {
		res1 = new(int)
		*res1 = x + 1
	}, func() {
		res2 = new(int)
		*res2 = 0
	})

	if *res1 != 43 || res2 != nil {
		t.Errorf("expected some result")
	}
}

func TestMatchNone(t *testing.T) {
	opt := option.None[int]()
	var res1, res2 *int
	opt.Match(func(x int) {
		res1 = new(int)
		*res1 = x + 1
	}, func() {
		res2 = new(int)
		*res2 = 0
	})

	if res1 != nil || *res2 != 0 {
		t.Errorf("expected none result")
	}
}

func TestTrySome(t *testing.T) {
	result := option.Try(func() (string, error) { return "foo", nil })
	if result.IsNone() {
		t.Errorf("expected Some result")
	}
	unwrapped := result.GetOrNil()
	if *unwrapped != "foo" {
		t.Errorf("expected Some result")
	}
}

func TestTry(t *testing.T) {
	result := option.Try(func() (*string, error) { return nil, errors.New("error") })
	if !result.IsNone() {
		t.Errorf("expected some result")
	}
}

func TestEqualsBothNone(t *testing.T) {
	opt1 := option.None[int]()
	opt2 := option.None[int]()
	result := opt1.Equals(opt2)
	if !result {
		t.Errorf("expected true")
	}
}

func TestEqualsOneNone(t *testing.T) {
	opt1 := option.Some(41)
	opt2 := option.None[int]()
	result := opt1.Equals(opt2)
	if result {
		t.Errorf("expected false")
	}
}

func TestEqualsBothSome(t *testing.T) {
	opt1 := option.Some(41)
	opt2 := option.Some(41)
	result := opt1.Equals(opt2)
	if !result {
		t.Errorf("expected true")
	}
}

func TestNequalsBothSome(t *testing.T) {
	opt1 := option.Some(42)
	opt2 := option.Some(41)
	result := opt1.Equals(opt2)
	if result {
		t.Errorf("expected false")
	}
}
