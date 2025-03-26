package main

import "fmt"

func main() {
	fmt.Println(printSquare(square, 3))
}

func printSquare(f func(int) int, a int) string {
	x := f(a)
	return fmt.Sprintf("the number %v squared is %v", a, x)
}

func square(n int) int {
	return n * n
}

// func main() {
// 	result := process(square, 5)
// 	fmt.Println(result)
// }

// // callback function type
// type callbackFunc func(int) int

// // A function that takes a callback
// func process(cb callbackFunc, num int) int {
// 	return cb(num)
// }

// // Callback function
// func square(n int) int {
// 	return n * n
// }

/*
● pass a func into a func as an argument
	○ func square(n int) int
	○ printSquare(f func(int)int, int) string
*/
