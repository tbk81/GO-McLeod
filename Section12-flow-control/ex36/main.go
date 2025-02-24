package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 40,
	}
	for k, v := range m {
		fmt.Printf("Key: %v \tAge: %v\n", k, v)
	}
}

/*
below is the code to create a data structure called a map
put this code into a program
m := map[string]int{
"James":
42,
"Moneypenny": 32,
}
● use a for range loop to print each value and the key associated with each value
*/
