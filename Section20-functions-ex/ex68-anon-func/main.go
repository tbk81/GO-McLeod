package main

import "fmt"

func main() {
	funcy := func(s string) string {
		return fmt.Sprint("This is an anon func - Your input was...", s)
	}("Gopherman!")
	fmt.Println(funcy)
}

/*
● Build and use an anonymous func
● anonymous func / func literal / lambda func
*/
