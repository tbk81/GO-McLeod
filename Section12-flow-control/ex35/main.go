package main

import "fmt"

func main() {
	xi := []int{42, 43, 45, 46, 47, 48, 49, 50}
	for i, v := range xi {
		fmt.Printf("Index: %v\tValue: %v\n", i, v)
	}
}

/*
below is the code to create a data structure called a slice of ints
use a for range loop to print each value and the index position of each value
*/
