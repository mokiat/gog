# GoG (Go Generics)

[![Go Reference](https://pkg.go.dev/badge/github.com/mokiat/gog.svg)](https://pkg.go.dev/github.com/mokiat/gog)
[![Go Report Card](https://goreportcard.com/badge/github.com/mokiat/gog)](https://goreportcard.com/report/github.com/mokiat/gog)

GoG is a Go library with useful generic functions and types.

Since the introduction of generics in Go 1.18, a number of useful and reusable
data structures, algorithms and utility functions are now possible. This library
attempts to cover some of the most common use cases.

It avoids duplicating functions that are already provided by
[slices](https://pkg.go.dev/slices) and
[maps](https://pkg.go.dev/maps).

For a complete list on available functions and types, check the
godoc documentation for this project:

- [gog](https://pkg.go.dev/github.com/mokiat/gog) - general utility functions
- [gog/ds](https://pkg.go.dev/github.com/mokiat/gog/ds) - data structures
- [gog/filter](https://pkg.go.dev/github.com/mokiat/gog) - data filtering
- [gog/opt](https://pkg.go.dev/github.com/mokiat/gog) - optional fields and types


## Examples

**Converting slice values from one type to another:**

```go
source := []int{1, 2, 3, 4}
target := gog.Map(source, func(item int) float64 {
  return float64(item) / 2.0
})
// target = []float64{0.5, 1.0, 1.5, 2.0}
```

**Removing duplicate items in a slice:**

```go
source := []string{"john", "bill", "eric", "john", "max", "eric"}
target := gog.Dedupe(source)
// target = []string{"john", "bill", "eric", "max"}
```

**Finding an item in a slice:**

```go
source := []string{"user 01", "user 02", "user 05", "user 33"}
foundItem, ok := gog.FindFunc(source, func(item string) bool {
  return strings.Contains(item, "02")
})
// ok = true
// foundItem = "user 02"
```

**Selecting specific slice items:**

```go
source := []int{1, 2, 3, 4}
target := gog.Select(source, func(item int) bool {
  return item % 2 == 0
})
// target = []int{2, 4}
```
