package xkcd

import (
	"testing"
)

func TestHit(t *testing.T) {
	var tests = []struct {
		comic *Comic
		terms []string
		want  bool
	}{
		{&Comic{}, []string{"foo"}, false},
		{&Comic{Alt: "foo"}, []string{"foo"}, true},
		{&Comic{Alt: "foo"}, []string{"bar"}, false},
		{&Comic{Alt: "foobar"}, []string{"foo", "bar"}, true},
		{&Comic{Alt: "foo", Day: "bar"}, []string{"foo", "bar"}, true},
	}

	for _, test := range tests {
		got := hit(test.comic, test.terms)
		if got != test.want {
			t.Errorf("hit(%v, %v) = %t, want %t", test.comic, test.terms, got, test.want)
		}
	}
}
