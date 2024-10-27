package result

// Map applies a function to the value inside the Result, if it exists (Ok).
//
// Type signature:
//
//	Map :: Result a -> (a -> b) -> Result b
//
// It transforms the value inside the Result, or returns the original error if it is Err.
func Map[T, U any](r Result[T], fn func(T) U) Result[U] {
	if r.isErr {
		return Err[U](r.err)
	}
	return Ok(fn(r.value))
}

// BiMap applies one of two functions depending on whether the Result is Ok or Err.
//
// Type signature:
//
//	BiMap :: Result a -> (a -> b) -> (error -> error) -> Result b
//
// If the Result is Ok, it applies the fn function to the value. If it's Err, it applies the errFn function to the error.
func BiMap[T, U any](r Result[T], fn func(T) U, errFn func(error) error) Result[U] {
	if r.isErr {
		return Err[U](errFn(r.err))
	}
	return Ok(fn(r.value))
}
