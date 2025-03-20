package main

import "fmt"

func foo() int {
	return 42
}

func bar() (int, string) {
	return 44, "trevor"
}

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

/*
○ create a func with the identifier foo that returns an int
○ create a func with the identifier bar that returns an int and a string
○ call both funcs
○ print out their results
*/
