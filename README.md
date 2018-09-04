# mathu [![GoDoc](https://godoc.org/github.com/clavoie/mathu?status.svg)](http://godoc.org/github.com/clavoie/mathu) [![Build Status](https://travis-ci.org/clavoie/mathu.svg?branch=master)](https://travis-ci.org/clavoie/mathu) [![codecov](https://codecov.io/gh/clavoie/mathu/branch/master/graph/badge.svg)](https://codecov.io/gh/clavoie/mathu) [![Go Report Card](https://goreportcard.com/badge/github.com/clavoie/mathu)](https://goreportcard.com/report/github.com/clavoie/mathu)

[Epsilon definitions](https://godoc.org/github.com/clavoie/mathu#pkg-constants) for Go.

```go

func Equal(f1, f2, threshold float64) bool {
	f3 := math.Abs(f1 - f2)
	return mathu.Zero(f3) || f3 <= threshold
}
```

## Dependency Injection

The epsilon constants as well as all top level package functions are wrapped in an interface if you would prefer to inject them into your code rather than referencing the package directly.

## Floats

A wrapper type is provided around float64 that integrates with the epsilon constants and testing for zero:

```go
f := mathu.Float(1.23)

if f.IsZero() {	/* ... */ }
if f.GteZero() { /* ... */ }
if f.LteZero() { /* ... */ }
if f.Equal(f2, 0.001) { /* ... */ }
```
