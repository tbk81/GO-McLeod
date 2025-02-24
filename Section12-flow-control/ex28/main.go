package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Printf("x = %v\ty = %v\n", x, y)
	switch {
	case x < 4 && y < 4:
		fmt.Println("x and y are less than 4")
	case x > 4 && y > 6:
		fmt.Println("x and y are greater than 6")
	case x >= 4 && x <= 6:
		fmt.Println("x is greater than or equal to 4 and less than or equal to 6")
	case y != 5:
		fmt.Println("y is not 5")
	default:
		fmt.Println("none of the previous cases were met")
	}
}

/*
Modify the previous program to use a switch statement
Create 2 random ints between 0 inclusive and 10 exclusive
	○ assign them to variables with the identifiers x and y
		● Print their values
		● use an if statement to print the correct description
	○ x and y are both less than 4
	○ x and y are both greater than 6
	○ x is greater than or equal to 4 and less than or equal to 6
	○ y is not 5
	○ none of the previous cases were met
*/
