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

func TestLoopPopCount(t *testing.T) {
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
		if got := LoopPopCount(test.x); got != test.want {
			t.Errorf("LoopPopCount(%d) = %d, want %d", test.x, got, test.want)
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

func BenchmarkLoopPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoopPopCount(0)
		LoopPopCount(1)
		LoopPopCount(1 << 8)
		LoopPopCount(1<<8 + 1)
		LoopPopCount(1<<8 - 1)
		LoopPopCount(1<<64 - 1)
	}
}
