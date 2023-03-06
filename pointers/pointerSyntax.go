package main

import "fmt"

func main() {
    a := 42

    b := &a

    fmt.Println(&a)
    fmt.Printf("%T\n", a)

    fmt.Println(*b)
    fmt.Printf("%T\n", b)
}
