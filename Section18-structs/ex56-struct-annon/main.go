package main

import "fmt"

// type person struct{
// 	first string
// 	friends map[string]int
// 	favDrinks []string
// }

func main() {
	p1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first:     "trevor",
		friends:   map[string]int{"frank": 42, "zora": 11, "john": 32},
		favDrinks: []string{"cherry coke", "coffee", "cranberry juice"},
	}

	fmt.Println(p1.first, "Friends")
	for k, v := range p1.friends {
		fmt.Println(k, v)
	}
	fmt.Println("----------------")
	fmt.Println(p1.first, "Favorite Drinks")
	for _, v := range p1.favDrinks {
		fmt.Println(v)
	}
}

/*
Create and use an anonymous struct with these fields:
	● first string
	● friends map[string]int
	● favDrinks []string
Print things.
*/
