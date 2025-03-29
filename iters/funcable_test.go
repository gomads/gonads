package iters_test

import "github.com/alsi-lawr/gonads/iters"

func AnyDeepEqual[T comparable](a []any, s iters.Iter[T]) bool {
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
