package main

import "fmt"

func main() {
	foo()
	bar("Trevor")
	z := zora("Fido")
	fmt.Println(z)
	chick1, a := chicken("Poppy", 1)
	fmt.Println(chick1, a)
	fmt.Println(chick1)
	fmt.Println(a)
}

// func no params
func foo() {
	fmt.Println("I am from foo")
}

// func 1 param, no returns
func bar(s string) {
	fmt.Println("My name is", s)
}

// func 1 param, 1 return
func zora(s string) string {
	return fmt.Sprint("I am a dog named ", s)
}

// func 2 params, 2 returns
func chicken(name string, age int) (string, int) {
	age *= 10
	return fmt.Sprint(name, " is this old in human years:"), age
}

/*
func (r reciever) indentifier (p parameter(s)) (return(s)) {...}
*/
