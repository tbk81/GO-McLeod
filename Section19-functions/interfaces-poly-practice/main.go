package main

import "fmt"

type character struct {
	stats
	name       string
	background string
	species    string
	class      string
}

type stats struct {
	str   int
	dex   int
	con   int
	intel int
	wis   int
	cha   int
}

func (c character) identify() {
	fmt.Printf("My name - %v\nBackground in - %v\nMy species - %v\nMy class - %v\n", c.name, c.background, c.species, c.class)
}

func (s stats) identify() {
	fmt.Printf("Str: %v\nDex: %v\nCon: %v\nInt: %v\nWis: %v\nCha: %v\n", s.str, s.dex, s.con, s.intel, s.wis, s.cha)
}

type player interface {
	identify()
}

func whoami(p player) {
	p.identify()
}

func main() {
	p1 := character{
		stats: stats{
			str:   10,
			dex:   15,
			con:   13,
			intel: 9,
			wis:   8,
			cha:   11,
		},
		name:       "Trevor",
		background: "Criminal",
		species:    "Human",
		class:      "Rogue",
	}

	p2 := character{
		stats: stats{
			str:   15,
			dex:   12,
			con:   11,
			intel: 9,
			wis:   8,
			cha:   10,
		},
		name:       "Zora",
		background: "Cloistered Scholar",
		species:    "Dog",
		class:      "Paladin",
	}

	whoami(p1)
	whoami(p1.stats)
	fmt.Println("----------------------")
	whoami(p2)
	whoami(p2.stats)
}
