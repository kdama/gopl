package github

import (
	"testing"
)

func TestMilestoneEquals(t *testing.T) {
	var tests = []struct {
		m    Milestone
		x    *Milestone
		want bool
	}{
		{Milestone{}, &Milestone{}, true},
		{Milestone{}, &Milestone{ID: 42}, false},
		{Milestone{ID: 42}, &Milestone{ID: 42}, true},
		{Milestone{ID: 42}, &Milestone{ID: 43}, false},
	}

	for _, test := range tests {
		got := test.m.Equals(test.x)
		if got != test.want {
			t.Errorf("(%q).Equals(%q) = %t, want %t", test.m, test.x, got, test.want)
		}
	}
}

func TestUserEquals(t *testing.T) {
	var tests = []struct {
		u    User
		x    *User
		want bool
	}{
		{User{}, &User{}, true},
		{User{}, &User{ID: 42}, false},
		{User{ID: 42}, &User{ID: 42}, true},
		{User{ID: 42}, &User{ID: 43}, false},
	}

	for _, test := range tests {
		got := test.u.Equals(test.x)
		if got != test.want {
			t.Errorf("(%q).Equals(%q) = %t, want %t", test.u, test.x, got, test.want)
		}
	}
}
