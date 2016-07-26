package swrpg

import (
	"github.com/CalebQ42/go-rpg"
)

//Character is an extension of rpg.Character with the necessary variables for the Star Wars RPG
type Character struct {
	rpg.Character
	Morality        int
	Credits         int
	ForceRating     int
	SoakVal         int
	WoundThreshold  int
	WoundCurrent    int
	StrainThreshold int
	StrainCurrent   int
	DefenseRanged   int
	DefenseMelee    int
}

//CreateBaseSWCharacter creates a swrpg.Character with the given name and rpg.BaseStats.
//Also figures out the starting SoakVal.
func CreateBaseSWCharacter(name string, stats rpg.BaseStats) (out Character) {
	out.Stats = stats
	out.Name = name
	out.SoakVal = out.Stats["Brawn"]
	return
}
