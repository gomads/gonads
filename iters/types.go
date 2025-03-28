package iters

type Collection[T any] []T

type Mappable[T, R any] Collection[T]

type Grouping[K comparable, T any] map[K][]T

type Aggregable[K comparable, T, R any] Grouping[K, T]

func LiftMap[T, R any](data []T) Mappable[T, R] {
	return data
}

func LiftSlice[T any](data []T) Collection[T] {
	return data
}

func LiftAggregable[K comparable, T, R any](data Grouping[K, T]) Aggregable[K, T, R] {
	return (Aggregable[K, T, R])(data)
}

func (s Collection[T]) ToSlice() []T {
	return s
}

func (s Mappable[T, R]) ToCollection() Collection[T] {
	return ([]T)(s)
}

func (s Aggregable[K, T, R]) ToGrouping() Grouping[K, T] {
	return (map[K][]T)(s)
}
