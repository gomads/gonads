package iters

type Iter[T any] []T

type Mappable[T, R any] Iter[T]

type Grouping[K comparable, T any] map[K][]T

type Aggregable[K comparable, T, R any] Grouping[K, T]

func LiftMap[T, R any](data []T) Mappable[T, R] {
	return data
}

func LiftSlice[T any](data []T) Iter[T] {
	return data
}

func LiftAggregable[K comparable, T, R any](data Grouping[K, T]) Aggregable[K, T, R] {
	return (Aggregable[K, T, R])(data)
}

func (s Iter[T]) ToSlice() []T {
	return s
}

func (s Mappable[T, R]) ToIter() Iter[T] {
	return ([]T)(s)
}

func (s Aggregable[K, T, R]) ToGrouping() Grouping[K, T] {
	return (map[K][]T)(s)
}
