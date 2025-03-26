package main

import "fmt"

func main() {
	x := 42
	fmt.Println("This is the value of x:", x)
	fmt.Printf("This is the type of x: %T\n", x)
	fmt.Println("This is the mem address of x:", &x)
	fmt.Println("This is the dereferenced value of x:", *&x)
	fmt.Println("-------------------------------")
	y := &x
	fmt.Println("This is the value of y:", y)
	fmt.Println("This is the dereferenced value of y:", *y)
	fmt.Printf("This is the type of y: %T\n", y)
	fmt.Println("This is the mem address of y:", &y)
	fmt.Println("-------------------------------")
	*y = 43
	fmt.Println("This is the value of y:", y)
	fmt.Println("This is the dereferenced value of y:", *y)
	fmt.Printf("This is the type of y: %T\n", y)
	fmt.Println("This is the mem address of y:", &y)
}
