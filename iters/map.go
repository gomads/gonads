package iters

// Map applies a function to each element of a slice, returning a new slice with the mapped values.
//
// Type signature:
//
//	Map :: [T] -> (T -> R) -> [R]
//
// Each element in the input slice is transformed using f.
func Map[T any, R any](s Collection[T], f func(T) R) Collection[R] {
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
//	MapSliceWithIndex :: [T] -> ((Int, T) -> R) -> [R]
//
// The function f receives both the index and the element for transformation.
func MapI[T any, R any](s Collection[T], f func(int, T) R) Collection[R] {
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
//	MapSliceErr :: [T] -> (T -> (R, error)) -> ([R], error)
//
// Each element is processed using f, and processing stops if an error is encountered.
func MapErr[T any, R any](s Collection[T], f func(T) (R, error)) (Collection[R], error) {
	result := make([]R, len(s))
	for i, v := range s {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		result[i] = r
	}
	return result, nil
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
//	Map :: [T] -> (T -> R) -> [R]
//
// Each element in the input slice is transformed using f.
func (s Mappable[T, R]) Map(f func(T) R) Collection[R] {
	return Map(s.ToCollection(), f)
}

// MapSliceWithIndex applies a function to each element of a slice along with its index, returning a new slice with the mapped values.
//
// Type signature:
//
//	MapSliceWithIndex :: [T] -> ((Int, T) -> R) -> [R]
//
// The function f receives both the index and the element for transformation.
func (s Mappable[T, R]) MapI(f func(int, T) R) Collection[R] {
	return MapI(s.ToCollection(), f)
}

// MapSliceErr applies a function to each element of a slice, returning a new slice with the mapped values or an error if one occurs.
//
// Type signature:
//
//	MapSliceErr :: [T] -> (T -> (R, error)) -> ([R], error)
//
// Each element is processed using f, and processing stops if an error is encountered.
func (s Mappable[T, R]) MapErr(f func(T) (R, error)) (Collection[R], error) {
	return MapErr(s.ToCollection(), f)
}

// Map applies a function to each element of a slice, returning a new slice with the mapped values.
//
// [Unsafe Variant] This method loses compile-time type safety by returning `[]any` rather than a strongly-typed slice.
// Use only when necessary.
//
// Type signature:
//
//	Map :: [T] -> (T -> R) -> [R]
//
// Each element in the input slice is transformed using f.
func (s Collection[T]) MapUnsafe(f func(T) any) Collection[any] {
	return Map(s, f)
}

// MapSliceWithIndex applies a function to each element of a slice along with its index, returning a new slice with the mapped values.
//
// [Unsafe Variant] This method loses compile-time type safety by returning `[]any` rather than a strongly-typed slice.
// Use only when necessary.
//
// Type signature:
//
//	MapSliceWithIndex :: [T] -> ((Int, T) -> R) -> [R]
//
// The function f receives both the index and the element for transformation.
func (s Collection[T]) MapIUnsafe(f func(int, T) any) Collection[any] {
	return MapI(s, f)
}

// MapSliceErr applies a function to each element of a slice, returning a new slice with the mapped values or an error if one occurs.
//
// [Unsafe Variant] This method loses compile-time type safety by returning `[]any` rather than a strongly-typed slice.
// Use only when necessary.
//
// Type signature:
//
//	MapSliceErr :: [T] -> (T -> (R, error)) -> ([R], error)
//
// Each element is processed using f, and processing stops if an error is encountered.
func (s Collection[T]) MapErrUnsafe(f func(T) (any, error)) (Collection[any], error) {
	return MapErr(s, f)
}
