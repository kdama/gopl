package intsetbench

import (
	"testing"

	"github.com/kdama/gopl/ch11/ex02/intset32"
	"github.com/kdama/gopl/ch11/ex02/intset64"
	"github.com/kdama/gopl/ch11/ex02/mapset"
)

var MAX = 1 << 16

// IntSet32

func BenchmarkIntSet32Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set := intset32.NewIntSet()
		for n := 0; n < MAX; n++ {
			set.Add(n / 2)
		}
	}
}

func BenchmarkIntSet32AddAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set := intset32.NewIntSet()
		for n := 0; n < MAX; n++ {
			set.AddAll(n*2, n*2+1)
		}
	}
}

func BenchmarkIntSet32Clear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set := intset32.NewIntSet()
		set.Clear()
	}
}

func BenchmarkIntSet32Copy(b *testing.B) {
	set := intset32.NewIntSet()
	for n := 0; n < MAX; n++ {
		set.Add(n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Copy()
	}
}

func BenchmarkIntSet32DifferenceWith(b *testing.B) {
	set1 := intset32.NewIntSet()
	set2 := intset32.NewIntSet()
	for n := 0; n < MAX; n++ {
		set1.Add(n)
		set2.Add(n + MAX/2)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set1.DifferenceWith(set2)
	}
}

func BenchmarkIntSet32Elems(b *testing.B) {
	set := intset32.NewIntSet()
	for n := 0; n < MAX; n++ {
		set.Add(n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Elems()
	}
}

func BenchmarkIntSet32Has(b *testing.B) {
	set := intset32.NewIntSet()
	for n := 0; n < MAX/2; n++ {
		set.Add(n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for n := 0; n < MAX; n++ {
			set.Has(n)
		}
	}
}

func BenchmarkIntSet32IntersectWith(b *testing.B) {
	set1 := intset32.NewIntSet()
	set2 := intset32.NewIntSet()
	for n := 0; n < MAX; n++ {
		set1.Add(n)
		set2.Add(n + MAX/2)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set1.IntersectWith(set2)
	}
}

func BenchmarkIntSet32Len(b *testing.B) {
	set := intset32.NewIntSet()
	for n := 0; n < MAX/2; n++ {
		set.Add(n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Len()
	}
}

func BenchmarkIntSet32Remove(b *testing.B) {
	set := intset32.NewIntSet()
	for n := 0; n < MAX; n = n + 2 {
		set.Add(n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for n := 0; n < MAX; n++ {
			set.Remove(n)
		}
	}
}

func BenchmarkIntSet32SymmetricDifference(b *testing.B) {
	set1 := intset32.NewIntSet()
	set2 := intset32.NewIntSet()
	for n := 0; n < MAX; n++ {
		set1.Add(n)
		set2.Add(n + MAX/2)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set1.SymmetricDifference(set2)
	}
}

func BenchmarkIntSet32UnionWith(b *testing.B) {
	set1 := intset32.NewIntSet()
	set2 := intset32.NewIntSet()
	for n := 0; n < MAX; n++ {
		set1.Add(n)
		set2.Add(n + MAX/2)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set1.UnionWith(set2)
	}
}

// IntSet64

func BenchmarkIntSet64Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set := intset64.NewIntSet()
		for n := 0; n < MAX; n++ {
			set.Add(n / 2)
		}
	}
}

func BenchmarkIntSet64AddAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set := intset64.NewIntSet()
		for n := 0; n < MAX; n++ {
			set.AddAll(n*2, n*2+1)
		}
	}
}

func BenchmarkIntSet64Clear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set := intset64.NewIntSet()
		set.Clear()
	}
}

