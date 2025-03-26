package main

import "fmt"

func main() {
	defer fmt.Println("This line runs after being defered")
	fmt.Println("This line runs first even though it's further down")
	fmt.Println("This line will run sec second")

}

/*
● “defer” multiple functions in main
	○ show that a deferred func runs after the func containing it exits.
	○ determine the order in which the multiple defer funcs run
*/
