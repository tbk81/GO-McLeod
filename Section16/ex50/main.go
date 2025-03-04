package main

import "fmt"

var m = map[string][]string{
	"bond_james":       {`shaken, not stirred`, `martinis`, `fast cars`},
	"moneypenny_jenny": {`intelligence`, `literature`, `computer science`},
	"no_dr":            {`cats`, `ice cream`, `sunsets`},
}

func main() {
	m[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}
	for k, v := range m {
		fmt.Println(k)
		for i, val := range v {
			fmt.Println(i, val)
		}
	}
}
