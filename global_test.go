package mathu_test

import (
	"math"
	"testing"

	"github.com/clavoie/mathu"
)

func TestZero(t *testing.T) {
	f := 1.0
	fe := f + mathu.Epsilon64

	if fe <= f || fe >= 2.0 {
		t.Fatal(f, mathu.Epsilon64, fe)
	}

	nf := -1.0
	nfe := nf + mathu.NegEpsilon64

	if nfe >= nf || nfe <= -2.0 {
		t.Fatal(nf, mathu.NegEpsilon64, nfe)
	}

	if mathu.Zero(f) {
		t.Fatal("Should not be zero", f)
	}

	if mathu.Zero(mathu.Epsilon64) == false {
		t.Fatal("Epsilon should be zero", mathu.Epsilon64)
	}

	if mathu.Zero(nf) {
		t.Fatal("Should not be zero", nf)
	}

	if mathu.Zero(mathu.NegEpsilon64) == false {
		t.Fatal("Should be zero", mathu.NegEpsilon64)
	}

	if mathu.Zero(math.SmallestNonzeroFloat64) == false {
		t.Fatal("Should be zero", math.SmallestNonzeroFloat64, math.SmallestNonzeroFloat64 <= mathu.Epsilon64, mathu.NegEpsilon64 <= math.SmallestNonzeroFloat64)
	}
}
