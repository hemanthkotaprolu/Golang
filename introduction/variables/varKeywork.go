package main

import "fmt"

// var keyword is used to create the variables with global context
// shorhand declaration is local scope

// BEST PRACTICE
//	Try to limit the scope as much as possible(means limit reducing "var")

func main() {
	// short declaration
	x := 42

	fmt.Println(x)

	// var keyword

	var y = 40

	fmt.Println(y)
}
