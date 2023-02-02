package main

import "fmt"

func main() {
	sayHello()

	sayHelloName("hemanth")

	s1 := sayHii("Hemanth")
	fmt.Println(s1)

	x, y := mouse("Hemanth", "Kotaprolu")
	fmt.Println(x)
	fmt.Println(y)
	// fmt.Println("Hello")
}

// func (r receiver) identifier(parameters) (return(s)) { ... }

func sayHello() {
	fmt.Println("Hello")
}

// everything in GO is "PASS BY VALUE"
func sayHelloName(s string) {
	fmt.Println("Hello", s)
}

func sayHii(str string) string {
	return fmt.Sprint("Hello from woo, ", str)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, `, says Hello`)
	b := false

	return a, b
}
