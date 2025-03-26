package main

import "fmt"

func main() {
	n := 5
	fmt.Printf("%v multiplied by 2 is: %v\n", n, call(5))
}

var call = func(i int) int {
	return i * 2
}

/*
● Assign a func to a variable, then call that func
*/
