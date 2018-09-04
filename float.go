package mathu

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

// Float represents a wrapper around float64 that provides some
// convenience functions arounf the math package, and some fuzzy
// tests for zero.
type Float float64

// ParseFloat parses a string into a Float, just like
// strconv.ParseFloat(s, 64).
func ParseFloat(s string) (Float, error) {
	v, err := strconv.ParseFloat(s, 64)
	return Float(v), err
}

// Abs returns the absolute value of this Float.
func (f Float) Abs() Float {
	return Float(math.Abs(float64(f)))
}

// Ceil returns the integer ceiling value of the Float.
func (f Float) Ceil() int {
	return int(math.Ceil(float64(f)))
}

// Equal tests if two Float values are equal within some threshold of
// each other.
func (f1 Float) Equal(f2, threshold Float) bool {
	f3 := (f1 - f2).Abs()
	return f3.IsZero() || f3 <= threshold
}

// Floorf returns the Float representation of the math.Floor
// value of the Float.
func (f Float) Floorf() Float {
	return Float(math.Floor(float64(f)))
}

// Floori returns the int representation of the math.Floor
// value of the Float.
func (f Float) Floori() int {
	return int(math.Floor(float64(f)))
}

// IsNaN tests if the Float is NaN.
func (f Float) IsNaN() bool {
	return math.IsNaN(float64(f))
}

// IsZero tests that a Float value n is between NegEpsilon64 and Epsilon64,
// and close enough to zero to be considered such for the purposes of comparison
func (f Float) IsZero() bool {
	return Zero(float64(f))
}

// GteZero tests that the Float value is greater than or equal to
// zero.
func (f Float) GteZero() bool {
	return f > 0 || f.IsZero()
}

// LteZero tests that the Float is less than or equal to zero.
func (f Float) LteZero() bool {
	return f < 0 || f.IsZero()
}

// Max returns the max Float value between f1 and f2.
func (f1 Float) Max(f2 Float) Float {
	return Float(math.Max(float64(f1), float64(f2)))
}

// Min returns the min Float value between f1 and f2.
func (f1 Float) Min(f2 Float) Float {
	return Float(math.Min(float64(f1), float64(f2)))
}

// RandomFloat returns a Float in the range [0.0, 1.0)
func RandomFloat() Float {
	return Float(rand.Float64())
}

// RoundToInt rounds the Float value up or down to the nearest integer value.
func (f Float) RoundToInt() int {
	return int(math.Floor(float64(f + 0.5)))
}

// Sqrt returns the square root of the Float.
func (f Float) Sqrt() Float {
	return Float(math.Sqrt(float64(f)))
}

// String returns a string representation of the Float.
func (f Float) String() string {
	return fmt.Sprintf("%v", float64(f))
}
