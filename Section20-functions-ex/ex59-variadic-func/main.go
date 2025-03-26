package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(foo(numbers...))
	fmt.Println(bar(numbers))
}

func foo(nums ...int) int {
	sum := 0
	for xi := range nums {
		sum += xi
	}
	return sum
}

func bar(slicenums []int) int {
	sumnum := 0
	for xint := range slicenums {
		sumnum += xint
	}
	return sumnum
}

/*
● create a func with the identifier foo that
	○ takes in a variadic parameter of type int
	○ pass in a value of type []int into your func (unfurl the []int)
	○ returns the sum of all values of type int passed in
● create a func with the identifier bar that
	○ takes in a parameter of type []int
	○ returns the sum of all values of type int passed in
*/
