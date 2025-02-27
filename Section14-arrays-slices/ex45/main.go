package main

import "fmt"

var xi = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

func main() {
	y := []int{56, 57, 58, 59, 60}
	fmt.Println(xi)
	xi = append(xi, 52)
	fmt.Println(xi)
	xi = append(xi, 53, 54, 55)
	fmt.Println(xi)
	xi = append(xi, y...)
	fmt.Println(xi)

}

/*
start with this slice
	○ x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
● append to that slice this value
	○ 52
● print out the slice
● in ONE STATEMENT append to that slice these values
	○ 53
	○ 54
	○ 55
● print out the slice
● append to the slice this slice
	○ y := []int{56, 57, 58, 59, 60}
● print out the slice
*/
