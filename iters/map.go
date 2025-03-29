package iters

import "github.com/alsi-lawr/gonads/result"

// Map applies a function to each element of a slice, returning a new slice with the mapped values.
//
// Type signature:
//
//	Map :: Iter T -> (T -> R) -> Iter R
//
// Each element in the input slice is transformed using f.
func Map[T any, R any](s Iter[T], f func(T) R) Iter[R] {
	result := make([]R, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

// MapSliceWithIndex applies a function to each element of a slice along with its index, returning a new slice with the mapped values.
//
// Type signature:
//
//	MapSliceWithIndex :: Iter T -> ((Int, T) -> R) -> Iter R
//
// The function f receives both the index and the element for transformation.
func MapI[T any, R any](s Iter[T], f func(int, T) R) Iter[R] {
	result := make([]R, len(s))
	for i, v := range s {
		result[i] = f(i, v)
	}
	return result
}

// MapSliceErr applies a function to each element of a slice, returning a new slice with the mapped values or an error if one occurs.
//
// Type signature:
//
//	MapSliceErr :: Iter T -> (T -> (R, error)) -> Result Iter R
//
// Each element is processed using f, and processing stops if an error is encountered.
func MapErr[T any, R any](s Iter[T], f func(T) (R, error)) result.Result[Iter[R]] {
	mapRes := make(Iter[R], len(s))
	for i, v := range s {
		r, err := f(v)
		if err != nil {
			return result.Err[Iter[R]](err)
		}
		mapRes[i] = r
	}
	return result.Ok[Iter[R]](mapRes)
}

// MapMap applies a function to each key/value pair of a map, returning a new map with transformed keys and values.
//
// Type signature:
//
//	MapMap :: Map K V -> ((K, V) -> (NK, NV)) -> Map NK NV
//
// Each key/value pair is processed using f to build the new map.
func MapMap[K comparable, V any, NK comparable, NV any](
	m map[K]V,
	f func(K, V) (NK, NV),
) map[NK]NV {
	result := make(map[NK]NV, len(m))
	for k, v := range m {
		nk, nv := f(k, v)
		result[nk] = nv
	}
	return result
}

// MapMapKeys applies a function to each key of a map, returning a new map with transformed keys and original values.
//
// Type signature:
//
//	MapMapKeys :: Map K V -> (K -> NK) -> Map NK V
//
// Only the keys are transformed by f; the values remain unchanged.
func MapMapKeys[K comparable, V any, NK comparable](m map[K]V, f func(K) NK) map[NK]V {
	result := make(map[NK]V, len(m))
	for k, v := range m {
		result[f(k)] = v
	}
	return result
}

// MapMapValues applies a function to each value of a map, returning a new map with original keys and transformed values.
//
// Type signature:
//
//	MapMapValues :: Map K V -> (V -> NV) -> Map K NV
//
// Only the values are transformed by f; the keys remain unchanged.
func MapMapValues[K comparable, V any, NV any](m map[K]V, f func(V) NV) map[K]NV {
	result := make(map[K]NV, len(m))
	for k, v := range m {

		result[k] = f(v)
	}
	return result
}

// MapChan applies a function to each element received on a channel, returning a new channel with the mapped values.
//
// Type signature:
//
//	MapChan :: Channel T -> (T -> R) -> Channel R
//
// Each element from the input channel is processed by f and sent to the output channel.
func MapChan[T any, R any](c <-chan T, f func(T) R) <-chan R {
	out := make(chan R)
	go func() {
		defer close(out)
		for v := range c {
			out <- f(v)
		}
	}()
	return out
}

// MapString applies a function to each rune in a string, returning a new string with the mapped runes.
//
// Type signature:
//
//	MapString :: String -> (rune -> rune) -> String
//
// Each rune in the input string is transformed using f.
func MapString(s string, f func(rune) rune) string {
	runes := []rune(s)
	for i, r := range runes {
		runes[i] = f(r)
	}
	return string(runes)
}

// MapStringI applies a function to each rune in a string along with its index, returning a new string with the mapped runes.
//
// Type signature:
//
//	MapStringI :: String -> ((Int, rune) -> rune) -> String
//
// The function f receives both the index and the rune for transformation.
func MapStringI(s string, f func(int, rune) rune) string {
	runes := []rune(s)
	for i, r := range runes {
		runes[i] = f(i, r)
	}
	return string(runes)
}

// Map applies a function to each element of a slice, returning a new slice with the mapped values.
//
// Type signature:
//
//	Map :: Mappable T -> (T -> R) -> Iter R
//
// Each element in the input slice is transformed using f.
func (s Mappable[T, R]) Map(f func(T) R) Iter[R] {
	return Map(s.ToIter(), f)
}

// MapSliceWithIndex applies a function to each element of a slice along with its index, returning a new slice with the mapped values.
//
// Type signature:
//
//	MapSliceWithIndex :: Mappable T -> ((Int, T) -> R) -> Iter R
//
// The function f receives both the index and the element for transformation.
func (s Mappable[T, R]) MapI(f func(int, T) R) Iter[R] {
	return MapI(s.ToIter(), f)
}

// MapSliceErr applies a function to each element of a slice, returning a new slice with the mapped values or an error if one occurs.
//
// Type signature:
//
//	MapSliceErr :: Mappable T -> (T -> (R, error)) -> (Iter R, error)
//
// Each element is processed using f, and processing stops if an error is encountered.
func (s Mappable[T, R]) MapErr(f func(T) (R, error)) result.Result[Iter[R]] {
	return MapErr(s.ToIter(), f)
}

// Map applies a function to each element of a slice, returning a new slice with the mapped values.
//
// [Unsafe Variant] This method loses compile-time type safety by returning `[]any` rather than a strongly-typed slice.
// Use only when necessary.
//
// Type signature:
//
//	Map :: Iter T -> (T -> any) -> Iter any
//
// Each element in the input slice is transformed using f.
func (s Iter[T]) MapUnsafe(f func(T) any) Iter[any] {
	return Map(s, f)
}

// MapSliceWithIndex applies a function to each element of a slice along with its index, returning a new slice with the mapped values.
//
// [Unsafe Variant] This method loses compile-time type safety by returning `[]any` rather than a strongly-typed slice.
// Use only when necessary.
//
// Type signature:
//
//	MapSliceWithIndex :: Iter T -> ((Int, T) -> any) -> Iter any
//
// The function f receives both the index and the element for transformation.
func (s Iter[T]) MapIUnsafe(f func(int, T) any) Iter[any] {
	return MapI(s, f)
}

// MapSliceErr applies a function to each element of a slice, returning a new slice with the mapped values or an error if one occurs.
//
// [Unsafe Variant] This method loses compile-time type safety by returning `[]any` rather than a strongly-typed slice.
// Use only when necessary.
//
// Type signature:
//
//	MapSliceErr :: Iter T -> (T -> (any, error)) -> Result Iter any
//
// Each element is processed using f, and processing stops if an error is encountered.
func (s Iter[T]) MapErrUnsafe(f func(T) (any, error)) result.Result[Iter[any]] {
	return MapErr(s, f)
}
