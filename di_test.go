package mathu_test

import (
	"testing"

	"github.com/clavoie/mathu"
)

func TestNewDiDefs(t *testing.T) {
	defs := mathu.NewDiDefs()

	if defs == nil {
		t.Fatal("Expecting non-nil defs")
	}

	defs2 := mathu.NewDiDefs()
	if defs[0] == defs2[0] {
		t.Fatal("Not expecting defs to match")
	}
}
