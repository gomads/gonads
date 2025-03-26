package funcs

// Find returns the first element in the slice that satisfies the predicate f.
//
// Type signature:
//
//	Find :: [T] -> (T -> bool) -> (T, bool)
//
// If an element is found, it returns the element and true; otherwise, it returns the zero value of T and false.
func Find[T any](s []T, f func(T) bool) (T, bool) {
	for _, v := range s {
		if f(v) {
			return v, true
		}
	}
	var zero T
	return zero, false
}

// FindIndex returns the index of the first element in the slice that satisfies the predicate f.
//
// Type signature:
//
//	FindIndex :: [T] -> (T -> bool) -> int
//
// If no element satisfies f, it returns -1.
func FindIndex[T any](s []T, f func(T) bool) int {
	for i, v := range s {
		if f(v) {
			return i
		}
	}
	return -1
}

// FindFirst returns the first element in the slice that satisfies the predicate f.
//
// Type signature:
//
//	FindFirst :: [T] -> (T -> bool) -> (T, bool)
//
// If an element is found, it returns the element and true; otherwise, it returns the zero value of T and false.
func FindFirst[T any](s []T, f func(T) bool) (T, bool) {
	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			return s[i], true
		}
	}

	var zero T
	return zero, false
}

// FindLast returns the last element in the slice that satisfies the predicate f.
//
// Type signature:
//
//	FindLast :: [T] -> (T -> bool) -> (T, bool)
//
// If an element is found, it returns the element and true; otherwise, it returns the zero value of T and false.
func FindLast[T any](s []T, f func(T) bool) (T, bool) {
	for i := len(s) - 1; i >= 0; i-- {
		if f(s[i]) {
			return s[i], true
		}
	}

	var zero T
	return zero, false
}

// Some returns true if any element in the slice satisfies the predicate f.
//
// Type signature:
//
//	Some :: [T] -> (T -> bool) -> bool
func Some[T any](s []T, f func(T) bool) bool {
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
//	All :: [T] -> (T -> bool) -> bool
func All[T any](s []T, f func(T) bool) bool {
	for _, v := range s {
		if !f(v) {
			return false
		}
	}
	return true
}

// None returns true if no element in the slice satisfies the predicate f.
//
// Type signature:
//
//	None :: [T] -> (T -> bool) -> bool
func None[T any](s []T, f func(T) bool) bool {
	for _, v := range s {
		if f(v) {
			return false
		}
	}
	return true
}

// Count returns the number of elements in the slice that satisfy the predicate f.
//
// Type signature:
//
//	Count :: [T] -> (T -> bool) -> int
func Count[T any](s []T, f func(T) bool) int {
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
//	Find :: [T] -> (T -> bool) -> (T, bool)
//
// If an element is found, it returns the element and true; otherwise, it returns the zero value of T and false.
func (s Collection[T]) Find(f func(T) bool) (T, bool) {
	return Find(s, f)
}

// FindIndex returns the index of the first element in the slice that satisfies the predicate f.
//
// Type signature:
//
//	FindIndex :: [T] -> (T -> bool) -> int
//
// If no element satisfies f, it returns -1.
func (s Collection[T]) FindIndex(f func(T) bool) int {
	return FindIndex(s, f)
}

// FindFirst returns the first element in the slice that satisfies the predicate f.
//
// Type signature:
//
//	FindFirst :: [T] -> (T -> bool) -> (T, bool)
//
// If an element is found, it returns the element and true; otherwise, it returns the zero value of T and false.
func (s Collection[T]) FindFirst(f func(T) bool) (T, bool) {
	return FindFirst(s, f)
}

// FindLast returns the last element in the slice that satisfies the predicate f.
//
// Type signature:
//
//	FindLast :: [T] -> (T -> bool) -> (T, bool)
//
// If an element is found, it returns the element and true; otherwise, it returns the zero value of T and false.
func (s Collection[T]) FindLast(f func(T) bool) (T, bool) {
	return FindLast(s, f)
}

// Some returns true if any element in the slice satisfies the predicate f.
//
// Type signature:
//
//	Some :: [T] -> (T -> bool) -> bool
func (s Collection[T]) Some(f func(T) bool) bool {
	return Some(s, f)
}

// All returns true if every element in the slice satisfies the predicate f.
//
// Type signature:
//
//	All :: [T] -> (T -> bool) -> bool
func (s Collection[T]) All(f func(T) bool) bool {
	return All(s, f)
}

// None returns true if no element in the slice satisfies the predicate f.
//
// Type signature:
//
//	None :: [T] -> (T -> bool) -> bool
func (s Collection[T]) None(f func(T) bool) bool {
	return None(s, f)
}

// Count returns the number of elements in the slice that satisfy the predicate f.
//
// Type signature:
//
//	Count :: [T] -> (T -> bool) -> int
func (s Collection[T]) Count(f func(T) bool) int {
	return Count(s, f)
}
