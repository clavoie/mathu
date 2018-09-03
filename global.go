package mathu

const (
	// Epsilon64 is the smallest positive float64 value that is significant in numeric
	// operations or comparisons
	Epsilon64 = float64(7.)/3 - float64(4.)/3 - float64(1.)

	// NegEpsilon64 is the smallest positive float64 value that is significant in
	// numeric operations or comparisons
	NegEpsilon64 = -Epsilon64
)

// Zero tests that a float64 value n is between NegEpsilon64 and Epsilon64, and
// close enough to zero to be considered such for the purposes of comparison
func Zero(n float64) bool {
	return NegEpsilon64 <= n && n <= Epsilon64
}
