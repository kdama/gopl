package intsetbench

import (
	"testing"

	impl "github.com/kdama/gopl/ch11/ex02/bitset32"
)

func BenchmarkBitSet32Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intset := impl.NewIntSet()
		for _, n := range set1 {
			intset.Add(n)
		}
	}
}

func BenchmarkBitSet32AddAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intset := impl.NewIntSet()
		intset.AddAll(set1...)
	}
}

func BenchmarkBitSet32Clear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intset := impl.NewIntSet()
		intset.Clear()
	}
}

func BenchmarkBitSet32Copy(b *testing.B) {
	intset := impl.NewIntSet()
	for _, n := range set1 {
		intset.Add(n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intset.Copy()
	}
}

func BenchmarkBitSet32DifferenceWith(b *testing.B) {
	intset1 := impl.NewIntSet()
	intset2 := impl.NewIntSet()
	for _, n := range set1 {
		intset1.Add(n)
	}
	for _, n := range set2 {
		intset2.Add(n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intset1.DifferenceWith(intset2)
	}
}

func BenchmarkBitSet32Elems(b *testing.B) {
	intset := impl.NewIntSet()
	for _, n := range set1 {
		intset.Add(n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intset.Elems()
	}
}

func BenchmarkBitSet32Has(b *testing.B) {
	intset := impl.NewIntSet()
	for _, n := range set1 {
		intset.Add(n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, n := range set1 {
			intset.Has(n)
		}
	}
}

func BenchmarkBitSet32IntersectWith(b *testing.B) {
	intset1 := impl.NewIntSet()
	intset2 := impl.NewIntSet()
	for _, n := range set1 {
		intset1.Add(n)
	}
	for _, n := range set2 {
		intset2.Add(n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intset1.IntersectWith(intset2)
	}
}

func BenchmarkBitSet32Len(b *testing.B) {
	intset := impl.NewIntSet()
	for _, n := range set1 {
		intset.Add(n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intset.Len()
	}
}

// 注意: このままでは Add のための時間が余計に含まれるので、BenchmarkBitSet32Add の結果との差分を使う必要があります。
func BenchmarkBitSet32Remove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intset := impl.NewIntSet()
		for _, n := range set1 {
			intset.Add(n)
		}
		for _, n := range set1 {
			intset.Remove(n)
		}
	}
}

func BenchmarkBitSet32SymmetricDifference(b *testing.B) {
	intset1 := impl.NewIntSet()
	intset2 := impl.NewIntSet()
	for _, n := range set1 {
		intset1.Add(n)
	}
	for _, n := range set2 {
		intset2.Add(n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intset1.SymmetricDifference(intset2)
	}
}

func BenchmarkBitSet32UnionWith(b *testing.B) {
	intset1 := impl.NewIntSet()
	intset2 := impl.NewIntSet()
	for _, n := range set1 {
		intset1.Add(n)
	}
	for _, n := range set2 {
		intset2.Add(n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intset1.UnionWith(intset2)
	}
}
