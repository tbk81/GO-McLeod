package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
}

func main() {

	veh1 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "Hyundai",
		model: "ioniq",
		doors: 4,
		color: "grey",
	}

	veh2 := vehicle{
		engine: engine{
			electric: false,
		},
		make:  "Honda",
		model: "CR-V",
		doors: 4,
		color: "white",
	}

	fmt.Println(veh1)
	fmt.Println(veh2)

	fmt.Println(veh1.make)
	fmt.Println(veh2.make)

	fmt.Println(veh1.engine.electric)
	fmt.Println(veh2.engine.electric)
}

/*
●Create a type engine struct, and include this field
	○ electric bool
● Create a type vehicle struct, and include these fields
	■ engine
	■ make
	■ model
	■ doors
	■ color
● Create two VALUES of TYPE vehicle
	○ use a composite literal
● Print out each of these values.
● Print out a single field from each of these values.
*/
