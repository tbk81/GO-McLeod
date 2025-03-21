package main

import "fmt"

func main() {
	p1 := person{
		first: "trevor",
		age:   44,
	}

	p1.speak()
}

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Printf("My name is %v and I am %v years old.\n", p.first, p.age)
}

/*
● Create a user defined struct with
	○ the identifier “person”
	○ the fields:
		■ first
		■ age
● attach a method to type person with
	○ the identifier “speak”
	○ the method should have the person say their name and age
● create a value of type person
● call the method from the value of type person
*/
