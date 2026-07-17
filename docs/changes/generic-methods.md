# Generic Method Migration

Issue: https://github.com/gomads/gonads/issues/1
Reference: https://github.com/golang/go/issues/77273

## Scope

Migrate iterator APIs that currently use helper types or unsafe `any` methods to simulate method-level type parameters. This work depends on Go 1.27, or `gotip` with `GOEXPERIMENT=genericmethods`, and applies only to concrete methods. Generic interface methods are out of scope.

Treat this as a breaking, greenfield cleanup. Do not keep compatibility wrappers for helper APIs that only exist because generic methods were unavailable.

## Current Helper APIs

- `iters.Mappable[T, R]` carries a method result type for chainable `Iter[T]` operations.
- `iters.Aggregable[K, T, R]` carries an aggregate result type for chainable `Grouping[K, T]` operations.
- `iters.LiftMap[T, R]` and `iters.LiftAggregable[K, T, R]` construct those helper types.
- `Mappable[T, R].ToIter` and `Aggregable[K, T, R].ToGrouping` only support the helper-type conversion path.

## Direct Method Targets

Add concrete generic methods once the toolchain supports them:

- `Iter[T].Map[R any](func(T) R) Iter[R]`
- `Iter[T].MapI[R any](func(int, T) R) Iter[R]`
- `Iter[T].MapErr[R any](func(T) (R, error)) result.Result[Iter[R]]`
- `Iter[T].FlatMap[R any](func(T) Iter[R]) Iter[R]`
- `Iter[T].FlatMapI[R any](func(int, T) Iter[R]) Iter[R]`
- `Iter[T].Fold[A any](A, func(A, T) A) A`
- `Iter[T].FoldI[A any](A, func(int, A, T) A) A`
- `Iter[T].GroupBy[K comparable](func(T) K) Grouping[K, T]`
- `Iter[T].GroupByI[K comparable](func(int, T) K) Grouping[K, T]`
- `Grouping[K, T].Aggregate[R any](func(Iter[T]) R) map[K]R`
- `Grouping[K, T].AggregateI[R any](func(K, Iter[T]) R) map[K]R`

Keep package-level functions such as `Map`, `FlatMap`, `Fold`, `GroupBy`, and `Aggregate` unless a separate cleanup decides to remove them.

## Removals

Remove APIs that only exist to carry method-level type parameters:

- `Mappable[T, R]`
- `Aggregable[K, T, R]`
- `LiftMap[T, R]`
- `LiftAggregable[K, T, R]`
- `Mappable[T, R].ToIter`
- `Aggregable[K, T, R].ToGrouping`

Remove unsafe method variants that become redundant after typed generic methods exist:

- `Iter[T].MapUnsafe`
- `Iter[T].MapIUnsafe`
- `Iter[T].MapErrUnsafe`
- `Iter[T].FlatMapUnsafe`
- `Iter[T].FlatMapIUnsafe`
- `Iter[T].FoldUnsafe`
- `Iter[T].FoldIUnsafe`
- `Iter[T].GroupByUnsafe`
- `Iter[T].GroupByIUnsafe`
- `Grouping[K, T].AggregateUnsafe`
- `Grouping[K, T].AggregateIUnsafe`

## File-Level Work

- `iters/types.go`: remove helper type definitions and lift/conversion functions; keep `Iter[T]`, `Grouping[K, T]`, `LiftSlice`, and `Iter[T].ToSlice`.
- `iters/map.go`: move chainable map methods from `Mappable[T, R]` to `Iter[T]` with method type parameters; remove unsafe variants.
- `iters/flatmap.go`: move chainable flat-map methods from `Mappable[T, R]` to `Iter[T]`; remove unsafe variants.
- `iters/fold.go`: move chainable fold methods from `Mappable[T, A]` to `Iter[T]`; remove unsafe variants.
- `iters/group.go`: add typed `Iter[T].GroupBy` and `Iter[T].GroupByI`; remove unsafe variants.
- `iters/aggregate.go`: move chainable aggregate methods from `Aggregable[K, T, R]` to `Grouping[K, T]`; remove unsafe variants.
- `README.md` and `iters/doc.go`: replace helper-type language and examples with direct method usage.

## Test Updates

- Remove tests for `LiftMap`, `LiftAggregable`, `Mappable`, `Aggregable`, and unsafe method variants.
- Rewrite method tests to start from `iters.Iter[T]` or `iters.LiftSlice`.
- Add typed method coverage for result type changes, including `Iter[int].Map[string]`, `Iter[int].FlatMap[string]`, `Iter[int].Fold[string]`, `Iter[int].GroupBy[string]`, and `Grouping[string, int].Aggregate[int]`.
- Keep package-level function tests intact unless behavior changes.

## Validation

- Use a Go toolchain that supports generic concrete methods: Go 1.27 or `gotip` with `GOEXPERIMENT=genericmethods`.
- Run `go test ./...` under that toolchain.
- Confirm no remaining references to `Mappable`, `Aggregable`, `LiftMap`, `LiftAggregable`, or removed unsafe method names remain in code, tests, or docs.
