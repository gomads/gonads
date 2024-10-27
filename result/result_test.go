package result_test

import (
	"errors"
	"testing"

	"github.com/alsi-lawr/gonads/result"
)

func TestOk(t *testing.T) {
	res := result.Ok(42)
	if !res.IsOk() {
		t.Errorf("expected Ok result")
	}
	if res.IsErr() {
		t.Errorf("expected result to not be an error")
	}
}

func TestErr(t *testing.T) {
	errMsg := errors.New("test error")
	res := result.Err[int](errMsg)
	if !res.IsErr() {
		t.Errorf("expected Err result")
	}
	if res.IsOk() {
		t.Errorf("expected result to not be successful")
	}
}

func TestMatchOk(t *testing.T) {
	var okValue int
	res := result.Ok(42)
	res.Match(
		func(val int) {
			okValue = val
		},
		func(err error) {
			t.Errorf("unexpected error: %v", err)
		},
	)
	if okValue != 42 {
		t.Errorf("expected Ok value to be 42, got %d", okValue)
	}
}

func TestMatchErr(t *testing.T) {
	var errValue error
	res := result.Err[int](errors.New("test error"))
	res.Match(
		func(val int) {
			t.Errorf("unexpected success value: %d", val)
		},
		func(err error) {
			errValue = err
		},
	)
	if errValue == nil || errValue.Error() != "test error" {
		t.Errorf("expected Err value to be 'test error', got %v", errValue)
	}
}

func TestBindOk(t *testing.T) {
	res := result.Ok(10).Bind(func(val int) result.Result[int] {
		return result.Ok(val * 2)
	})
	if !res.IsOk() || res.IsErr() {
		t.Errorf("expected Ok result")
	}
}

func TestBindErr(t *testing.T) {
	res := result.Err[int](errors.New("test error")).Bind(func(val int) result.Result[int] {
		return result.Ok(val * 2)
	})
	if !res.IsErr() || res.IsOk() {
		t.Errorf("expected Err result")
	}
}

func TestBindNewOk(t *testing.T) {
	res := result.Bind(result.Ok(10), func(val int) result.Result[uint] {
		return result.Ok((uint)(val * 2))
	})
	if !res.IsOk() || res.IsErr() {
		t.Errorf("expected Ok result")
	}
}

func TestBindNewErr(t *testing.T) {
	res := result.Bind(result.Err[int](errors.New("test error")), func(val int) result.Result[uint] {
		return result.Ok((uint)(val * 2))
	})
	if !res.IsErr() || res.IsOk() {
		t.Errorf("expected Err result")
	}
}

func TestBiBindOk(t *testing.T) {
	res := result.BiBind(result.Ok(10), func(val int) result.Result[uint] {
		return result.Ok((uint)(val * 2))
	}, func(err error) result.Result[uint] {
		return result.Err[uint](err)
	})
	if !res.IsOk() || res.IsErr() {
		t.Errorf("expected Ok result")
	}
}

func TestBiBindErr(t *testing.T) {
	res := result.BiBind(result.Err[int](errors.New("test error")), func(val int) result.Result[uint] {
		return result.Ok((uint)(val * 2))
	}, func(err error) result.Result[uint] {
		return result.Err[uint](errors.New("New Error!"))
	})

	if !res.IsErr() || res.IsOk() {
		t.Errorf("expected Err result")
	}
	res.Match(func(val uint) {
		t.Errorf("unexpected success value: %d", val)
	}, func(err error) {
		if err == nil || err.Error() != "New Error!" {
			t.Errorf("expected Err value to be 'New Error!', got %v", err)
		}
	})
}

func TestBiBindInstanceOk(t *testing.T) {
	res := result.Ok(10).BiBind(func(val int) result.Result[int] {
		return result.Ok(val * 2)
	}, func(err error) result.Result[int] {
		return result.Err[int](err)
	})
	if !res.IsOk() || res.IsErr() {
		t.Errorf("expected Ok result")
	}
}

func TestBiBindInstanceErr(t *testing.T) {
	res := result.Err[int](errors.New("test error")).BiBind(func(val int) result.Result[int] {
		return result.Ok(val * 2)
	}, func(err error) result.Result[int] {
		return result.Err[int](errors.New("New Error!"))
	})

	if !res.IsErr() || res.IsOk() {
		t.Errorf("expected Err result")
	}
	res.Match(func(val int) {
		t.Errorf("unexpected success value: %d", val)
	}, func(err error) {
		if err == nil || err.Error() != "New Error!" {
			t.Errorf("expected Err value to be 'New Error!', got %v", err)
		}
	})
}

func TestMapOk(t *testing.T) {
	res := result.Map(result.Ok(21), func(val int) int {
		return val * 2
	})
	if !res.IsOk() || res.IsErr() {
		t.Errorf("expected Ok result")
	}
	res.Match(func(val int) {
		if val != 42 {
			t.Errorf("expected value to be 42, got %d", val)
		}
	}, func(err error) {
		t.Errorf("unexpected error")
	})
}

func TestMapErr(t *testing.T) {
	res := result.Map(result.Err[int](errors.New("test error")), func(val int) int {
		return val * 2
	})
	if !res.IsErr() || res.IsOk() {
		t.Errorf("expected Err result")
	}
}

func TestBiMapOk(t *testing.T) {
	res := result.BiMap(result.Ok(42), func(val int) uint {
		return (uint)(val * 2)
	}, func(err error) error {
		return errors.New("test error")
	})
	if !res.IsOk() || res.IsErr() {
		t.Errorf("expected Ok result")
	}
	res.Match(func(val uint) {
		if val != 84 {
			t.Errorf("expected value to be 84, got %d", val)
		}
	}, func(err error) {
		t.Errorf("unexpected error")
	})
}

func TestBiMapErr(t *testing.T) {
	res := result.BiMap(result.Err[int](errors.New("test error")), func(val int) uint {
		return (uint)(val * 2)
	}, func(err error) error {
		return errors.New("test error 2")
	})

	if res.IsOk() && !res.IsErr() {
		t.Errorf("expected error result")
	}
	res.Match(func(val uint) {
		t.Errorf("unexpected error")
	}, func(err error) {
		if err.Error() != "test error 2" {
			t.Errorf("expected error to be 'test error 2', got %v", err)
		}
	})
}
