package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

type myNumbers interface {
	~int | ~float64 // ~ include all values of type int that have an underlying value of int
}

// This ia a generic
func addT[T myNumbers](a, b T) T {
	return a + b
}

type myAlias int

func main() {
	var n myAlias = 42
	fmt.Println(add(1, 2))
	fmt.Println(addF(1.2, 2.5))

	fmt.Println("------Generic func------")
	// Using gernerics
	fmt.Println(addT(n, 2))
	fmt.Println(addT(1.2, 2.5))
}
