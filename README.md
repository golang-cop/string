<p align="center"><img src="https://raw.githubusercontent.com/go-composites/brand/main/social/go-composites.png" alt="go-composites/string" width="720"></p>

# string

[![ci](https://github.com/go-composites/string/actions/workflows/ci.yml/badge.svg)](https://github.com/go-composites/string/actions/workflows/ci.yml)

The text composite of [go-composites](https://github.com/go-composites). A
`String` wraps a Go `string` and exposes its operations either as plain Go
predicates (`bool`/`int`) or, when they produce a new value, as a
[`Result`](https://github.com/go-composites/result) — so failures and chaining
are *values*, never panics or `nil`.

## Install

```sh
go get github.com/go-composites/string
```

## Usage

```golang
package main

import (
    "fmt"

    String "github.com/go-composites/string/src"
    Result "github.com/go-composites/result/src"
    Array  "github.com/go-composites/array/src"
)

func main() {
    s := String.New(String.WithGoString(`Hello World!`))
    fmt.Println(s.ToGoString())        // Hello World!
    fmt.Println(s.Length())            // 12 (rune count)
    fmt.Println(s.Upper().Payload().(String.Interface).ToGoString()) // HELLO WORLD!

    if r := s.Split(` `); !r.HasError() {
        words := r.Payload().(Array.Interface)
        first := words.First().Payload().(String.Interface)
        fmt.Println(first.ToGoString())  // Hello
    }
}
```

## API

### Construction & conversion

- `New(opts ...Option) Interface` — empty `String` by default.
- `WithGoString(string) Option` — set the initial value.
- `Set(string) Result` — set the value; payload is the `String`.
- `ToGoString() string` — the underlying Go string.

### Inspection (return plain Go values)

| method | returns | notes |
| --- | --- | --- |
| `Length()` | Go `int` | rune count (multibyte-safe, not byte length) |
| `Contains(substr)` | Go `bool` | substring test |
| `Equal(other)` | Go `bool` | same underlying string |
| `StartsWith(prefix)` | Go `bool` | |
| `EndsWith(suffix)` | Go `bool` | |
| `IsNull()` | Go `bool` | `true` only for the Null-Object |

### Transformation (return a `Result` wrapping a new `String`)

| method | result payload | notes |
| --- | --- | --- |
| `Concat(other)` | a new `String` | receiver followed by `other` |
| `Replace(old, new)` | a new `String` | replace all occurrences |
| `Upper()` | a new `String` | upper case |
| `Lower()` | a new `String` | lower case |
| `Trim()` | a new `String` | strip leading/trailing white space |
| `Format(args...)` | a new `String` | `fmt.Sprintf` with the value as the format |
| `Split(sep)` | an [`Array`](https://github.com/go-composites/array) of `String`s | split on `sep` |

## Null-Object

`Null()` returns the never-nil Null-Object `String`: `ToGoString()` is `""`,
predicates report `false`, and value operations return a `Result` wrapping the
null `String`. `IsNull()` reports `true` for it. A sibling package
`github.com/go-composites/string/src/null` (`NullString`) provides an equivalent
Null `String` over the shared `go-composites/null` core, with its own `New()`.

## License

BSD-3-Clause © the go-composites/string authors.
