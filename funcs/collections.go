package funcs

// Partition splits a slice into two slices based on a predicate.
//
// Type signature:
//
//	Partition :: [T] -> (T -> bool) -> ([T], [T])
//
// Returns two slices: the first containing elements that satisfy f,
// and the second containing elements that do not.
func Partition[T any](s []T, f func(T) bool) ([]T, []T) {
	var yes, no []T
	for _, v := range s {
		if f(v) {
			yes = append(yes, v)
		} else {
			no = append(no, v)
		}
	}
	return yes, no
}

// PartitionI splits a slice into two slices based on a predicate that takes both
// the element's index and value.
//
// Type signature:
//
//	PartitionI :: [T] -> ((Int, T) -> bool) -> ([T], [T])
//
// Returns two slices: the first with elements that satisfy f, and the second with those that don't.
func PartitionI[T any](s []T, f func(int, T) bool) ([]T, []T) {
	var yes, no []T
	for i, v := range s {
		if f(i, v) {
			yes = append(yes, v)
		} else {
			no = append(no, v)
		}
	}
	return yes, no
}

// FlatMap maps each element of a slice to a slice using the function f and then flattens the result
// into a single slice.
//
// Type signature:
//
//	FlatMap :: [T] -> (T -> [R]) -> [R]
func FlatMap[T any, R any](s []T, f func(T) []R) []R {
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
func FlatMapI[T any, R any](s []T, f func(int, T) []R) []R {
	var result []R
	for i, v := range s {
		result = append(result, f(i, v)...)
	}
	return result
}
