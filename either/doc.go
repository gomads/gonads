/*
Package either provides an implementation of the Either monad, commonly used in functional programming to represent a value that can be one of two types: a success (Right) or a failure (Left).

The Either type is useful for handling computations or operations that can result in two distinct outcomes—typically a successful result or an error. This avoids the need for explicit error handling or the use of panics, by keeping the result and error in the same structure.

Either consists of:

	Left: Represents an error or alternative outcome.
	Right: Represents a successful value or result.

Usage Example:

	func divide(a, b int) Either[string, int] {
	    if b == 0 {
	        return Left[string, int]("division by zero")
	    }
	    return Right[string, int](a / b)
	}

In this example, Left contains the error string, while Right holds the result of the division if successful.
*/
package either
