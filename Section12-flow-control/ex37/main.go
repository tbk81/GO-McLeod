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

	age := m["James"]
	fmt.Printf("Age of James: %v\n", age)

	for k := range m {
		if k != "Q" {
			fmt.Println("Q not found!")
		}
	}

	if v, ok := m["Q"]; !ok {
		fmt.Println("\"Q\" is not found! Value is ", v)
	}
}

/*
use the code from the previous exercise
add this code to print a single value stored in the map

age := m["James"]
fmt.Println(age)


now use similar code to use the lookup of "Q" and print that value
now use the "comma ok" idiom to test whether "Q" is actually stored in the map, then
print out a statement if it is not stored in the map
	○ hint: check effective go for the "comma ok" idiom
*/
