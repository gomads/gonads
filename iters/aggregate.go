package iters

// Aggregate applies an aggregation function to each group and returns a map of the aggregated results.
//
// Type signature:
//
//	Aggregate :: Grouping K [T] -> ([T] -> R) -> Map K R
func Aggregate[K comparable, T, R any](g Grouping[K, T], agg func(items Iter[T]) R) map[K]R {
	result := make(map[K]R, len(g))
	for k, items := range g {
		result[k] = agg(items)
	}
	return result
}

// AggregateI applies an aggregation function to each group and returns a map of the aggregated results.
//
// Type signature:
//
//	AggregateI :: Grouping K [T] -> ((K, [T]) -> R) -> Map K R
func AggregateI[K comparable, T, R any](g Grouping[K, T], agg func(key K, items Iter[T]) R) map[K]R {
	result := make(map[K]R, len(g))
	for k, items := range g {
		result[k] = agg(k, items)
	}
	return result
}

// Aggregate applies an aggregation function to each group and returns a map of the aggregated results.
//
// Type signature:
//
//	Aggregate :: Grouping K [T] -> ((T] -> R) -> Map K R
func (gq Grouping[K, T]) Aggregate[R any](agg func(items Iter[T]) R) map[K]R {
	return Aggregate(gq, agg)
}

// AggregateI applies an aggregation function to each group and returns a map of the aggregated results.
//
// Type signature:
//
//	AggregateI :: Grouping K [T] -> ((K, [T]) -> R) -> Map K R
func (gq Grouping[K, T]) AggregateI[R any](agg func(key K, items Iter[T]) R) map[K]R {
	return AggregateI(gq, agg)
}
