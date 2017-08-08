package intsetbench

import (
	"math/rand"
	"time"
)

const (
	max  = 1 << 24
	size = 1 << 23
)

var set1, set2 []int

func init() {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < size; i++ {
		set1 = append(set1, rng.Intn(max+1))
		set2 = append(set2, rng.Intn(max+1))
	}
}
