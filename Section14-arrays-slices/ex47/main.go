package main

import "fmt"

var states = []string{`Alabama`, `Alaska`, `Arizona`, `Arkansas`, `California`, `Colorado`, `Connecticut`, `Delaware`, `Florida`, `Georgia`, `Hawaii`, `Idaho`, `Illinois`, `Indiana`, `Iowa`, `Kansas`, `Kentucky`, `Louisiana`, `Maine`, `Maryland`, `Massachusetts`, `Michigan`, `Minnesota`, `Mississippi`, `Missouri`, `Montana`, `Nebraska`, `Nevada`, `New Hampshire`, `New Jersey`, `New Mexico`, `New York`, `North Carolina`, `North Dakota`, `Ohio`, `Oklahoma`, `Oregon`, `Pennsylvania`, ` Rhode Island`, `South Carolina`, `South Dakota`, `Tennessee`, `Texas`, `Utah`, `Vermont`, `Virginia`, `Washington`, `West Virginia`, `Wisconsin`, `Wyoming`}

func main() {
	fmt.Printf("List len: %v\nList cap: %v\n", len(states), cap(states))
	for i := 0; i < len(states); i++ {
		fmt.Printf("index: %v\tstate: %v\n", i, states[i])
	}
}

/*
● Create a slice to store the names of all of the states in the United States of America.
	○ Use make and append to do this.
	○ Goal: do not have the array that underlies the slice created more than once.
● Print out
	○ the len
	○ the cap
	○ the values, along with their index position, without using the range clause.
*/
