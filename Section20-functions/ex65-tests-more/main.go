package main

import "fmt"

func main() {
	fmt.Println(saySomething("Hellooooo gopherman"))
}

func saySomething(s string) string {
	return fmt.Sprint("This is something I said to print, ", s)
}
