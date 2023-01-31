package main

import "fmt"

func main() {

	// x := type{values}

	x := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(len(x))

	// for range
	for i, v := range x {
		fmt.Printf("%v - %v\n", i, v)
	}

	// slicing a slice
	y := x[1:4]
	fmt.Println(y)

	// append a slice
	z := append(x, x...)

	fmt.Println(len(z))

	// if the size of the slice exceed then the length doubled
	a := make([]int, 5, 10)
	fmt.Println(a)
	fmt.Println(cap(a))
}
