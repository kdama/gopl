package intsetbench

import (
	"testing"

	impl "github.com/kdama/gopl/ch11/ex02/mapset"
)

func BenchmarkMapSetAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intset := impl.NewIntSet()
		for _, n := range set1 {
			intset.Add(n)
		}
	}
}

func BenchmarkMapSetAddAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intset := impl.NewIntSet()
		intset.AddAll(set1...)
	}
}

func BenchmarkMapSetClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intset := impl.NewIntSet()
		intset.Clear()
	}
}

func BenchmarkMapSetCopy(b *testing.B) {
	intset := impl.NewIntSet()
	for _, n := range set1 {
		intset.Add(n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intset.Copy()
	}
}

func BenchmarkMapSetDifferenceWith(b *testing.B) {
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

func BenchmarkMapSetElems(b *testing.B) {
	intset := impl.NewIntSet()
	for _, n := range set1 {
		intset.Add(n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intset.Elems()
	}
}

func BenchmarkMapSetHas(b *testing.B) {
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

func BenchmarkMapSetIntersectWith(b *testing.B) {
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

func BenchmarkMapSetLen(b *testing.B) {
	intset := impl.NewIntSet()
	for _, n := range set1 {
		intset.Add(n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intset.Len()
	}
}

// 注意: このままでは Add のための時間が余計に含まれるので、BenchmarkMapSetAdd の結果との差分を使う必要があります。
func BenchmarkMapSetRemove(b *testing.B) {
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

func BenchmarkMapSetSymmetricDifference(b *testing.B) {
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

func BenchmarkMapSetUnionWith(b *testing.B) {
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
