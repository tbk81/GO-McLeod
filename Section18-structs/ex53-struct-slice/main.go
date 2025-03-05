package main

import "fmt"

type person struct {
	first   string
	last    string
	flavors []string
}

func main() {
	p1 := person{
		first:   "trevor",
		last:    "kimbler",
		flavors: []string{"chocolate", "vanilla", "straberry"},
	}

	p2 := person{
		first:   "zora",
		last:    "dog",
		flavors: []string{"lamb", "beef", "chicken"},
	}

	fmt.Printf("First: %v\nLast: %v\nFlavors: %v\n", p1.first, p1.last, p1.flavors)
	fmt.Println("--------------------------------")
	for _, v := range p1.flavors {
		fmt.Println(p1.first, p1.last)
		fmt.Println(v)

	}
	fmt.Println("--------------------------------")
	for _, v := range p2.flavors {
		fmt.Println(p2.first, p2.last)
		fmt.Println(v)
	}
}

/*
Create your own type “person” which will have an underlying type of “struct” so that it can store
the following data:
	● first name
	● last name
	● favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice
which stores the favorite flavors.
*/
