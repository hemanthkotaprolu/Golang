package main

import "fmt"

var a int

type hotdog int
var b hotdog

func main() {
   a = 20

   b = 35


   fmt.Println(a)
   fmt.Println(b)

   fmt.Printf("%T\n", a)
   fmt.Printf("%T\n", b)

   a = int(b)

   fmt.Println(a)
}
