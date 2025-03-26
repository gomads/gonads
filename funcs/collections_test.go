package funcs_test

import (
	"reflect"
	"testing"

	"github.com/alsi-lawr/gonads/funcs"
)

func TestPartition(t *testing.T) {
	ints := []int{1, 2, 3, 4}
	got1, got2 := funcs.Partition(ints, func(x int) bool { return x%2 == 0 })
	want1, want2 := []int{2, 4}, []int{1, 3}
	if len(got1) != 2 || len(got2) != 2 {
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
	input := []int{1, 2, 3, 4}
	got1, got2 := funcs.PartitionI(input, func(idx, x int) bool { return x%2 == 0 || idx == 0 })
	want1, want2 := []int{1, 3, 4}, []int{2}
	if len(got1) != 3 || len(got2) != 1 {
		t.Errorf(
			"funcs.PartitionI() = [%v, %v], want [%v, %v]",
			got1,
			got2,
			want1,
			want2,
		)
	}
}

func TestFlatMap(t *testing.T) {
	input := []int{1, 2, 3}
	got := funcs.FlatMap(input, func(x int) []int {
		return []int{x, x * 10}
	})
	want := []int{1, 10, 2, 20, 3, 30}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("FlatMap() = %v, want %v", got, want)
	}
}

func TestFlatMapI(t *testing.T) {
	input := []string{"a", "b", "c"}
	got := funcs.FlatMapI(input, func(i int, s string) []string {
		return []string{s, string(rune('A' + i))}
	})
	want := []string{"a", "A", "b", "B", "c", "C"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FlatMapI() = %v, want %v", got, want)
	}
}

func TestGroupByStatic(t *testing.T) {
	input := funcs.LiftSlice([]int{1, 2, 3, 4, 5, 6})
	got := funcs.GroupBy(input, func(x int) bool {
		return x%2 == 0
	})
	want := funcs.Grouping[bool, int]{
		false: {1, 3, 5},
		true:  {2, 4, 6},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GroupBy() = %v, want %v", got, want)
	}
}

func TestGroupByIStatic(t *testing.T) {
	input := funcs.LiftSlice([]string{"apple", "banana", "apricot", "blueberry", "avocado"})
	got := funcs.GroupByI(input, func(i int, s string) string {
		if i%2 == 0 {
			return "even"
		}
		return "odd"
	})
	want := funcs.Grouping[string, string]{
		"even": {"apple", "apricot", "avocado"},
		"odd":  {"banana", "blueberry"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GroupByI() = %v, want %v", got, want)
	}
}

func TestGroupByUnsafe(t *testing.T) {
	input := funcs.LiftSlice([]int{1, 2, 3, 4, 5, 6})
	got := input.GroupByUnsafe(func(x int) any {
		return x%2 == 0
	})
	want := funcs.Grouping[any, int]{
		false: {1, 3, 5},
		true:  {2, 4, 6},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GroupBy() = %v, want %v", got, want)
	}
}

func TestGroupByIUnsafe(t *testing.T) {
	input := funcs.LiftSlice([]string{"apple", "banana", "apricot", "blueberry", "avocado"})
	got := input.GroupByIUnsafe(func(i int, s string) any {
		if i%2 == 0 {
			return "even"
		}
		return "odd"
	})
	want := funcs.Grouping[any, string]{
		"even": {"apple", "apricot", "avocado"},
		"odd":  {"banana", "blueberry"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GroupByI() = %v, want %v", got, want)
	}
}
