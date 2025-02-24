package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("Outer loop: %v\t", i)
			fmt.Printf("Inner loop: %v\n", j)
		}
	}
}
