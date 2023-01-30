package main

import "fmt"

func main() {

	x := 2

	fmt.Printf("%d\t\t%b", x, x)
	fmt.Println()

	y := x << 1
	fmt.Printf("%d\t\t%b", y, y)
	fmt.Println()

	kb := 1024
	mb := kb * 1024
	gb := mb * 1024

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t\t%b\n", gb, gb)
}
