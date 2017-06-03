package main

import (
	"testing"
)

func TestToStringSlice(t *testing.T) {
	tests := []struct {
		stack []string
		attrs []map[string]string
		want  []string
	}{
		{
			[]string{},
			[]map[string]string{},
			[]string{},
		},
		{
			[]string{},
			[]map[string]string{map[string]string{}},
			[]string{},
		},
		{
			[]string{"div"},
			[]map[string]string{map[string]string{}},
			[]string{"div"},
		},
		{
			[]string{"div"},
			[]map[string]string{map[string]string{"id": "foo"}},
			[]string{"div", "id=foo"},
		},
		{
			[]string{"div", "span"},
			[]map[string]string{map[string]string{"id": "foo"}, map[string]string{"class": "baz"}},
			[]string{"div", "id=foo", "span", "class=baz"},
		},
	}
	for _, test := range tests {
		got := toStringSlice(test.stack, test.attrs)
		if !equals(got, test.want) {
			t.Errorf("toStringSlice(%v, %v) = %v, want %v", test.stack, test.attrs, got, test.want)
		}
	}
}

func equals(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
