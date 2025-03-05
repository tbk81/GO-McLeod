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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v)
		for _, v2 := range v.flavors {
			fmt.Println(v2)
		}
	}
}

/*
Take the code from the previous exercise, then store the VALUES of type person in a map with
the KEY of last name. Access each value in the map. Print out the values, ranging over the
slice.
*/
