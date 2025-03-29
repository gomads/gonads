package iters

// GroupBy groups elements of a slice into a map keyed by the output of the key function f.
//
// Type signature:
//
//	GroupBy :: [T] -> (T -> K) -> Map K [T]
//
// Returns a map where each key is produced by f and the corresponding value is a slice of all
// elements for which f returned that key.
func GroupBy[T any, K comparable](s Iter[T], f func(T) K) Grouping[K, T] {
	groups := make(map[K][]T)
	for _, v := range s {
		key := f(v)
		groups[key] = append(groups[key], v)
	}
	return groups
}

// GroupByI groups elements of a slice into a map keyed by a function that takes the element's
// index and value.
//
// Type signature:
//
//	GroupByI :: [T] -> ((Int, T) -> K) -> Map K [T]
//
// Returns a map where each key is produced by f and the corresponding value is a slice of all
// elements for which f returned that key.
func GroupByI[T any, K comparable](s Iter[T], f func(int, T) K) Grouping[K, T] {
	groups := make(map[K][]T)
	for i, v := range s {
		key := f(i, v)
		groups[key] = append(groups[key], v)
	}
	return groups
}

// GroupBy groups elements of a slice into a map keyed by the output of the key function f.
//
// Type signature:
//
//	GroupBy :: [T] -> (T -> K) -> Map K [T]
//
// Returns a map where each key is produced by f and the corresponding value is a slice of all
// elements for which f returned that key.
func (s Iter[T]) GroupByUnsafe(f func(T) any) Grouping[any, T] {
	return GroupBy(s, f)
}

// GroupByI groups elements of a slice into a map keyed by a function that takes the element's
// index and value.
//
// [Unsafe Variant] This method loses compile-time type safety by returning `any` rather than a strongly-typed grouping.
// Use only when necessary.
//
// Type signature:
//
//	GroupByI :: [T] -> ((Int, T) -> K) -> Map K [T]
//
// Returns a map where each key is produced by f and the corresponding value is a slice of all
// elements for which f returned that key.
func (s Iter[T]) GroupByIUnsafe(f func(int, T) any) Grouping[any, T] {
	return GroupByI(s, f)
}
