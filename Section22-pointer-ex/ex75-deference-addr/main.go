package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func main() {
	w := "My name is trevor"
	x := "There is a dog named Zora"
	y := "We live in a house"
	z := 44
	a = &w
	b = &x
	c = &y
	d = &z
	fmt.Printf("%v\t%T\n", a, a)
	fmt.Printf("%v\t%T\n", c, c)
	fmt.Printf("%v\t%T\n", d, d)
	fmt.Printf("%v\t%T\n", b, b)
	fmt.Println(*a)
	fmt.Println(*c)
	fmt.Println(*d)
	fmt.Println(*b)

}

/*
● print the VALUE stored in each variable
	○ these will be memory addresses
● print the TYPE of each variable
● print the data stored at memory locations
	○ dereference the stored memory address *
*/
