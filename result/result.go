package result

// Result represents a computation that can either return a successful value (Ok)
// or an error (Err). It is used to handle operations that may fail.
//
// Type signature:
//
//	Result[T] :: a -> Result a
type Result[T any] struct {
	value T
	err   error
	isErr bool
}

// Lift lifts a function that returns a value and an error into a Result.
//
// Type signature:
//
//	Lift :: (a -> (b, error)) -> Result b
//
// It returns a Result that is either Ok or Err depending on the error.
func Lift[T any](fn func() (T, error)) Result[T] {
	val, err := fn()
	if err != nil {
		return Err[T](err)
	}
	return Ok(val)
}

// Ok creates a Result representing a successful computation with a value.
//
// Type signature:
//
//	Ok :: a -> Result a
//
// It returns a Result where isErr is false, indicating success.
func Ok[T any](val T) Result[T] {
	return Result[T]{value: val, isErr: false}
}

// Err creates a Result representing a failed computation with an error.
//
// Type signature:
//
//	Err :: error -> Result a
//
// It returns a Result where isErr is true, indicating failure.
func Err[T any](err error) Result[T] {
	return Result[T]{err: err, isErr: true}
}

// IsErr returns true if the Result is an error (Err), false otherwise.
//
// Type signature:
//
//	IsErr :: Result a -> Bool
func (r Result[T]) IsErr() bool {
	return r.isErr
}

// IsOk returns true if the Result is successful (Ok), false otherwise.
//
// Type signature:
//
//	IsOk :: Result a -> Bool
func (r Result[T]) IsOk() bool {
	return !r.isErr
}

// Match applies one of two functions depending on whether the Result is Ok or Err.
//
// Type signature:
//
//	Match :: Result a -> (a -> ()) -> (error -> ()) -> ()
//
// If the Result is Ok, it applies the ifOk function to the value. If it's Err, it applies the ifErr function to the error.
func (r Result[T]) Match(ifOk func(T), ifErr func(error)) {
	if r.isErr {
		ifErr(r.err)
		return
	}
	ifOk(r.value)
}
