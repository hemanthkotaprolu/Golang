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

    // use break statement to exit the loop
    // use continue to go to next iteration without executing the next lines of code in the loop

    for i := 33; i <= 122; i++ {
        fmt.Printf("%v\t%#U\n", i, i)
    }
}
