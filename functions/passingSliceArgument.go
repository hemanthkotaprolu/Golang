package main

import "fmt"

func main() {
    x1 := []int{1,2,3,4,5,6,7,8,9}

    fmt.Println(getSum(x1)
}

func getSum(x []int) int {
    sum := 0

    for _, v := range x {
        sum += v
    }
    return sum
}
