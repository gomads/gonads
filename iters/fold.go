package iters

// Fold applies a function to each element of a slice, reducing it to a single value.
//
// Type signature:
//
//	Fold :: [T] -> A -> ((A, T) -> A) -> A
func Fold[T any, A any](s []T, init A, f func(A, T) A) A {
	acc := init
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

// FoldI applies a function to each element of a slice with its index, reducing it to a single value.
//
// Type signature:
//
//	FoldI :: [T] -> A -> ((Int, A, T) -> A) -> A
func FoldI[T any, A any](s []T, init A, f func(int, A, T) A) A {
	acc := init
	for i, v := range s {
		acc = f(i, acc, v)
	}
	return acc
}

// FoldMap applies a function to each key/value pair in a map, reducing it to a single value.
// Note: The order of iteration over a map is not guaranteed.
//
// Type signature:
//
//	FoldMap :: Map K V -> A -> ((A, K, V) -> A) -> A
func FoldMap[K comparable, V any, A any](m map[K]V, init A, f func(A, K, V) A) A {
	acc := init
	for k, v := range m {
		acc = f(acc, k, v)
	}
	return acc
}

// FoldChan applies a function to each element received from a channel, reducing it to a single value.
//
// Type signature:
//
//	FoldChan :: Channel T -> A -> ((A, T) -> A) -> A
func FoldChan[T any, A any](c <-chan T, init A, f func(A, T) A) A {
	acc := init
	for v := range c {
		acc = f(acc, v)
	}
	return acc
}

// FoldString applies a function to each rune in a string, reducing it to a single value.
//
// Type signature:
//
//	FoldString :: String -> A -> ((A, rune) -> A) -> A
func FoldString[A any](s string, init A, f func(A, rune) A) A {
	acc := init
	for _, r := range s {
		acc = f(acc, r)
	}
	return acc
}

// FoldStringI applies a function to each rune in a string along with its index, reducing it to a single value.
//
// Type signature:
//
//	FoldStringI :: String -> A -> ((Int, A, rune) -> A) -> A
func FoldStringI[A any](s string, init A, f func(int, A, rune) A) A {
	acc := init
	for i, r := range s {
		acc = f(i, acc, r)
	}
	return acc
}

// Fold applies a function to each element of a slice, reducing it to a single value.
//
// Type signature:
//
//	Fold :: [T] -> A -> ((A, T) -> A) -> A
func (s Mappable[T, A]) Fold(init A, f func(A, T) A) A {
	return Fold(s, init, f)
}

// FoldI applies a function to each element of a slice with its index, reducing it to a single value.
//
// Type signature:
//
//	FoldI :: [T] -> A -> ((Int, A, T) -> A) -> A
func (s Mappable[T, A]) FoldI(init A, f func(int, A, T) A) A {
	return FoldI(s, init, f)
}

// Fold applies a function to each element of a slice, reducing it to a single value.
//
// Type signature:
//
//	Fold :: [T] -> A -> ((A, T) -> A) -> A
func (s Iter[T]) FoldUnsafe(init any, f func(any, T) any) any {
	return Fold(s, init, f)
}

// FoldI applies a function to each element of a slice with its index, reducing it to a single value.
//
// Type signature:
//
//	FoldI :: [T] -> A -> ((Int, A, T) -> A) -> A
func (s Iter[T]) FoldIUnsafe(init any, f func(int, any, T) any) any {
	return FoldI(s, init, f)
}
