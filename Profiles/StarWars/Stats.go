package swrpg

import (
	"github.com/CalebQ42/go-rpg"
)

var (
	//EmptyStats is an rpg.BaseStats with the six stats used in the Star Wars RPG set to 0
	EmptyStats = rpg.CreateBaseStats([]string{"Brawn", "Agility", "Intellect", "Cunning", "Willpower", "Presence"}, 0)
)
