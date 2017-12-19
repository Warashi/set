package set_test

import (
	"testing"

	"github.com/Warashi/gorex"
	"github.com/Warashi/set"
)

func BenchmarkNewPatricia(b *testing.B) {
	g, err := gorex.New("[abc]{1,10}")
	if err != nil {
		panic(err)
	}
	items, err := g.Expand()
	if err != nil {
		panic(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = set.NewPatricia(items...)
	}
}

func BenchmarkPatriciaHas(b *testing.B) {
	g, err := gorex.New("[abc]{1,10}")
	if err != nil {
		panic(err)
	}
	items, err := g.Expand()
	if err != nil {
		panic(err)
	}
	s := set.NewPatricia(items...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, item := range items {
			s.Has(item)
		}
	}
}

func BenchmarkPatriciaItems(b *testing.B) {
	g, err := gorex.New("[abc]{1,10}")
	if err != nil {
		panic(err)
	}
	items, err := g.Expand()
	if err != nil {
		panic(err)
	}
	s := set.NewPatricia(items...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for range s.Items() {
		}
	}
}

func BenchmarkNewMap(b *testing.B) {
	g, err := gorex.New("[abc]{1,10}")
	if err != nil {
		panic(err)
	}
	items, err := g.Expand()
	if err != nil {
		panic(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = set.NewMap(items...)
	}
}

func BenchmarkMapHas(b *testing.B) {
	g, err := gorex.New("[abc]{1,10}")
	if err != nil {
		panic(err)
	}
	items, err := g.Expand()
	if err != nil {
		panic(err)
	}
	s := set.NewMap(items...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, item := range items {
			s.Has(item)
		}
	}
}

func BenchmarkMapItems(b *testing.B) {
	g, err := gorex.New("[abc]{1,10}")
	if err != nil {
		panic(err)
	}
	items, err := g.Expand()
	if err != nil {
		panic(err)
	}
	s := set.NewMap(items...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for range s.Items() {
		}
	}
}
