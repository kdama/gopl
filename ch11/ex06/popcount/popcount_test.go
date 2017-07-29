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

func benchmarkTablePopCount(b *testing.B, n uint64) {
	for i := 0; i < b.N; i++ {
		for x := uint64(0); x < 1<<n-1; x++ {
			TablePopCount(x)
		}
	}
}

func BenchmarkTablePopCount0(b *testing.B) {
	benchmarkTablePopCount(b, 0)
}

func BenchmarkTablePopCount1(b *testing.B) {
	benchmarkTablePopCount(b, 1)
}

func BenchmarkTablePopCount2(b *testing.B) {
	benchmarkTablePopCount(b, 2)
}

func BenchmarkTablePopCount4(b *testing.B) {
	benchmarkTablePopCount(b, 4)
}

func BenchmarkTablePopCount8(b *testing.B) {
	benchmarkTablePopCount(b, 8)
}

func BenchmarkTablePopCount16(b *testing.B) {
	benchmarkTablePopCount(b, 16)
}

func benchmarkBitShiftPopCount(b *testing.B, n uint64) {
	for i := 0; i < b.N; i++ {
		for x := uint64(0); x < 1<<n-1; x++ {
			BitShiftPopCount(x)
		}
	}
}

func BenchmarkBitShiftPopCount0(b *testing.B) {
	benchmarkBitShiftPopCount(b, 0)
}

func BenchmarkBitShiftPopCount1(b *testing.B) {
	benchmarkBitShiftPopCount(b, 1)
}

func BenchmarkBitShiftPopCount2(b *testing.B) {
	benchmarkBitShiftPopCount(b, 2)
}

func BenchmarkBitShiftPopCount4(b *testing.B) {
	benchmarkBitShiftPopCount(b, 4)
}

func BenchmarkBitShiftPopCount8(b *testing.B) {
	benchmarkBitShiftPopCount(b, 8)
}

func BenchmarkBitShiftPopCount16(b *testing.B) {
	benchmarkBitShiftPopCount(b, 16)
}

func benchmarkLSBPopCount(b *testing.B, n uint64) {
	for i := 0; i < b.N; i++ {
		for x := uint64(0); x < 1<<n-1; x++ {
			LSBPopCount(x)
		}
	}
}

func BenchmarkLSBPopCount0(b *testing.B) {
	benchmarkLSBPopCount(b, 0)
}

func BenchmarkLSBPopCount1(b *testing.B) {
	benchmarkLSBPopCount(b, 1)
}

func BenchmarkLSBPopCount2(b *testing.B) {
	benchmarkLSBPopCount(b, 2)
}

func BenchmarkLSBPopCount4(b *testing.B) {
	benchmarkLSBPopCount(b, 4)
}

func BenchmarkLSBPopCount8(b *testing.B) {
	benchmarkLSBPopCount(b, 8)
}

func BenchmarkLSBPopCount16(b *testing.B) {
	benchmarkLSBPopCount(b, 16)
}
