package iters

// Partition splits a slice into two slices based on a predicate.
//
// Type signature:
//
//	Partition :: Iter T -> (T -> bool) -> (Iter T, Iter T)
//
// Returns two slices: the first containing elements that satisfy f,
// and the second containing elements that do not.
func Partition[T any](s Iter[T], f func(T) bool) (Iter[T], Iter[T]) {
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
//	PartitionI :: Iter T -> ((Int, T) -> bool) -> (Iter T, Iter T)
//
// Returns two slices: the first with elements that satisfy f, and the second with those that don't.
func PartitionI[T any](s Iter[T], f func(int, T) bool) (Iter[T], Iter[T]) {
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

// Partition splits a slice into two slices based on a predicate.
//
// Type signature:
//
//	Partition :: Iter T -> (T -> bool) -> (Iter T, Iter T)
//
// Returns two slices: the first containing elements that satisfy f,
// and the second containing elements that do not.
func (s Iter[T]) Partition(f func(T) bool) (Iter[T], Iter[T]) {
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
//	PartitionI :: Iter T -> ((Int, T) -> bool) -> (Iter T, Iter T)
//
// Returns two slices: the first with elements that satisfy f, and the second with those that don't.
func (s Iter[T]) PartitionI(f func(int, T) bool) (Iter[T], Iter[T]) {
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
