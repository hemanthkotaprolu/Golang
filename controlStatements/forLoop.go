package main

import "fmt"

func main() {
    // method - 1
    for i := 0; i <= 20; i++ {
        fmt.Print(" ")
        fmt.Print(i)
    }
    fmt.Println()

    // method - 2
    x := 1
    for x < 15 {
        fmt.Println(x)
        x++
    }
}
