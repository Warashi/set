package set_test

import (
	"testing"

	"github.com/Warashi/set"
)

func TestNew(t *testing.T) {
	ex := [][]string{{"a"}, {"a", "ab"}, {"a", "ab", "ac", "abc"}, {"a", "b", "c"}, {""}, {"", "a"}}
	for _, e := range ex {
		s := set.New(e...)
		if len(e) != s.Len() {
			t.Errorf("expect %d, but %d", len(e), s.Len())
		}
		for _, i := range e {
			if !s.Has(i) {
				t.Error("s does not have", i)
			}
		}
	}
}

func TestItems(t *testing.T) {
	ex := [][]string{{"a"}, {"a", "ab"}, {"a", "ab", "ac", "abc"}, {"a", "b", "c"}, {""}, {"", "a"}}
	for _, e := range ex {
		s := set.New(e...)
		var c int
		for i := range s.Items() {
			if !s.Has(i) {
				t.Errorf("s does not have %s, but items generated", i)
			}
			c++
		}
		if c != s.Len() {
			t.Errorf("generated %d items, but set has %d items", c, s.Len())
		}
	}
}

func TestDelete(t *testing.T) {
	ex := [][]string{{"a"}, {"a", "ab"}, {"a", "ab", "ac", "abc"}, {"a", "b", "c"}, {""}, {"", "a"}}
	for _, e := range ex {
		t.Logf("start example %#v", e)
		s := set.New(e...)
		c := s.Len()
		for _, i := range e {
			c--
			t.Logf("delete %s", i)
			s.Delete(i)
			if c != s.Len() {
				t.Errorf("expect %d, but %d in case of %#v", c, s.Len(), e)
			}
		}
	}
}

func TestAnd(t *testing.T) {
	examples := []struct {
		a, b, e []string
	}{
		{a: []string{"a"}, b: []string{"b"}, e: []string{}},
		{a: []string{"a", "c"}, b: []string{"b", "c"}, e: []string{"c"}},
		{a: []string{"a"}, b: []string{"a"}, e: []string{"a"}}}

	dst := set.New()
	for _, ex := range examples {
		set.And(dst, set.New(ex.a...), set.New(ex.b...))
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
