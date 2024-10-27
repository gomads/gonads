package result

// Bind applies a function to the value inside the Result, if it exists (Ok).
//
// Type signature:
//
//	Bind :: Result a -> (a -> Result a) -> Result a
//
// If the Result is Err, it returns the original error without applying the function.
func (r Result[T]) Bind(fn func(T) Result[T]) Result[T] {
	if r.isErr {
		return r
	}
	return fn(r.value)
}

// Bind applies a function to the value inside the Result, if it exists (Ok),
// allowing for transformations that may also return a Result.
//
// Type signature:
//
//	Bind :: Result a -> (a -> Result b) -> Result b
//
// If the Result is Err, it propagates the error.
func Bind[T, U any](r Result[T], fn func(T) Result[U]) Result[U] {
	return BiBind(r, fn, func(e error) Result[U] { return Err[U](e) })
}

// BiBind applies one of two functions depending on whether the Result is Ok or Err.
//
// Type signature:
//
//	BiBind :: Result a -> (a -> Result a) -> (error -> Result a) -> Result a
//
// If the Result is Ok, it applies the fn function to the value. If it's Err, it applies the errFn function to the error.
func (r Result[T]) BiBind(fn func(T) Result[T], errFn func(error) Result[T]) Result[T] {
	if r.isErr {
		return errFn(r.err)
	}
	return fn(r.value)
}

// BiBind applies one of two functions depending on whether the Result is Ok or Err.
// Transforms a result of type Result[T] into a result of type Result[U]
//
// Type signature:
//
//	BiBind :: Result a -> (a -> Result b) -> (error -> Result b) -> Result b
//
// If the Result is Ok, it applies the fn function to the value. If it's Err, it applies the errFn function to the error.
func BiBind[T, U any](r Result[T], fn func(T) Result[U], errFn func(error) Result[U]) Result[U] {
	if r.isErr {
		return errFn(r.err)
	}
	return fn(r.value)
}
