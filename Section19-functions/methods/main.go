package main

import "fmt"

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("My name is", p.first)
}

func main() {
	p1 := person{
		first: "Trevor",
	}

	p2 := person{
		first: "Zora",
	}

	p1.speak()
	p2.speak()
}

/*
func (r reciever) indentifier (p parameter(s)) (return(s)) {...}
*/
