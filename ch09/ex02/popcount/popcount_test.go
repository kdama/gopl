package popcount

import (
	"sync"
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

	var wg sync.WaitGroup
	for _, test := range tests {
		test := test
		wg.Add(1)
		go func() {
			defer wg.Done()
			if got := PopCount(test.x); got != test.want {
				t.Errorf("PopCount(%d) = %d, want %d", test.x, got, test.want)
			}
		}()
	}
	wg.Wait()
}
