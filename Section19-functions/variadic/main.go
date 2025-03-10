package main

import "fmt"

func main() {
	s := sum(3, 5, 6, 8, 10)
	fmt.Println("The sum is", s)
}

func sum(ii ...int) int {
	fmt.Println(ii)
	fmt.Printf("%T\n", ii)

	n := 0
	for _, v := range ii {
		n += v
	}
	return n
}

/*
func (r reciever) indentifier (p parameter(s)) (return(s)) {...}
*/
