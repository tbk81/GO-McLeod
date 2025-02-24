package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}

	for j := 0; j < 100; j++ {
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
}

/*
there are two parts ot this hands on exercise
	○ create a program that has a loop that prints every number from 0 to 99
	○ modify the program from the previous hands on exercise to run 100 times
*/
