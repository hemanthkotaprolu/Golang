package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am ", s.first, s.last, "- the secretAgent speaks")
}

func (p person) speak() {
	fmt.Println("I am ", p.first, p.last, "- the person speaks")
}

type human interface {
    speak()
}

func bar(h human)  {
    fmt.Println("I was passed into bar", h)
    
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}
    
	sa2 := secretAgent{

		person: person{
			first: "Miss",
			last:  "Moneypenny",
		},
		ltk: true,
	}

    p1 := person {
        first: "Dr.",
        last: "Yes",
    }

	fmt.Println(sa1)
	sa1.speak()
    sa2.speak()

    fmt.Println(p1)

    bar(sa1)
    bar(sa2)
    bar(p1)
}
