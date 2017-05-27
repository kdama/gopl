// ch05/ex15 は、可変子引数関数としての max や min の実装です。
package main

import "fmt"

func max(vals ...int) (int, error) {
	if len(vals) < 1 {
		return 0, fmt.Errorf("'vals' must have at least one value")
	}
	result := vals[0]
	for _, val := range vals {
		if result < val {
			result = val
		}
	}
	return result, nil
}

func min(vals ...int) (int, error) {
	if len(vals) < 1 {
		return 0, fmt.Errorf("'vals' must have at least one value")
	}
	result := vals[0]
	for _, val := range vals {
		if result > val {
			result = val
		}
	}
	return result, nil
}

func alternativeMax(val int, others ...int) int {
	result := val
	for _, other := range others {
		if result < other {
			result = other
		}
	}
	return result
}

func alternativeMin(val int, others ...int) int {
	result := val
	for _, other := range others {
		if result > other {
			result = other
		}
	}
	return result
}

func main() {
	if got, err := max(1, 2, 3, 4); err == nil {
		fmt.Println(got) // "4"
	}
	if got, err := min(1, 2, 3, 4); err == nil {
		fmt.Println(got) // "1"
	}

	fmt.Println(alternativeMax(1, 2, 3, 4)) // "4"
	fmt.Println(alternativeMin(1, 2, 3, 4)) // "1"
}
