package main

import (
	"testing"
)

func TestTopoSort(t *testing.T) {
	var tests = []struct {
		m    map[string][]string
		want []string
		err  bool
	}{
		{map[string][]string{}, []string{}, false},
		{map[string][]string{
			"A": {},
		}, []string{
			"A",
		}, false},
		{map[string][]string{
			"A": {},
			"B": {"A"},
			"C": {"B"},
		}, []string{
			"A", "B", "C",
		}, false},
		{map[string][]string{
			"A": {"B"},
			"B": {"C"},
			"C": {},
		}, []string{
			"C", "B", "A",
		}, false},
		{map[string][]string{
			"A": {"C"},
			"B": {"C"},
			"C": {},
		}, []string{
			"C", "A", "B",
		}, false},
		{map[string][]string{
			"A": {"A"},
		}, nil, true},
		{map[string][]string{
			"A": {"B"},
			"B": {"A"},
		}, nil, true},
		{map[string][]string{
			"A": {"B"},
			"B": {"C"},
			"C": {"A"},
		}, nil, true},
	}

	for _, test := range tests {
		got, err := topoSort(test.m)
		if err != nil && !test.err {
			t.Errorf("expects no error for topoSort(%v), but error: %v", test.m, err)
		} else if err == nil && test.err {
			t.Errorf("expects error for topoSort(%v), but no error", test.m)
		} else if len(got) != len(test.want) {
			t.Errorf("len(topoSort(%v)) = %d, want %d", test.m, len(got), len(test.want))
		} else {
			for i := range got {
				if got[i] != test.want[i] {
					t.Errorf("topoSort(%v)[%d] = %q, want %q", test.m, i, got[i], test.want[i])
				}
			}
		}
	}
}
