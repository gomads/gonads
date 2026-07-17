package iters

type Iter[T any] []T

type Grouping[K comparable, T any] map[K][]T

func LiftSlice[T any](data []T) Iter[T] {
	return data
}

func (s Iter[T]) ToSlice() []T {
	return s
}
