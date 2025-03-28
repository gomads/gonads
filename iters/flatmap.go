package iters

// FlatMap maps each element of a slice to a slice using the function f and then flattens the result
// into a single slice.
//
// Type signature:
//
//	FlatMap :: [T] -> (T -> [R]) -> [R]
func FlatMap[T any, R any](s Collection[T], f func(T) Collection[R]) Collection[R] {
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
//	FlatMapI :: [T] -> ((Int, T) -> [R]) -> [R]
func FlatMapI[T any, R any](s Collection[T], f func(int, T) Collection[R]) Collection[R] {
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
//	FlatMap :: [T] -> (T -> [R]) -> [R]
func (s Mappable[T, R]) FlatMap(f func(T) Collection[R]) Collection[R] {
	return FlatMap(s.ToCollection(), f)
}

// FlatMapI maps each element of a slice along with its index to a slice using the function f,
// and then flattens the result into a single slice.
//
// Type signature:
//
//	FlatMapI :: [T] -> ((Int, T) -> [R]) -> [R]
func (s Mappable[T, R]) FlatMapI(f func(int, T) Collection[R]) Collection[R] {
	return FlatMapI(s.ToCollection(), f)
}

// FlatMap maps each element of a slice to a slice using the function f and then flattens the result
// into a single slice.
//
// [Unsafe Variant] This method loses compile-time type safety by returning `[]any` rather than a strongly-typed slice.
// Use only when necessary.
//
// Type signature:
//
//	FlatMap :: [T] -> (T -> [R]) -> [R]
func (s Collection[T]) FlatMapUnsafe(f func(T) Collection[any]) Collection[any] {
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
//	FlatMapI :: [T] -> ((Int, T) -> [R]) -> [R]
func (s Collection[T]) FlatMapIUnsafe(f func(int, T) Collection[any]) Collection[any] {
	return FlatMapI(s, f)
}
