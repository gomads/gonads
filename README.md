# Gonads | Golang

[![ver](https://img.shields.io/github/tag/alsi-lawr/gonads)](https://github.com/alsi-lawr/gonads/releases)
![Gover](https://img.shields.io/badge/Go-%3E%3D%201.18-%23007d9c)
[![Godoc](https://godoc.org/github.com/alsi-lawr/gonads?status.svg)](https://pkg.go.dev/github.com/alsi-lawr/gonads)
[![Goreport](https://goreportcard.com/badge/github.com/alsi-lawr/gonads)](https://goreportcard.com/report/github.com/alsi-lawr/gonads)
[![codecov](https://codecov.io/gh/alsi-lawr/gonads/graph/badge.svg?token=FyxqW2TQEY)](https://codecov.io/gh/alsi-lawr/gonads)
[![License](https://img.shields.io/github/license/alsi-lawr/gonads)](./LICENSE)

**`gonads`** brings a set of minimal monads to go. It contains only a set of functional monads that enables FP-style function composition.

## Ethos

- To provide a minimalist library of monads to make certain parts of the code easier to compose.
- To create a strongly typed return type for `error` to enforce error handling.
- To provide a stongly typed union that doesn't require `nil` checks.
- `nil` checks kind of suck, so fewer `nil` checks by wrapping your results in `Option` or `Result` types provides a strongly typed solution to the `nil` check problem.
- Most of all, to have fun (and have a funny name)

## 💡 Features of `gonads`

Currently supported monads:

- **`Option[T]`**, also called a **`Maybe`**: provides a concise and safe way to wrap optional values, enforcing `nil` checks as a drop-in replacement for `nil`-able types.
- **`Either[L, R]`**: provides a concise and safe way to create unions, allowing for enforced union type checking through `Left` and `Right` conditional evaluation.
- **`Result[T]`**: provides the ability to create a strongly typed return type for `error` to enforce error handling.

## 🚀 Getting Started

```sh
go get github.com/alsi-lawr/gonads
```

This library has no other dependencies beyond the go standard library.

## Quick start

Try out an `Option`:

First, import the `option` package:

```go
import (
    "github.com/alsi-lawr/gonads/option"
)
```

Then, try it out:

```go
my_option := option.Some("Hello")

bound := my_option.Bind(
 func(s string) option.Option[string] {
  return option.Some(s + " World!")
 },
)

bound.Match(
 func(s string) {
  fmt.Println(s)
 },
 func() {
  fmt.Println("None")
 },
)
```

You've just used your first `gonad` 🥳!

## Documentation and examples

Detailed documentation is available on [pkg.go.dev](https://pkg.go.dev/github.com/alsi-lawr/gonads), or in your IDE as you play around with your `gonads`.

## 🤝 Contributing

If you *really* want to contribute to this repository, just fork it and start working.
I'd prefer that you start off with an issue of the proposed addition to see if it fits into the gonads ethos,
but if you want to contribute to the gonads library, feel free to do so.

## License

Do whatever you want with this code!
