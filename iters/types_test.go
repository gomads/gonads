package iters_test

import (
	"reflect"
	"testing"

	"github.com/alsi-lawr/gonads/iters"
)

func TestLiftSlice(t *testing.T) {
	got := iters.LiftSlice([]int{1, 2, 3})
	want := iters.Iter[int]{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("LiftSlice() = %v, want %v", got, want)
	}
}

func TestToSlice(t *testing.T) {
	got := iters.Iter[int]{1, 2, 3}.ToSlice()
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToSlice() = %v, want %v", got, want)
	}
}
