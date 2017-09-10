package cycle

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCycle(t *testing.T) {
	one := 1

	type CyclePtr *CyclePtr
	var cyclePtr1, cyclePtr2 CyclePtr
	cyclePtr1 = &cyclePtr1
	cyclePtr2 = &cyclePtr2

	type CycleSlice []CycleSlice
	var cycleSlice = make(CycleSlice, 1)
	cycleSlice[0] = cycleSlice

	ch1 := make(chan int)

	type mystring string

	var iface1 interface{} = &one

	for _, test := range []struct {
		x    interface{}
		want bool
	}{
		// basic types
		{1, false},
		{"foo", false},
		{mystring("foo"), false},
		// slices
		{[]string{"foo"}, false},
		{[]string{}, false},
		{[]string(nil), false},
		// slice cycles
		{cycleSlice, true},
		// maps
		{
			map[string][]int{"foo": {1, 2, 3}},
			false,
		},
		{
			map[string][]int{},
			false,
		},
		{
			map[string][]int(nil),
			false,
		},
		// pointers
		{&one, false},
		{new(bytes.Buffer), false},
		// pointer cycles
		{cyclePtr1, true},
		{cyclePtr2, true},
		// functions
		{(func())(nil), false},
		{func() {}, false},
		// arrays
		{[...]int{1, 2, 3}, false},
		// channels
		{ch1, false},
		// interfaces
		{&iface1, false},
	} {
		if Cycle(test.x) != test.want {
			t.Errorf("Cycle(%v) = %t",
				test.x, !test.want)
		}
	}
}

func Example() {
	//!+cycle
	// Circular linked lists a -> b -> a and c -> c.
	type link struct {
		value string
		tail  *link
	}
	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c
	fmt.Println(Cycle(a)) // "true"
	fmt.Println(Cycle(b)) // "true"
	fmt.Println(Cycle(c)) // "true"
	//!-cycle

	// Output:
	// true
	// true
	// true
}
