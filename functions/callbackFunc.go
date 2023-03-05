package main

import "fmt"

// Callbacks are funcitons that are passed as an arguments

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := even(sum, ii...) // ... at the last makes the silce an variadic parameter. that means it unfolds the slice

	fmt.Println(s)
}

func sum(xi ...int) int {
	total := 0

	for _, v := range xi {
		total += v
	}

	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int

	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}

	return sum(yi...)
}
