package main

import "fmt"


func main() {
    f := func(){
        fmt.Println("inside the func expression")
    }

    f()
}
