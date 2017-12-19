package set_test

import (
	"testing"

	"github.com/Warashi/set"
	"github.com/k0kubun/pp"
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
				pp.Println(s)
			}
		}
	}
}
