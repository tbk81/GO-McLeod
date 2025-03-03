package main

import "fmt"

// var bondy = [][]string{
// 	{"James", "Bond", "Shaken, not stirred"},
// 	{"Miss", "Moneypenny", "I'm 008."},
// }

var jb = []string{"James", "Bond", "Shaken, not stirred"}
var mp = []string{"Miss", "Moneypenny", "I'm 008."}
var bondy = [][]string{jb, mp}

func main() {
	for i, v := range bondy {
		fmt.Println(i, v)
		for x, y := range v {
			fmt.Println(x, y)
		}
	}
}

/*
Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional
slice:
"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "I'm 008."

Range over the records, then range over the data in each record.
*/
