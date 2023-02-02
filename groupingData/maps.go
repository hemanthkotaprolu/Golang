package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	v, ok := m["barnabas"]

	fmt.Println(v)

	fmt.Println(ok)

	// add something to the map

	m["todd"] = 33

	if v, ok := m["todd"]; ok {
		fmt.Println("the age of todd is: ", v)
	} else {
		fmt.Println("no element")
	}

	// var a map[string]int
	a := make(map[string]int)

	a["a"] = 25

	fmt.Println(a["a"])

	fmt.Println(len(m))

	fmt.Println("*************************")
	delete(m, "James")
	fmt.Println(len(m))
}
