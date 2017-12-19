package set_test

import (
	"testing"

	"github.com/Warashi/set"
)

func TestAnd(t *testing.T) {
	examples := []struct {
		a, b, e []string
	}{
		{a: []string{"a"}, b: []string{"b"}, e: []string{}},
		{a: []string{"a", "c"}, b: []string{"b", "c"}, e: []string{"c"}},
		{a: []string{"a"}, b: []string{"a"}, e: []string{"a"}}}

	dst := set.NewPatricia()
	for _, ex := range examples {
		set.And(dst, set.NewPatricia(ex.a...), set.NewPatricia(ex.b...))
		if len(ex.e) != dst.Len() {
			t.Errorf("expect %d, but %d", len(ex.e), dst.Len())
		}
		for _, i := range ex.e {
			if !dst.Has(i) {
				t.Error("dst does not have", i)
			}
		}

	}
}
