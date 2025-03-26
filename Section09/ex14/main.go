package main

import "fmt"

var x string

const y = 42

func main() {
	x = "I'm outside main block!"
	j := "This is inside main func"
	fmt.Printf("Outside var = %v\nConst var = %v\nInside var = %v\n", j, x, y)
}

/*
create the following variables with the following scopes:
- package level
	■ create outside of func main
	■ use the
		● var keyword
		● const keyword

- block level
	■ inside func main
		● use the short declaration operator
*/
