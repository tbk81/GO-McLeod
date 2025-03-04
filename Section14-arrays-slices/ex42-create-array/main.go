package main

import "fmt"

func main() {
	ar := [5]int{5, 6, 7, 8, 9}
	for x := range ar {
		fmt.Println(ar[x])
	}
}

/*
Using a COMPOSITE LITERAL:
● create an ARRAY which holds 5 VALUES of TYPE int
● assign VALUES to each index position.
● Range over the array and print the values out.
	○ print out the VALUE and the TYPE
*/
