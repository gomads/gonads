package iters_test

import (
	"reflect"
	"testing"

	"github.com/alsi-lawr/gonads/iters"
)

func TestLiftSlice(t *testing.T) {
	got := iters.LiftSlice([]int{1, 2, 3})
	want := iters.Collection[int]{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("LiftSlice() = %v, want %v", got, want)
	}
}

func TestLiftMap(t *testing.T) {
	got := iters.LiftMap[int, int]([]int{1, 2, 3})
	want := iters.Mappable[int, int]{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("LiftMap() = %v, want %v", got, want)
	}
}
