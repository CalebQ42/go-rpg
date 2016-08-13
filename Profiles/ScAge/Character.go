package scage

import rpg "github.com/CalebQ42/go-rpg"

type Character struct {
	rpg.Character
	MaxHP       int
	HP          int
	ArmorRating int
	Strain      int
	MaxNT       int
	NT          int
	Level       int
	XP          int
}

func CreateEmptyScageCharacter(name string) (out Character) {
	out.Name = name
	out.Stats = EmptyStats
	return
}
