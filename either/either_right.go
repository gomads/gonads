package either

// Right creates an Either with a Right value.
//
// Type signature:
//
//	Right :: R -> Either L R
//
// It returns an Either where isLeft is false, indicating that it contains a Right value.
func Right[L, R any](r R) Either[L, R] {
	return Either[L, R]{right: r, isLeft: false}
}

// IsRight checks if the Either contains a Right value.
//
// Type signature:
//
//	IsRight :: Either L R -> Bool
func (e Either[L, R]) IsRight() bool {
	return !e.isLeft
}

// RBind applies a function to the Right value if it exists, otherwise it returns the current Left value.
//
// Type signature:
//
//	RBind :: Either L R -> (R -> Either L R) -> Either L R
//
// If the Either is Right, it applies the r function to the right value. If it's Left, it returns the current Left value.
func (e Either[L, R]) RBind(r func(R) Either[L, R]) Either[L, R] {
	if e.isLeft {
		return Left[R](e.left)
	}
	return r(e.right)
}

// RBind applies a function to the Right value if it exists, otherwise it returns the current Left value.
//
// Type signature:
//
//	RBind :: Either L R -> (R -> Either L U) -> Either L U
//
// If the Either is Right, it applies the r function to the right value. If it's Left, it returns the current Left value.
func RBind[L, R, U any](e Either[L, R], r func(R) Either[L, U]) Either[L, U] {
	if e.isLeft {
		return Left[U](e.left)
	}
	return r(e.right)
}

// RMap applies a function to the Right value if it exists.
//
// Type signature:
//
//	RMap :: Either L R -> (R -> U) -> Either L U
//
// If the Either is Right, it applies the r function to the right value and wraps the result in a new Either.
// If the Either is Left, it returns the current Left value.
func RMap[L, R, U any](e Either[L, R], r func(R) U) Either[L, U] {
	if e.isLeft {
		return Left[U](e.left)
	}
	return Right[L](r(e.right))
}

// RightOrNil returns the Right value if it exists, or nil if it doesn't.
//
// Type signature:
//
//	RightOrNil :: Either L R -> R
func (r Either[L, R]) RightOrNil() *R {
	if !r.isLeft {
		return &r.right
	}
	return nil
}
