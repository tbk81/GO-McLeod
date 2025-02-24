package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			continue
		} else {
			fmt.Println(i)
		}
	}
}

/*
use modulus and the continue statement in a loop to print all ODD numbers
*/
