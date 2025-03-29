package iters

// Filter applies a predicate to each element of a slice, returning a new slice with only the elements that satisfy the predicate.
//
// Type signature:
//
//	Filter :: [T] -> (T -> bool) -> [T]
func Filter[T any](s []T, f func(T) bool) []T {
	var result []T
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// FilterI applies a predicate to each element of a slice along with its index,
// returning a new slice with only the elements that satisfy the predicate.
//
// Type signature:
//
//	FilterI :: [T] -> ((Int, T) -> bool) -> [T]
func FilterI[T any](s []T, f func(int, T) bool) []T {
	var result []T
	for i, v := range s {
		if f(i, v) {
			result = append(result, v)
		}
	}
	return result
}

// FilterMap applies a predicate to each key/value pair of a map,
// returning a new map containing only the pairs that satisfy the predicate.
//
// Type signature:
//
//	FilterMap :: Map K V -> ((K, V) -> bool) -> Map K V
func FilterMap[K comparable, V any](m map[K]V, f func(K, V) bool) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		if f(k, v) {
			result[k] = v
		}
	}
	return result
}

// FilterChan applies a predicate to each element received on a channel,
// returning a new channel that only outputs elements that satisfy the predicate.
//
// Type signature:
//
//	FilterChan :: Channel T -> (T -> bool) -> Channel T
func FilterChan[T any](c <-chan T, f func(T) bool) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for v := range c {
			if f(v) {
				out <- v
			}
		}
	}()
	return out

}

// FilterString applies a predicate to each rune in a string,
// returning a new string containing only the runes that satisfy the predicate.
//
// Type signature:
//
//	FilterString :: String -> (rune -> bool) -> String
func FilterString(s string, f func(rune) bool) string {
	runes := []rune(s)
	var filtered []rune
	for _, r := range runes {
		if f(r) {
			filtered = append(filtered, r)
		}
	}
	return string(filtered)
}

// FilterStringI applies a predicate to each rune in a string along with its index,
// returning a new string containing only the runes that satisfy the predicate.
//
// Type signature:
//
//	FilterStringI :: String -> ((Int, rune) -> bool) -> String
func FilterStringI(s string, f func(int, rune) bool) string {
	runes := []rune(s)
	var filtered []rune
	for i, r := range runes {
		if f(i, r) {

			filtered = append(filtered, r)
		}
	}
	return string(filtered)
}

// Filter applies a predicate to each element of a slice, returning a new slice with only the elements that satisfy the predicate.
//
// Type signature:
//
//	Filter :: [T] -> (T -> bool) -> [T]
func (s Collection[T]) Filter(f func(T) bool) Collection[T] {
	return Filter(s, f)
}

// FilterI applies a predicate to each element of a slice along with its index,
// returning a new slice with only the elements that satisfy the predicate.
//
// Type signature:
//
//	FilterI :: [T] -> ((Int, T) -> bool) -> [T]
func (s Collection[T]) FilterI(f func(int, T) bool) Collection[T] {
	return FilterI(s, f)
}