func BenchmarkIntSet64Copy(b *testing.B) {
	set := intset64.NewIntSet()
	for n := 0; n < MAX; n++ {
		set.Add(n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Copy()
	}
}

func BenchmarkIntSet64DifferenceWith(b *testing.B) {
	set1 := intset64.NewIntSet()
	set2 := intset64.NewIntSet()
	for n := 0; n < MAX; n++ {
		set1.Add(n)
		set2.Add(n + MAX/2)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set1.DifferenceWith(set2)
	}
}

func BenchmarkIntSet64Elems(b *testing.B) {
	set := intset64.NewIntSet()
	for n := 0; n < MAX; n++ {
		set.Add(n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Elems()
	}
}

func BenchmarkIntSet64Has(b *testing.B) {
	set := intset64.NewIntSet()
	for n := 0; n < MAX/2; n++ {
		set.Add(n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for n := 0; n < MAX; n++ {
			set.Has(n)
		}
	}
}

func BenchmarkIntSet64IntersectWith(b *testing.B) {
	set1 := intset64.NewIntSet()
	set2 := intset64.NewIntSet()
	for n := 0; n < MAX; n++ {
		set1.Add(n)
		set2.Add(n + MAX/2)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set1.IntersectWith(set2)
	}
}

func BenchmarkIntSet64Len(b *testing.B) {
	set := intset64.NewIntSet()
	for n := 0; n < MAX/2; n++ {
		set.Add(n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Len()
	}
}

func BenchmarkIntSet64Remove(b *testing.B) {
	set := intset64.NewIntSet()
	for n := 0; n < MAX; n = n + 2 {
		set.Add(n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for n := 0; n < MAX; n++ {
			set.Remove(n)
		}
	}
}

func BenchmarkIntSet64SymmetricDifference(b *testing.B) {
	set1 := intset64.NewIntSet()
	set2 := intset64.NewIntSet()
	for n := 0; n < MAX; n++ {
		set1.Add(n)
		set2.Add(n + MAX/2)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set1.SymmetricDifference(set2)
	}
}

func BenchmarkIntSet64UnionWith(b *testing.B) {
	set1 := intset64.NewIntSet()
	set2 := intset64.NewIntSet()
	for n := 0; n < MAX; n++ {
		set1.Add(n)
		set2.Add(n + MAX/2)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set1.UnionWith(set2)
	}
}

// MapSet

func BenchmarkMapSetAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set := mapset.NewIntSet()
		for n := 0; n < MAX; n++ {
			set.Add(n / 2)
		}
	}
}

func BenchmarkMapSetAddAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set := mapset.NewIntSet()
		for n := 0; n < MAX; n++ {
			set.AddAll(n*2, n*2+1)
		}
	}
}

func BenchmarkMapSetClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set := mapset.NewIntSet()
		set.Clear()
	}
}

func BenchmarkMapSetCopy(b *testing.B) {
	set := mapset.NewIntSet()
	for n := 0; n < MAX; n++ {
		set.Add(n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Copy()
	}
}

func BenchmarkMapSetDifferenceWith(b *testing.B) {
	set1 := mapset.NewIntSet()
	set2 := mapset.NewIntSet()
	for n := 0; n < MAX; n++ {
		set1.Add(n)
		set2.Add(n + MAX/2)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set1.DifferenceWith(set2)
	}
}

func BenchmarkMapSetElems(b *testing.B) {
	set := mapset.NewIntSet()
	for n := 0; n < MAX; n++ {
		set.Add(n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Elems()
	}
}

func BenchmarkMapSetHas(b *testing.B) {
	set := mapset.NewIntSet()
	for n := 0; n < MAX/2; n++ {
		set.Add(n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for n := 0; n < MAX; n++ {
			set.Has(n)
		}
	}
}

func BenchmarkMapSetIntersectWith(b *testing.B) {
	set1 := mapset.NewIntSet()
	set2 := mapset.NewIntSet()
	for n := 0; n < MAX; n++ {
		set1.Add(n)
		set2.Add(n + MAX/2)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set1.IntersectWith(set2)
	}
}

func BenchmarkMapSetLen(b *testing.B) {
	set := mapset.NewIntSet()
	for n := 0; n < MAX/2; n++ {
		set.Add(n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Len()
	}
}

func BenchmarkMapSetRemove(b *testing.B) {
	set := mapset.NewIntSet()
	for n := 0; n < MAX; n = n + 2 {
		set.Add(n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for n := 0; n < MAX; n++ {
			set.Remove(n)
		}
	}
}

func BenchmarkMapSetSymmetricDifference(b *testing.B) {
	set1 := mapset.NewIntSet()
	set2 := mapset.NewIntSet()
	for n := 0; n < MAX; n++ {
		set1.Add(n)
		set2.Add(n + MAX/2)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set1.SymmetricDifference(set2)
	}
}

func BenchmarkMapSetUnionWith(b *testing.B) {
	set1 := mapset.NewIntSet()
	set2 := mapset.NewIntSet()
	for n := 0; n < MAX; n++ {
		set1.Add(n)
		set2.Add(n + MAX/2)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set1.UnionWith(set2)
	}
}
