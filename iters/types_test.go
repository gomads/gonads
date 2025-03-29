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

func TestLiftAggregable(t *testing.T) {
	got := iters.LiftAggregable[int, int, int](
		map[int][]int{
			1: {1, 2, 3},
			2: {4, 5, 6},
		})
	want := iters.Aggregable[int, int, int]{
		1: {1, 2, 3},
		2: {4, 5, 6},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("LiftAggregable() = %v, want %v", got, want)
	}
}

func TestToSlice(t *testing.T) {
	got := iters.Collection[int]{1, 2, 3}.ToSlice()
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToSlice() = %v, want %v", got, want)
	}
}

func TestToCollection(t *testing.T) {
	got := iters.Mappable[int, int]{1, 2, 3}.ToCollection()
	want := iters.Collection[int]{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToCollection() = %v, want %v", got, want)
	}
}

func TestToGrouping(t *testing.T) {
	got := iters.Aggregable[int, int, int]{
		1: {1, 2, 3},
		2: {4, 5, 6},
	}.ToGrouping()
	want := iters.Grouping[int, int]{
		1: {1, 2, 3},
		2: {4, 5, 6},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToGrouping() = %v, want %v", got, want)
	}
}
