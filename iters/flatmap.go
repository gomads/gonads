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
func (s Mappable[T, R]) FlatMap(f func(T) Iter[R]) Iter[R] {
	return FlatMap(s.ToIter(), f)
}

// FlatMapI maps each element of a slice along with its index to a slice using the function f,
// and then flattens the result into a single slice.
//
// Type signature:
//
//	FlatMapI :: Mappable T R -> ((Int, T) -> Iter R) -> Iter R
func (s Mappable[T, R]) FlatMapI(f func(int, T) Iter[R]) Iter[R] {
	return FlatMapI(s.ToIter(), f)
}

// FlatMap maps each element of a slice to a slice using the function f and then flattens the result
// into a single slice.
//
// [Unsafe Variant] This method loses compile-time type safety by returning `[]any` rather than a strongly-typed slice.
// Use only when necessary.
//
// Type signature:
//
//	FlatMap :: Iter T -> (T -> Iter any) -> Iter any
func (s Iter[T]) FlatMapUnsafe(f func(T) Iter[any]) Iter[any] {
	return FlatMap(s, f)
}

// FlatMapI maps each element of a slice along with its index to a slice using the function f,
// and then flattens the result into a single slice.
//
// [Unsafe Variant] This method loses compile-time type safety by returning `[]any` rather than a strongly-typed slice.
// Use only when necessary.
//
// Type signature:
//
//	FlatMapI :: Iter T -> ((Int, T) -> Iter any) -> Iter any
func (s Iter[T]) FlatMapIUnsafe(f func(int, T) Iter[any]) Iter[any] {
	return FlatMapI(s, f)
}
