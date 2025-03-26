package funcs

type Collection[T any] []T

type Mappable[T, R any] []T

type Grouping[K comparable, T any] map[K][]T

func LiftMap[T, R any](data []T) Mappable[T, R] {
	return data
}

func LiftSlice[T any](data []T) Collection[T] {
	return data
}

func (s Collection[T]) ToSlice() []T {
	return s
}
