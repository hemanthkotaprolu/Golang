package main

import "fmt"

func main() {
    s1 := foo()
    fmt.Println(s1)

    f1 := bar()
    fmt.Println(f1())
}



func foo() string {
    return "Hello world"
}

func bar() func() int {
    return func() int {
        return 451
    }
}
