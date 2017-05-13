package main

import (
	"testing"

	"github.com/kdama/gopl/ch04/ex14/github"
)

func TestAppendMilestoneAsSet(t *testing.T) {
	var tests = []struct {
		set       []github.Milestone
		milestone *github.Milestone
		want      []github.Milestone
	}{
		{[]github.Milestone{}, &github.Milestone{}, []github.Milestone{github.Milestone{}}},
		{[]github.Milestone{github.Milestone{}}, &github.Milestone{}, []github.Milestone{github.Milestone{}}},
		{[]github.Milestone{}, &github.Milestone{ID: 42}, []github.Milestone{github.Milestone{ID: 42}}},
		{[]github.Milestone{github.Milestone{ID: 42}}, &github.Milestone{ID: 42}, []github.Milestone{github.Milestone{ID: 42}}},
		{[]github.Milestone{github.Milestone{ID: 42}}, &github.Milestone{ID: 43}, []github.Milestone{github.Milestone{ID: 42}, github.Milestone{ID: 43}}},
	}

	for _, test := range tests {
		got := appendMilestoneAsSet(test.set, test.milestone)
		if len(got) != len(test.want) {
			t.Errorf("len(appendMilestoneAsSet(%q, %q)) = %d, want %d", test.set, test.milestone, len(got), len(test.want))
		} else {
			for idx, gotValue := range got {
				if gotValue != test.want[idx] {
					t.Errorf("appendMilestoneAsSet(%q, %q)[%d] = %q, want %q", test.set, test.milestone, idx, gotValue, test.want[idx])
				}
			}
		}
	}
}

func TestIncludesMilestone(t *testing.T) {
	var tests = []struct {
		array     []github.Milestone
		milestone *github.Milestone
		want      bool
	}{
		{[]github.Milestone{}, &github.Milestone{}, false},
		{[]github.Milestone{github.Milestone{}}, &github.Milestone{}, true},
		{[]github.Milestone{github.Milestone{ID: 42}}, &github.Milestone{ID: 42}, true},
		{[]github.Milestone{github.Milestone{ID: 42}}, &github.Milestone{ID: 43}, false},
	}

	for _, test := range tests {
		if got := includesMilestone(test.array, test.milestone); got != test.want {
			t.Errorf("includesMilestone(%q, %q) = %t, want %t", test.array, test.milestone, got, test.want)
		}
	}
}

func TestAppendUserAsSet(t *testing.T) {
	var tests = []struct {
		set  []github.User
		user *github.User
		want []github.User
	}{
		{[]github.User{}, &github.User{}, []github.User{github.User{}}},
		{[]github.User{github.User{}}, &github.User{}, []github.User{github.User{}}},
		{[]github.User{}, &github.User{ID: 42}, []github.User{github.User{ID: 42}}},
		{[]github.User{github.User{ID: 42}}, &github.User{ID: 42}, []github.User{github.User{ID: 42}}},
		{[]github.User{github.User{ID: 42}}, &github.User{ID: 43}, []github.User{github.User{ID: 42}, github.User{ID: 43}}},
	}

	for _, test := range tests {
		got := appendUserAsSet(test.set, test.user)
		if len(got) != len(test.want) {
			t.Errorf("len(appendUserAsSet(%q, %q)) = %d, want %d", test.set, test.user, len(got), len(test.want))
		} else {
			for idx, gotValue := range got {
				if gotValue != test.want[idx] {
					t.Errorf("appendUserAsSet(%q, %q)[%d] = %q, want %q", test.set, test.user, idx, gotValue, test.want[idx])
				}
			}
		}
	}
}

func TestIncludesUser(t *testing.T) {
	var tests = []struct {
		array []github.User
		user  *github.User
		want  bool
	}{
		{[]github.User{}, &github.User{}, false},
		{[]github.User{github.User{}}, &github.User{}, true},
		{[]github.User{github.User{ID: 42}}, &github.User{ID: 42}, true},
		{[]github.User{github.User{ID: 42}}, &github.User{ID: 43}, false},
	}

	for _, test := range tests {
		if got := includesUser(test.array, test.user); got != test.want {
			t.Errorf("includesUser(%q, %q) = %t, want %t", test.array, test.user, got, test.want)
		}
	}
}
