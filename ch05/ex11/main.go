// ch05/ex11 は、循環が存在した場合にそれを報告する topoSort です。
package main

import (
	"fmt"
	"log"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},

	"linear algebra": {"calculus"},
}

func main() {
	sorted, err := topoSort(prereqs)
	if err != nil {
		log.Fatalf("ch05/ex11: %v", err)
	}
	for i, course := range sorted {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string) error

	visitAll = func(items []string) error {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				err := visitAll(m[item])
				if err != nil {
					return err
				}
				order = append(order, item)
			} else {
				cyclic := true
				for _, s := range order {
					if s == item {
						cyclic = false
					}
				}
				if cyclic {
					return fmt.Errorf("cyclic: %s", item)
				}
			}
		}
		return nil
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	err := visitAll(keys)
	if err != nil {
		return nil, err
	}
	return order, nil
}
