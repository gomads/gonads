package iters

// Aggregate applies an aggregation function to each group and returns a map of the aggregated results.
//
// Type signature:
//
//	Aggregate :: Grouping K [T] -> ([T] -> R) -> Map K R
func Aggregate[K comparable, T, R any](g Grouping[K, T], agg func(items Collection[T]) R) map[K]R {
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
func AggregateI[K comparable, T, R any](g Grouping[K, T], agg func(key K, items Collection[T]) R) map[K]R {
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
//	Aggregate :: Grouping K [T] -> ((T] -> any) -> Map K any
func (gq Aggregable[K, T, R]) Aggregate(agg func(items Collection[T]) R) map[K]R {
	return Aggregate(gq.ToGrouping(), agg)
}

// AggregateI applies an aggregation function to each group and returns a map of the aggregated results.
//
// Type signature:
//
//	AggregateI :: Grouping K [T] -> ((K, [T]) -> any) -> Map K any
func (gq Aggregable[K, T, R]) AggregateI(agg func(key K, items Collection[T]) R) map[K]R {
	return AggregateI(gq.ToGrouping(), agg)
}

// Aggregate applies an aggregation function to each group and returns a map of the aggregated results.
//
// [Unsafe Variant] This method loses compile-time type safety by returning `any` rather than a strongly-typed map.
// Use only when necessary.
//
// Type signature:
//
//	Aggregate :: Grouping K [T] -> ((T] -> any) -> Map K any
func (gq Grouping[K, T]) AggregateUnsafe(agg func(items Collection[T]) any) map[K]any {
	return Aggregate(gq, agg)
}

// AggregateI applies an aggregation function to each group and returns a map of the aggregated results.
//
// [Unsafe Variant] This method loses compile-time type safety by returning `any` rather than a strongly-typed map.
// Use only when necessary.
//
// Type signature:
//
//	AggregateI :: Grouping K [T] -> ((K, [T]) -> any) -> Map K any
func (gq Grouping[K, T]) AggregateIUnsafe(agg func(key K, items Collection[T]) any) map[K]any {
	return AggregateI(gq, agg)
}
