package main

import "fmt"

var y = 42
var z string = "hemanth"

// DECLARED VARIABLE can't be changed.
// GO is a STATICALLY TYPED language
// variable is declared only to store a particular type!!

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Printf("%T\n", z)
}
