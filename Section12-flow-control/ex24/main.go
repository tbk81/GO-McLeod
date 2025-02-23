package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)
	fmt.Printf("x = %v\n", x)
	switch {
	case x <= 100:
		fmt.Println("Number is between 0 and 100")
	case x > 100 && x <= 200:
		fmt.Println("Number is between 101 and 200")
	default:
		fmt.Println("Number is between 201 and 250")
	}
}

/*
create a program that uses both SEQUENTIAL and CONDITIONAL control flow. Your
program should do the following
	○ create a random int between 0 and 250
	○ store the value of that int in a variable with the identifier of x
		■ func Intn(n int) int
			● rand.Intn()
	○ print out the name and value of the variable
	use an if statement to do the following
		■ if the value is between 0 and 100
			● print between 0 and 100
		■ if the value is between 101 and 200
			● print between 101 and 200
		■ if the value is between 201 and 250
			● print between 201 and 250
			● re: inclusive, exclusive – does rand.Intn()
	○ include zero in its output?
	○ include 250 in its output?
	○ show this in code using the numbers 0 - 3
			● hint:
				○ &&
*/
