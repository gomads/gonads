package iters

import "github.com/alsi-lawr/gonads/option"

// Find returns the first element in the slice that satisfies the predicate f.
//
// Type signature:
//
//	Find :: Iter T -> (T -> bool) -> Option T
//
// If an element is found, it returns the Some(T); otherwise, it returns None.
func Find[T any](s Iter[T], f func(T) bool) option.Option[T] {
	for _, v := range s {
		if f(v) {
			return option.Some[T](v)
		}
	}
	return option.None[T]()
}

// FindIndex returns the index of the first element in the slice that satisfies the predicate f.
//
// Type signature:
//
//	FindIndex :: Iter T -> (T -> bool) -> Option int
//
// If no element satisfies f, it returns None.
func FindIndex[T any](s Iter[T], f func(T) bool) option.Option[int] {
	for i, v := range s {
		if f(v) {
			return option.Some(i)
		}
	}
	return option.None[int]()
}

// FindFirst returns the first element in the slice that satisfies the predicate f.
//
// Type signature:
//
//	FindFirst :: Iter T -> (T -> bool) -> Option T
//
// If an element is found, it returns Some(T); otherwise, it returns None.
func FindFirst[T any](s Iter[T], f func(T) bool) option.Option[T] {
	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			return option.Some(s[i])
		}
	}
	return option.None[T]()
}

// FindLast returns the last element in the slice that satisfies the predicate f.
//
// Type signature:
//
//	FindLast :: Iter T -> (T -> bool) -> Option T
//
// If an element is found, it returns Some(T); otherwise, it returns None.
func FindLast[T any](s Iter[T], f func(T) bool) option.Option[T] {
	for i := len(s) - 1; i >= 0; i-- {
		if f(s[i]) {
			return option.Some(s[i])
		}
	}
	return option.None[T]()
}

// Any returns true if any element in the slice satisfies the predicate f.
//
// Type signature:
//
//	Any :: Iter T -> (T -> bool) -> bool
func Any[T any](s Iter[T], f func(T) bool) bool {
	for _, v := range s {
		if f(v) {
			return true
		}
	}
	return false
}

// All returns true if every element in the slice satisfies the predicate f.
//
// Type signature:
//
//	All :: Iter T -> (T -> bool) -> bool
func All[T any](s Iter[T], f func(T) bool) bool {
	for _, v := range s {
		if !f(v) {
			return false
		}
	}
	return true
}

// Count returns the number of elements in the slice that satisfy the predicate f.
//
// Type signature:
//
//	Count :: Iter T -> (T -> bool) -> int
func Count[T any](s Iter[T], f func(T) bool) int {
	count := 0
	for _, v := range s {
		if f(v) {
			count++
		}
	}
	return count
}

// Find returns the first element in the slice that satisfies the predicate f.
//
// Type signature:
//
//	Find :: Iter T -> (T -> bool) -> Option T
//
// If an element is found, it returns Some(T); otherwise, it returns None.
func (s Iter[T]) Find(f func(T) bool) option.Option[T] {
	return Find(s, f)
}

// FindIndex returns the index of the first element in the slice that satisfies the predicate f.
//
// Type signature:
//
//	FindIndex :: Iter T -> (T -> bool) -> int
//
// If no element satisfies f, it returns None.
func (s Iter[T]) FindIndex(f func(T) bool) option.Option[int] {
	return FindIndex(s, f)
}

// FindFirst returns the first element in the slice that satisfies the predicate f.
//
// Type signature:
//
//	FindFirst :: Iter T -> (T -> bool) -> Option T
//
// If an element is found, it returns Some(T); otherwise, it returns None.
func (s Iter[T]) FindFirst(f func(T) bool) option.Option[T] {
	return FindFirst(s, f)
}

// FindLast returns the last element in the slice that satisfies the predicate f.
//
// Type signature:
//
//	FindLast :: Iter T -> (T -> bool) -> Option T
//
// If an element is found, it returns Some(T); otherwise, it returns None.
func (s Iter[T]) FindLast(f func(T) bool) option.Option[T] {
	return FindLast(s, f)
}

// Any returns true if any element in the slice satisfies the predicate f.
//
// Type signature:
//
//	Any :: Iter T -> (T -> bool) -> bool
func (s Iter[T]) Any(f func(T) bool) bool {
	return Any(s, f)
}

// All returns true if every element in the slice satisfies the predicate f.
//
// Type signature:
//
//	All :: Iter T -> (T -> bool) -> bool
func (s Iter[T]) All(f func(T) bool) bool {
	return All(s, f)
}

// Count returns the number of elements in the slice that satisfy the predicate f.
//
// Type signature:
//
//	Count :: Iter T -> (T -> bool) -> int
func (s Iter[T]) Count(f func(T) bool) int {
	return Count(s, f)
}
