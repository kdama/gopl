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

func TestBitShiftPopCount(t *testing.T) {
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
		if got := BitShiftPopCount(test.x); got != test.want {
			t.Errorf("BitShiftPopCount(%d) = %d, want %d", test.x, got, test.want)
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

func BenchmarkBitShiftPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitShiftPopCount(0)
		BitShiftPopCount(1)
		BitShiftPopCount(1 << 8)
		BitShiftPopCount(1<<8 + 1)
		BitShiftPopCount(1<<8 - 1)
		BitShiftPopCount(1<<64 - 1)
	}
}
