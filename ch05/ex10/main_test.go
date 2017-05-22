package main

import (
	"testing"
)

func TestTopoSort(t *testing.T) {
	var tests = []struct {
		m    map[string][]string
		want []string
	}{
		{map[string][]string{}, []string{}},
		{map[string][]string{
			"A": {},
		}, []string{
			"A",
		}},
		{map[string][]string{
			"A": {},
			"B": {"A"},
			"C": {"B"},
		}, []string{
			"A", "B", "C",
		}},
		{map[string][]string{
			"A": {"B"},
			"B": {"C"},
			"C": {},
		}, []string{
			"C", "B", "A",
		}},
		{map[string][]string{
			"A": {"C"},
			"B": {"C"},
			"C": {},
		}, []string{
			"C", "A", "B",
		}},
	}

	for _, test := range tests {
		got := topoSort(test.m)
		if len(got) != len(test.want) {
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
