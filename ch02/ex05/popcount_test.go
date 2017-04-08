package popcount

import (
	"testing"
)

func TestTablePopCount(t *testing.T) {
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
		if got := TablePopCount(test.x); got != test.want {
			t.Errorf("TablePopCount(%d) = %d, want %d", test.x, got, test.want)
		}
	}
}

func TestLSBPopCount(t *testing.T) {
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
		if got := LSBPopCount(test.x); got != test.want {
			t.Errorf("LSBPopCount(%d) = %d, want %d", test.x, got, test.want)
		}
	}
}

func BenchmarkTablePopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TablePopCount(0)
		TablePopCount(1)
		TablePopCount(1 << 8)
		TablePopCount(1<<8 + 1)
		TablePopCount(1<<8 - 1)
		TablePopCount(1<<64 - 1)
	}
}

func BenchmarkLSBPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LSBPopCount(0)
		LSBPopCount(1)
		LSBPopCount(1 << 8)
		LSBPopCount(1<<8 + 1)
		LSBPopCount(1<<8 - 1)
		LSBPopCount(1<<64 - 1)
	}
}
