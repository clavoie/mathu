package mathu_test

import (
	"testing"

	"github.com/clavoie/mathu"
)

func TestEpsilon(t *testing.T) {
	e := mathu.NewEpsilon()

	if e.E64() != mathu.Epsilon64 {
		t.Fatal(e.E64(), mathu.Epsilon64)
	}

	if e.NegE64() != mathu.NegEpsilon64 {
		t.Fatal(e.NegE64(), mathu.NegEpsilon64)
	}

	if e.IsZero(1) || e.IsZero(-1) {
		t.Fatal("1 or -1 should not be zero")
	}

	if e.IsZero(mathu.Epsilon64) == false || e.IsZero(mathu.NegEpsilon64) == false {
		t.Fatal("Epsilon or neg epsilon should be zero")
	}
}
