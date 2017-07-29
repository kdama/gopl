package mapset

import (
	"reflect"
	"testing"
)

func TestLen(t *testing.T) {
	var tests = []struct {
		in   []int
		want int
	}{
		{[]int{}, 0},
		{[]int{0}, 1},
		{[]int{1}, 1},
		{[]int{1, 2, 255, 256, 1024}, 5},
	}

	for _, test := range tests {
		intset := NewIntSet()
		for _, num := range test.in {
			intset.Add(num)
		}
		if got := intset.Len(); got != test.want {
			t.Errorf("Len of %v = %d, want %d", intset, got, test.want)
		}
	}
}

func TestRemove(t *testing.T) {
	var tests = []struct {
		in     []int
		remove int
		want   string
	}{
		{[]int{}, 0, "{}"},
		{[]int{0}, 0, "{}"},
		{[]int{0}, 1, "{0}"},
		{[]int{1}, 0, "{1}"},
		{[]int{1}, 1, "{}"},
		{[]int{1, 2, 255, 256, 1024}, 256, "{1 2 255 1024}"},
		{[]int{1, 2, 255, 256, 1024}, 255, "{1 2 256 1024}"},
	}

	for _, test := range tests {
		intset := NewIntSet()
		for _, num := range test.in {
			intset.Add(num)
		}
		intset.Remove(test.remove)
		if got := intset.String(); got != test.want {
			t.Errorf("(%v).Remove(%d) -> %s, want %s", intset, test.remove, got, test.want)
		}
	}
}

func TestClear(t *testing.T) {
	var tests = []struct {
		in   []int
		want string
	}{
		{[]int{}, "{}"},
		{[]int{0}, "{}"},
		{[]int{1}, "{}"},
		{[]int{1, 2, 255, 256, 1024}, "{}"},
	}

	for _, test := range tests {
		intset := NewIntSet()
		for _, num := range test.in {
			intset.Add(num)
		}
		intset.Clear()
		if got := intset.String(); got != test.want {
			t.Errorf("(%v).Clear() -> %s, want %s", intset, got, test.want)
		}
	}
}

func TestCopy(t *testing.T) {
	var tests = []struct {
		in   []int
		want string
	}{
		{[]int{}, "{}"},
		{[]int{0}, "{0}"},
		{[]int{1}, "{1}"},
		{[]int{1, 2, 255, 256, 1024}, "{1 2 255 256 1024}"},
	}

	for _, test := range tests {
		intset := NewIntSet()
		for _, num := range test.in {
			intset.Add(num)
		}
		got := intset.Copy()
		if &intset == &got {
			t.Errorf("(%v).Copy() returns itself", intset)
		}
		if got := intset.String(); got != test.want {
			t.Errorf("(%v).Copy() -> %s, want %s", intset, got, test.want)
		}
	}
}

func TestAddAll(t *testing.T) {
	var tests = []struct {
		in, add []int
		want    string
	}{
		{[]int{}, []int{}, "{}"},
		{[]int{0}, []int{0}, "{0}"},
		{[]int{1}, []int{2}, "{1 2}"},
		{[]int{1, 3, 255, 1024}, []int{2, 3, 256, 1024}, "{1 2 3 255 256 1024}"},
	}

	for _, test := range tests {
		intset := NewIntSet()
		for _, num := range test.in {
			intset.Add(num)
		}
		intset.AddAll(test.add...)
		if got := intset.String(); got != test.want {
			t.Errorf("(%v).AddAll(%v) -> %s, want %s", test.in, test.add, got, test.want)
		}
	}
}

func TestIntersectWith(t *testing.T) {
	var tests = []struct {
		s, t []int
		want string
	}{
		{[]int{}, []int{}, "{}"},
		{[]int{0}, []int{0}, "{0}"},
		{[]int{1}, []int{2}, "{}"},
		{[]int{1, 3, 255, 1024}, []int{2, 3, 256, 1024}, "{3 1024}"},
	}

	for _, test := range tests {
		intsetS := NewIntSet()
		for _, num := range test.s {
			intsetS.Add(num)
		}
		intsetT := NewIntSet()
		for _, num := range test.t {
			intsetT.Add(num)
		}
		intsetS.IntersectWith(intsetT)
		if got := intsetS.String(); got != test.want {
			t.Errorf("(%v).IntersectWith(%v) -> %s, want %s", test.s, test.t, got, test.want)
		}
	}
}

func TestDifferenceWith(t *testing.T) {
	var tests = []struct {
		s, t []int
		want string
	}{
		{[]int{}, []int{}, "{}"},
		{[]int{0}, []int{0}, "{}"},
		{[]int{1}, []int{2}, "{1}"},
		{[]int{1, 3, 255, 1024}, []int{2, 3, 256, 1024}, "{1 255}"},
	}

	for _, test := range tests {
		intsetS := NewIntSet()
		for _, num := range test.s {
			intsetS.Add(num)
		}
		intsetT := NewIntSet()
		for _, num := range test.t {
			intsetT.Add(num)
		}
		intsetS.DifferenceWith(intsetT)
		if got := intsetS.String(); got != test.want {
			t.Errorf("(%v).DifferenceWith(%v) -> %s, want %s", test.s, test.t, got, test.want)
		}
	}
}

func TestSymmetricDifference(t *testing.T) {
	var tests = []struct {
		s, t []int
		want string
	}{
		{[]int{}, []int{}, "{}"},
		{[]int{0}, []int{0}, "{}"},
		{[]int{1}, []int{2}, "{1 2}"},
		{[]int{1, 3, 255, 1024}, []int{2, 3, 256, 1024}, "{1 2 255 256}"},
	}

	for _, test := range tests {
		intsetS := NewIntSet()
		for _, num := range test.s {
			intsetS.Add(num)
		}
		intsetT := NewIntSet()
		for _, num := range test.t {
			intsetT.Add(num)
		}
		intsetS.SymmetricDifference(intsetT)
		if got := intsetS.String(); got != test.want {
			t.Errorf("(%v).SymmetricDifference(%v) -> %s, want %s", test.s, test.t, got, test.want)
		}
	}
}

func TestElems(t *testing.T) {
	var tests = []struct {
		in []int
	}{
		{[]int{}},
		{[]int{0}},
		{[]int{1}},
		{[]int{1, 2, 255, 256, 1024}},
	}

	for _, test := range tests {
		intset := NewIntSet()
		for _, num := range test.in {
			intset.Add(num)
		}
		if got := intset.Elems(); !reflect.DeepEqual(got, test.in) {
			t.Errorf("(%v).Elems() = %v, want %v", test.in, got, test.in)
		}
	}
}
