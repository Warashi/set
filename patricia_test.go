package set_test

import (
	"testing"

	"github.com/Warashi/set"
)

func TestNewPatricia(t *testing.T) {
	ex := [][]string{{"a"}, {"a", "ab"}, {"a", "ab", "ac", "abc"}, {"a", "b", "c"}, {""}, {"", "a"}}
	for _, e := range ex {
		s := set.NewPatricia(e...)
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
		s := set.NewPatricia(e...)
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
		s := set.NewPatricia(e...)
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
