package option

import "reflect"

// Option represents a type that may or may not contain a value (Some or None).
// It is used to model the presence (Some) or absence (None) of a value.
//
// Type signature:
//
//	Option[T] :: a -> Option a
type Option[T any] struct {
	value  T
	isSome bool
}

// Wrap lifts a nullable pointer (*T) into an Option[T].
//
// Type signature:
//
//	Wrap :: Maybe a -> Option a
//
// If the input is nil, it returns None. Otherwise, it returns Some with the dereferenced value.
func Wrap[T any](val *T) Option[T] {
	if val == nil {
		return None[T]()
	}
	return Some(*val)
}

// Some creates an Option containing a value.
//
// Type signature:
//
//	Some :: a -> Option a
//
// It returns an Option with the provided value set and isSome as true.
func Some[T any](val T) Option[T] {
	return Option[T]{value: val, isSome: true}
}

// None creates an Option with no value (None).
//
// Type signature:
//
//	None :: Option a
//
// It returns an Option where isSome is false, representing the absence of a value.
func None[T any]() Option[T] {
	return Option[T]{isSome: false}
}

// IsSome returns true if the Option contains a value (Some), false otherwise.
//
// Type signature:
//
//	IsSome :: Option a -> Bool
func (opt Option[T]) IsSome() bool {
	return opt.isSome
}

// IsNone returns true if the Option contains no value (None), false otherwise.
//
// Type signature:
//
//	IsNone :: Option a -> Bool
func (opt Option[T]) IsNone() bool {
	return !opt.isSome
}

// Match applies one of two functions depending on whether the Option contains a value (Some) or not (None).
//
// Type signature:
//
//	Match :: Option a -> (a -> ()) -> (() -> ()) -> ()
//
// If isSome is true, it invokes the some function with the value. Otherwise, it invokes the none function.
func (opt Option[T]) Match(some func(T), none func()) {
	if opt.isSome {
		some(opt.value)
		return
	}
	none()
}

// GetOrNil returns the value inside the Option if it exists(Some).
//
// Type signature:
//
//	GetOrNil :: Option a -> a
//
// If the Option is None, it returns nil.
func (opt Option[T]) GetOrNil() *T {
	if opt.isSome {
		return &opt.value
	}
	return nil
}

// GetOrElse returns the value inside the Option if it exists (Some).
//
// Type signature:
//
//	GetOrElse :: Option a -> (() -> a) -> a
//
// If the Option is None, it calls the provided function to return a default value.
func (opt Option[T]) GetOrElse(fn func() T) T {
	if opt.isSome {
		return opt.value
	}
	return fn()
}

// OrElse returns the Option if it contains a value (Some).
//
// Type signature:
//
//	OrElse :: Option a -> (() -> Option a) -> Option a
//
// Otherwise, it calls the provided function and returns its result.
func (opt Option[T]) OrElse(fn func() Option[T]) Option[T] {
	if opt.isSome {
		return opt
	}
	return fn()
}

// Flatten removes one level of nesting in an Option containing another Option.
//
// Type signature:
//
//	Flatten :: Option (Option a) -> Option a
//
// If the outer Option is Some, it returns the inner Option. If the outer Option is None, it returns None.
func Flatten[T any](opt Option[Option[T]]) Option[T] {
	if opt.isSome {
		return opt.value
	}
	return None[T]()
}

// ToSlice converts an Option to a slice containing one element if the Option is Some, or an empty slice if None.
//
// Type signature:
//
//	ToSlice :: Option a -> [a]
func (opt Option[T]) ToSlice() []T {
	if opt.isSome {
		return []T{opt.value}
	}
	return []T{}
}

// Zip combines two Option values into one Option containing a struct with both values, if both are Some.
//
// Type signature:
//
//	Zip :: Option a -> Option b -> Option (a, b)
//
// If either Option is None, it returns None.
func Zip[T, U any](opt1 Option[T], opt2 Option[U]) Option[struct {
	First  T
	Second U
}] {
	if opt1.isSome && opt2.isSome {
		return Some(struct {
			First  T
			Second U
		}{First: opt1.value, Second: opt2.value})
	}
	return None[struct {
		First  T
		Second U
	}]()
}

// Bind applies a function to the value inside the Option, if it exists (Some).
//
// Type signature:
//
//	Bind :: Option a -> (a -> Option b) -> Option b
//
// If the Option is None, it returns None without applying the function.
func Bind[T, U any](opt Option[T], fn func(T) Option[U]) Option[U] {
	if opt.isSome {
		return fn(opt.value)
	}
	return None[U]()
}

// Bind applies a function to the value inside the Option, if it exists (Some).
//
// Type signature:
//
//	Bind :: Option a -> (a -> Option a) -> Option a
//
// If the Option is None, it returns None without applying the function.
func (opt Option[T]) Bind(fn func(T) Option[T]) Option[T] {
	return Bind(opt, fn)
}

// Map applies a function to the value inside the Option, if it exists (Some).
//
// Type signature:
//
//	Map :: Option a -> (a -> b) -> Option b
//
// It returns a new Option with the result of the function, or None if the original Option is None.
func Map[T, U any](opt Option[T], fn func(T) U) Option[U] {
	if opt.isSome {
		return Some(fn(opt.value))
	}
	return None[U]()
}

// BiMap applies one of two functions depending on whether the Option contains a value (Some) or not (None).
//
// Type signature:
//
//	BiMap :: Option a -> (a -> b) -> (() -> b) -> b
//
// If the Option is Some, it applies the some function to the value and returns the result.
// If the Option is None, it calls the none function and returns its result.
func BiMap[T, U any](opt Option[T], ifSome func(T) U, ifNone func() U) U {
	if opt.isSome {
		return ifSome(opt.value)
	}
	return ifNone()
}

// Filter returns the Option itself if it contains a value and the provided function returns true for that value.
//
// Type signature:
//
//	Filter :: Option a -> (a -> Bool) -> Option a
//
// If the Option is None or the function returns false, it returns None.
func Filter[T any](opt Option[T], fn func(T) bool) Option[T] {
	if opt.isSome && fn(opt.value) {
		return opt
	}
	return None[T]()
}

// Try applies a function that returns a value and an error.
//
// Type signature:
//
//	Try :: (() -> (a, Error)) -> Option a
//
// If the function returns an error, it converts the result into None. Otherwise, it returns an Option containing the value (Some).
func Try[T any](fn func() (T, error)) Option[T] {
	val, err := fn()
	if err != nil {
		return None[T]()
	}
	return Some(val)
}

// Equals checks if two Option values are equal.
//
// Type signature:
//
//	Equals :: Option a -> Option a -> Bool
//
// It compares the values if both Options are Some, using reflect.DeepEqual to handle complex types.
// If both Options are None, it returns true. Otherwise, it returns false.
func (opt Option[T]) Equals(other Option[T]) bool {
	if opt.isSome != other.isSome {
		return false
	}

	if !opt.isSome && !other.isSome {
		return true // none is equal to none
	}

	return reflect.DeepEqual(opt.value, other.value) // T and T are not comparable with ==
}
