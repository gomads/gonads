package iters

// FlatMap maps each element of a slice to a slice using the function f and then flattens the result
// into a single slice.
//
// Type signature:
//
//	FlatMap :: Iter T -> (T -> Iter R) -> Iter R
func FlatMap[T any, R any](s Iter[T], f func(T) Iter[R]) Iter[R] {
	var result []R
	for _, v := range s {
		result = append(result, f(v)...)
	}
	return result
}

// FlatMapI maps each element of a slice along with its index to a slice using the function f,
// and then flattens the result into a single slice.
//
// Type signature:
//
//	FlatMapI :: Iter T -> ((Int, T) -> Iter R) -> Iter R
func FlatMapI[T any, R any](s Iter[T], f func(int, T) Iter[R]) Iter[R] {
	var result []R
	for i, v := range s {
		result = append(result, f(i, v)...)
	}
	return result
}

// FlatMap maps each element of a slice to a slice using the function f and then flattens the result
// into a single slice.
//
// Type signature:
//
//	FlatMap :: Iter T -> (T -> Iter R) -> Iter R
func (s Iter[T]) FlatMap[R any](f func(T) Iter[R]) Iter[R] {
	return FlatMap(s, f)
}

// FlatMapI maps each element of a slice along with its index to a slice using the function f,
// and then flattens the result into a single slice.
//
// Type signature:
//
//	FlatMapI :: Iter T -> ((Int, T) -> Iter R) -> Iter R
func (s Iter[T]) FlatMapI[R any](f func(int, T) Iter[R]) Iter[R] {
	return FlatMapI(s, f)
}
