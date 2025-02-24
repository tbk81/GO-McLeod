package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("Iterateion %v\tx is 3\n", i)
		}
	}
}

/*
● use the "statement statement" idiom to
	○ initialize x with and random int between 0 inclusive and 5 exclusive
	○ if x is 3
		■ print "x is 3"
● run that code 100 times
*/
