package mathu_test

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/clavoie/mathu"
)

func TestFloat(t *testing.T) {
	strF := "1.23"
	f64, err := strconv.ParseFloat(strF, 64)

	if err != nil {
		t.Fatal(err)
	}

	eq := func(f mathu.Float) bool { return f64 == float64(f) }

	t.Run("ParseFloat", func(t *testing.T) {
		_, err := mathu.ParseFloat("X")

		if err == nil {
			t.Fatal("Was expecting err")
		}

		f, err := mathu.ParseFloat(strF)

		if err != nil {
			t.Fatal(err)
		}

		if eq(f) == false {
			t.Fatal(f, f64)
		}
	})

	t.Run("Abs", func(t *testing.T) {
		f := -mathu.Float(f64)
		f = f.Abs()

		if eq(f) == false {
			t.Fatal(f, f64)
		}
	})

	t.Run("Ceil", func(t *testing.T) {
		c := mathu.Float(f64).Ceil()
		c64 := int(math.Ceil(f64))

		if c != c64 {
			t.Fatal(c, c64)
		}
	})

	t.Run("Equal", func(t *testing.T) {
		f := mathu.Float(f64)
		thresh := mathu.Float(0.001)
		f2 := f + thresh

		if f.Equal(f2, thresh) == false {
			t.Fatal(f, f2, thresh, f2-f)
		}

		f2 = f2 + thresh

		if f.Equal(f2, thresh) {
			t.Fatal(f, f2, thresh)
		}
	})

	t.Run("Floorf", func(t *testing.T) {
		f := mathu.Float(f64).Floorf()
		ff := math.Floor(f64)

		if float64(f) != ff {
			t.Fatal(f, ff)
		}
	})

	t.Run("Floori", func(t *testing.T) {
		fi := mathu.Float(f64).Floori()
		ff := int(math.Floor(f64))

		if fi != ff {
			t.Fatal(fi, ff)
		}
	})

	t.Run("NaN", func(t *testing.T) {
		f := mathu.Float(f64)

		if f.IsNaN() {
			t.Fatal(f)
		}

		fnan := mathu.Float(math.NaN())
		if fnan.IsNaN() == false {
			t.Fatal(fnan)
		}
	})

	t.Run("IsZero", func(t *testing.T) {
		f := mathu.Float(f64)

		if f.IsZero() {
			t.Fatal(f)
		}

		f = mathu.Float(mathu.Epsilon64)
		if f.IsZero() == false {
			t.Fatal(f)
		}

		f = mathu.Float(0)
		if f.IsZero() == false {
			t.Fatal(f)
		}
	})

	t.Run("GteZero", func(t *testing.T) {
		f := mathu.Float(f64)

		if f.GteZero() == false {
			t.Fatal(f)
		}

		f = -f
		if f.GteZero() {
			t.Fatal(f)
		}

		f = mathu.Float(0)
		if f.GteZero() == false {
			t.Fatal(f)
		}
	})

	t.Run("LteZero", func(t *testing.T) {
		f := mathu.Float(f64)

		if f.LteZero() {
			t.Fatal(f)
		}

		f = -f
		if f.LteZero() == false {
			t.Fatal(f)
		}

		f = mathu.Float(0)
		if f.LteZero() == false {
			t.Fatal(f)
		}
	})

	t.Run("Max", func(t *testing.T) {
		f := mathu.Float(f64)
		f2 := f + 1

		if f.Max(f2) != f2 {
			t.Fatal(f, f2)
		}

		if f2.Max(f) != f2 {
			t.Fatal(f, f2)
		}
	})

	t.Run("Min", func(t *testing.T) {
		f := mathu.Float(f64)
		f2 := f + 1

		if f.Min(f2) != f {
			t.Fatal(f, f2)
		}

		if f2.Min(f) != f {
			t.Fatal(f, f2)
		}
	})

	t.Run("RandomFloat", func(t *testing.T) {
		f1 := mathu.RandomFloat()
		f2 := mathu.RandomFloat()
		one := mathu.Float(1.0)
		zero := mathu.Float(0)

		if f1 == f2 {
			t.Fatal(f1, f2)
		}

		if f1 == one || f2 == one {
			t.Fatal(f1, f2)
		}

		if f1 < zero || f2 < zero {
			t.Fatal(f1, f2)
		}
	})

	t.Run("RoundToInt", func(t *testing.T) {
		f := mathu.Float(1.1)

		if f.RoundToInt() != 1 {
			t.Fatal(f.RoundToInt())
		}

		f = mathu.Float(1.5)
		if f.RoundToInt() != 2 {
			t.Fatal(f.RoundToInt())
		}

		f = mathu.Float(-1.0)
		if f.RoundToInt() != -1 {
			t.Fatal(f.RoundToInt())
		}
	})

	t.Run("Sqrt", func(t *testing.T) {
		f := mathu.Float(f64)
		fq := f.Sqrt()
		ff := math.Sqrt(f64)

		if float64(fq) != ff {
			t.Fatal(fq, ff)
		}
	})

	t.Run("String", func(t *testing.T) {
		f := mathu.Float(f64)
		fs := f.String()
		ff := fmt.Sprintf("%v", f64)

		if fs != ff {
			t.Fatal(fs, ff)
		}
	})
}
