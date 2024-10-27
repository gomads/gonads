/*
Package result provides an implementation of the Result monad, used to represent the outcome of computations that can either succeed (Ok) or fail (Err).

The Result type is particularly useful for operations that may return a value upon success or an error upon failure. It provides a safer and more structured alternative to Go's traditional error handling by keeping both the value and error in the same structure, preventing the need for error return values or panics.

Result consists of:

	Ok: Represents a successful result containing a value.
	Err: Represents a failure or error.

Usage Example:

	func divide(a, b int) Result[int] {
	    if b == 0 {
	        return Err[int](errors.New("division by zero"))
	    }
	    return Ok(a / b)
	}

In this example, `Err` is used to return an error message, while `Ok` contains the successful result.
*/
package result
