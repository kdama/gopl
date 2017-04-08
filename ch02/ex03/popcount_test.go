package popcount

import (
	"testing"
)

func TestPopCount(t *testing.T) {
	var tests = []struct {
		x    uint64
		want int
	}{
		{0, 0},
		{1, 1},
		{1 << 8, 1},
		{1<<8 + 1, 2},
		{1<<8 - 1, 8},
		{1<<64 - 1, 64},
	}

	for _, test := range tests {
		if got := PopCount(test.x); got != test.want {
			t.Errorf("PopCount(%d) = %d, want %d", test.x, got, test.want)
		}
	}
}

func TestAlternativePopCount(t *testing.T) {
	var tests = []struct {
		x    uint64
		want int
	}{
		{0, 0},
		{1, 1},
		{1 << 8, 1},
		{1<<8 + 1, 2},
		{1<<8 - 1, 8},
		{1<<64 - 1, 64},
	}

	for _, test := range tests {
		if got := AlternativePopCount(test.x); got != test.want {
			t.Errorf("AlternativePopCount(%d) = %d, want %d", test.x, got, test.want)
		}
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0)
		PopCount(1)
		PopCount(1 << 8)
		PopCount(1<<8 + 1)
		PopCount(1<<8 - 1)
		PopCount(1<<64 - 1)
	}
}

func BenchmarkAlternativePopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AlternativePopCount(0)
		AlternativePopCount(1)
		AlternativePopCount(1 << 8)
		AlternativePopCount(1<<8 + 1)
		AlternativePopCount(1<<8 - 1)
		AlternativePopCount(1<<64 - 1)
	}
}
