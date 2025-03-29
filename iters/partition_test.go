package iters_test

import (
	"reflect"
	"testing"

	"github.com/alsi-lawr/gonads/iters"
)

func TestPartitionStatic(t *testing.T) {
	ints := []int{1, 2, 3, 4}
	got1, got2 := iters.Partition(ints, func(x int) bool { return x%2 == 0 })
	want1, want2 := iters.Iter[int]{2, 4}, iters.Iter[int]{1, 3}
	if !reflect.DeepEqual(got1, want1) || !reflect.DeepEqual(got2, want2) {
		t.Errorf(
			"funcs.Partition() = [%v, %v], want [%v, %v]",
			got1,
			got2,
			want1,
			want2,
		)
	}
}

func TestPartitionIStatic(t *testing.T) {
	input := []int{1, 2, 3, 4}
	got1, got2 := iters.PartitionI(input, func(idx, x int) bool { return x%2 == 0 || idx == 0 })
	want1, want2 := iters.Iter[int]{1, 2, 4}, iters.Iter[int]{3}
	if !reflect.DeepEqual(got1, want1) || !reflect.DeepEqual(got2, want2) {
		t.Errorf(
			"funcs.PartitionI() = [%v, %v], want [%v, %v]",
			got1,
			got2,
			want1,
			want2,
		)
	}
}

func TestPartition(t *testing.T) {
	ints := iters.Iter[int]{1, 2, 3, 4}
	got1, got2 := ints.Partition(func(x int) bool { return x%2 == 0 })
	want1, want2 := iters.Iter[int]{2, 4}, iters.Iter[int]{1, 3}
	if !reflect.DeepEqual(got1, want1) || !reflect.DeepEqual(got2, want2) {
		t.Errorf(
			"funcs.Partition() = [%v, %v], want [%v, %v]",
			got1,
			got2,
			want1,
			want2,
		)
	}
}

func TestPartitionI(t *testing.T) {
	input := iters.Iter[int]{1, 2, 3, 4}
	got1, got2 := input.PartitionI(func(idx, x int) bool { return x%2 == 0 || idx == 0 })
	want1, want2 := iters.Iter[int]{1, 2, 4}, iters.Iter[int]{3}
	if !reflect.DeepEqual(got1, want1) || !reflect.DeepEqual(got2, want2) {
		t.Errorf(
			"funcs.PartitionI() = [%v, %v], want [%v, %v]",
			got1,
			got2,
			want1,
			want2,
		)
	}
}
