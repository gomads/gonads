package funcs_test

import "github.com/alsi-lawr/gonads/funcs"

func AnyDeepEqual[T comparable](a []any, s funcs.Collection[T]) bool {
	if len(a) != len(s) {
		return false
	}
	for i, v := range a {
		tv, ok := v.(T)
		if !ok || tv != s[i] {
			return false
		}
	}
	return true
}
