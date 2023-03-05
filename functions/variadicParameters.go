package main

import "fmt"

func main() {
	fmt.Println(foo(1, 2, 3, 4, 5, 6, 7, 8, 9))
}

// Here we the function foo expects a silce object. Which is actually dynamic sized array!!
// If we pass array then we get an error

func foo(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0

	for _, v := range x {
		sum += v
	}

	return sum
}
