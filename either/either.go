package either

// Either represents a value that can either be a Left or a Right.
// Left is typically used for error cases, and Right is used for success cases.
//
// Type signature:
//
//	Either[L, R] :: L -> R -> Either L R
type Either[L, R any] struct {
	left   L
	right  R
	isLeft bool
}

// Match applies one of two functions depending on whether the Either contains a Left or Right value.
//
// Type signature:
//
//	Match :: Either L R -> (L -> ()) -> (R -> ()) -> ()
//
// If the Either is Left, it applies the onLeft function to the left value.
// If the Either is Right, it applies the onRight function to the right value.
func (e Either[L, R]) Match(onLeft func(L), onRight func(R)) {
	if e.isLeft {
		onLeft(e.left)
		return
	}
	onRight(e.right)
}

// BiBind applies one of two functions depending on whether the Either contains a Left or Right value.
//
// Type signature:
//
//	BiBind :: Either L R -> (L -> Either U V) -> (R -> Either U V) -> Either U V
//
// If the Either is Left, it applies the l function to the left value.
// If the Either is Right, it applies the r function to the right value.
func BiBind[L, R, U, V any](e Either[L, R], l func(L) Either[U, V], r func(R) Either[U, V]) Either[U, V] {
	if e.isLeft {
		return l(e.left)
	}
	return r(e.right)
}

// BiMap applies one of two functions depending on whether the Either contains a Left or Right value.
//
// Type signature:
//
//	BiMap :: Either L R -> (L -> U) -> (R -> U) -> U
//
// If the Either is Left, it applies the l function to the left value.
// If the Either is Right, it applies the r function to the right value.
func BiMap[L, R, U any](e Either[L, R], l func(L) U, r func(R) U) U {
	if e.isLeft {
		return l(e.left)
	}
	return r(e.right)
}

// BiBindMap applies one of two functions depending on whether the Either contains a Left or Right value, and maps the results into a new Either.
//
// Type signature:
//
//	BiBindMap :: Either L R -> (L -> U) -> (R -> V) -> Either U V
//
// If the Either is Left, it applies the l function to the left value and wraps the result in a new Either.
// If the Either is Right, it applies the r function to the right value and wraps the result in a new Either.
func BiBindMap[L, R, U, V any](e Either[L, R], l func(L) U, r func(R) V) Either[U, V] {
	if e.isLeft {
		return Left[V](l(e.left))
	}
	return Right[U](r(e.right))
}
