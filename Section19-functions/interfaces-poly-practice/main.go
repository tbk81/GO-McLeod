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

func (c character) indentify() {
	fmt.Printf("My name - %v\nBackground in - %v\nMy species - %v\n My class - %v\n", c.name, c.background, c.species, c.class)
}

func (s stats) statout() {
	fmt.Printf("Str: %v\nDex: %v\nCon: %v\nInt: %v\nWis: %v\nCha: %v\n", s.str, s.dex, s.con, s.intel, s.wis, s.cha)
}

type player interface {
	identify()
	statout()
}

func whoami(p player) {
	p.identify()
	p.statout()
}

func main() {
	p1 := character{
		name:       "Trevor",
		background: "Criminal",
		species:    "Human",
		class:      "Rogue",
		stats: stats{
			str:   10,
			dex:   15,
			con:   13,
			intel: 9,
			wis:   8,
			cha:   11,
		},
	}

	p2 := character{
		name:       "Zora",
		background: "Cloistered Scholar",
		species:    "Dog",
		class:      "Paladin",
		stats: stats{
			str:   15,
			dex:   12,
			con:   11,
			intel: 9,
			wis:   8,
			cha:   10,
		},
	}

	whoami(p1)
}
