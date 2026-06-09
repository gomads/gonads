/*
Package iters provides an implementation of iterables,
which are types that can replace direct iteration with functional chains.

Iter[T] supports concrete generic methods such as Map, FlatMap, Fold,
and GroupBy. Grouping[K, T] supports direct generic aggregation methods.
These APIs require Go 1.27 or gotip with GOEXPERIMENT=genericmethods.
*/
package iters
