package mathu

// Epsilon is an injectable wrapper around the package level
// epsilon definitions
type Epsilon interface {
	// E64 returns the positive float64 representation of
	// epsilon.
	E64() float64

	// NegE64 returns the negative float64 representation of
	// epsilon.
	NegE64() float64

	// IsZero tests if the given value is between E64() and NegE64()
	IsZero(float64) bool
}

// epsilon is an implementation of Epsilon
type epsilon struct{}

// NewEpsilon returns a new instance of an Epsilon implementation
func NewEpsilon() Epsilon { return new(epsilon) }

func (e *epsilon) E64() float64            { return Epsilon64 }
func (e *epsilon) NegE64() float64         { return NegEpsilon64 }
func (e *epsilon) IsZero(f64 float64) bool { return Zero(f64) }
