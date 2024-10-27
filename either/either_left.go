package either

// Left creates an Either with a Left value.
//
// Type signature:
//
//	Left :: L -> Either L R
//
// It returns an Either where isLeft is true, indicating that it contains a Left value.
func Left[R, L any](l L) Either[L, R] {
	return Either[L, R]{left: l, isLeft: true}
}

// IsLeft checks if the Either contains a Left value.
//
// Type signature:
//
//	IsLeft :: Either L R -> Bool
func (e Either[L, R]) IsLeft() bool {
	return e.isLeft
}

// LBind applies a function to the Left value if it exists, otherwise it returns the current Right value.
//
// Type signature:
//
//	LBind :: Either L R -> (L -> Either L R) -> Either L R
//
// If the Either is Left, it applies the l function to the left value. If it's Right, it returns the current Right value.
func (e Either[L, R]) LBind(l func(L) Either[L, R]) Either[L, R] {
	if e.isLeft {
		return l(e.left)
	}
	return Right[L](e.right)
}

// LBind applies a function to the Left value if it exists, otherwise it returns the current Right value.
//
// Type signature:
//
//	LBind :: Either L R -> (L -> Either U R) -> Either U R
//
// If the Either is Left, it applies the l function to the left value. If it's Right, it returns the current Right value.
func LBind[L, R, U any](e Either[L, R], l func(L) Either[U, R]) Either[U, R] {
	if e.isLeft {
		return l(e.left)
	}
	return Right[U](e.right)
}

// LMap applies a function to the Left value if it exists.
//
// Type signature:
//
//	LMap :: Either L R -> (L -> U) -> Either U R
//
// If the Either is Left, it applies the l function to the left value and wraps the result in a new Either.
// If the Either is Right, it returns the current Right value.
func LMap[L, R, U any](e Either[L, R], l func(L) U) Either[U, R] {
	if e.isLeft {
		return Left[R](l(e.left))
	}
	return Right[U](e.right)
}

// LeftOrNil returns the Left value if it exists, or nil if it doesn't.
//
// Type signature:
//
//	LeftOrNil :: Either L R -> L
func (l *Either[L, R]) LeftOrNil() *L {
	if l.isLeft {
		return &l.left
	}
	return nil
}
