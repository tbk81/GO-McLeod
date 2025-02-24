package main

import "fmt"

func main() {
	i := 0
	for {
		if i > 42 {
			break
		}
		fmt.Printf("Iteration: %v\n", i)
		i++
	}
}

/*
create a for loop that uses break statement
	● increment or decrement the variable being checked as a condition in the loop
*/
