package main

import "fmt"

var xi = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

func main() {
	s1 := xi[:5]
	s2 := xi[5:]
	s3 := xi[2:7]
	s4 := xi[1:6]

	fmt.Printf("s1: %v\ns2: %v\ns3: %v\ns4: %v\n", s1, s2, s3, s4)
}

/*
Using the code from the previous example, use SLICING to create the following new slices
which are then printed:
● [42 43 44 45 46]
● [47 48 49 50 51]
● [44 45 46 47 48]
● [43 44 45 46 47]
*/
