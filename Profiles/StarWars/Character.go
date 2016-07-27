package swrpg

import (
	"github.com/CalebQ42/go-rpg"
)

//Character is an extension of rpg.Character with the necessary variables for the Star Wars RPG
type Character struct {
	rpg.Character
	Species         string
	Xp              int
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

//CreateEmptySWCharacter creates a swrpg.Character with the given name and empty stats.
func CreateEmptySWCharacter(name string) (out Character) {
	out.Stats = EmptyStats
	out.Name = name
	return
}

//CreateSWCharacter creates a swrpg.Character with the given name and Stats
//for the species. Supported species should be found in the README.
func CreateSWCharacter(name, species string) (out Character) {
	out.Name = name
	out.Stats = EmptyStats
	switch species {
	case "Bothan":
		out.Xp = 100
		out.Stats["Brawn"] = 1
		out.Stats["Agility"] = 2
		out.Stats["Intellect"] = 2
		out.Stats["Cunning"] = 3
		out.Stats["Willpower"] = 2
		out.Stats["Presence"] = 2
		out.Skills["Streetwise"] = 1
	case "Droid":
		out.Xp = 175
		out.Stats["Brawn"] = 1
		out.Stats["Agility"] = 1
		out.Stats["Intellect"] = 1
		out.Stats["Cunning"] = 1
		out.Stats["Willpower"] = 1
		out.Stats["Presence"] = 1
	case "Duros":
		out.Xp = 100
		out.Stats["Brawn"] = 1
		out.Stats["Agility"] = 2
		out.Stats["Intellect"] = 3
		out.Stats["Cunning"] = 2
		out.Stats["Willpower"] = 2
		out.Stats["Presence"] = 2
		out.Skills["Piloting (Space)"] = 1
	case "Gran":
		out.Xp = 100
		out.Stats["Brawn"] = 2
		out.Stats["Agility"] = 2
		out.Stats["Intellect"] = 2
		out.Stats["Cunning"] = 1
		out.Stats["Willpower"] = 2
		out.Stats["Presence"] = 3
	case "Human":
		out.Xp = 110
		out.Stats["Brawn"] = 2
		out.Stats["Agility"] = 2
		out.Stats["Intellect"] = 2
		out.Stats["Cunning"] = 2
		out.Stats["Willpower"] = 2
		out.Stats["Presence"] = 2
	case "Ithorian":
		out.Xp = 90
		out.Stats["Brawn"] = 2
		out.Stats["Agility"] = 1
		out.Stats["Intellect"] = 2
		out.Stats["Cunning"] = 2
		out.Stats["Willpower"] = 3
		out.Stats["Presence"] = 2
		out.Skills["Survival"] = 1
	case "Mon Calamari":
		out.Xp = 100
		out.Stats["Brawn"] = 2
		out.Stats["Agility"] = 2
		out.Stats["Intellect"] = 3
		out.Stats["Cunning"] = 1
		out.Stats["Willpower"] = 2
		out.Stats["Presence"] = 2
		out.Skills["Knowledge (Education)"] = 1
	case "Sullustan":
		out.Xp = 100
		out.Stats["Brawn"] = 2
		out.Stats["Agility"] = 3
		out.Stats["Intellect"] = 2
		out.Stats["Cunning"] = 1
		out.Stats["Willpower"] = 2
		out.Stats["Presence"] = 2
		out.Skills["Astrogation"] = 1
	case "Gand":
		out.Xp = 100
		out.Stats["Brawn"] = 2
		out.Stats["Agility"] = 2
		out.Stats["Intellect"] = 2
		out.Stats["Cunning"] = 2
		out.Stats["Willpower"] = 3
		out.Stats["Presence"] = 1
		out.Skills["Discipline"] = 1
	case "Rodian":
		out.Xp = 100
		out.Stats["Brawn"] = 2
		out.Stats["Agility"] = 3
		out.Stats["Intellect"] = 2
		out.Stats["Cunning"] = 2
		out.Stats["Willpower"] = 1
		out.Stats["Presence"] = 2
		out.Skills["Survival"] = 1
	case "Trandoshan":
		out.Xp = 90
		out.Stats["Brawn"] = 3
		out.Stats["Agility"] = 1
		out.Stats["Intellect"] = 2
		out.Stats["Cunning"] = 2
		out.Stats["Willpower"] = 2
		out.Stats["Presence"] = 2
		out.Skills["Perception"] = 1
	case "Twi'lek":
		out.Xp = 100
		out.Stats["Brawn"] = 1
		out.Stats["Agility"] = 2
		out.Stats["Intellect"] = 2
		out.Stats["Cunning"] = 2
		out.Stats["Willpower"] = 2
		out.Stats["Presence"] = 3
	case "Wookiee":
		out.Xp = 90
		out.Stats["Brawn"] = 3
		out.Stats["Agility"] = 2
		out.Stats["Intellect"] = 2
		out.Stats["Cunning"] = 2
		out.Stats["Willpower"] = 1
		out.Stats["Presence"] = 2
		out.Skills["Brawl"] = 1
	case "Cerean":
		out.Xp = 90
		out.Stats["Brawn"] = 2
		out.Stats["Agility"] = 1
		out.Stats["Intellect"] = 3
		out.Stats["Cunning"] = 2
		out.Stats["Willpower"] = 2
		out.Stats["Presence"] = 2
		out.Skills["Vigilance"] = 1
	case "Kel Dor":
		out.Xp = 100
		out.Stats["Brawn"] = 1
		out.Stats["Agility"] = 2
		out.Stats["Intellect"] = 2
		out.Stats["Cunning"] = 2
		out.Stats["Willpower"] = 3
		out.Stats["Presence"] = 2
		out.Skills["Knowledge (Education)"] = 1
	case "Mirialan":
		out.Xp = 100
		out.Stats["Brawn"] = 2
		out.Stats["Agility"] = 3
		out.Stats["Intellect"] = 2
		out.Stats["Cunning"] = 1
		out.Stats["Willpower"] = 2
		out.Stats["Presence"] = 2
		out.Skills["Discipline"] = 1
		out.Skills["Cool"] = 1
	case "Nautolan":
		out.Xp = 100
		out.Stats["Brawn"] = 3
		out.Stats["Agility"] = 2
		out.Stats["Intellect"] = 2
		out.Stats["Cunning"] = 2
		out.Stats["Willpower"] = 1
		out.Stats["Presence"] = 2
		out.Skills["Athletics"] = 1
	case "Togruta":
		out.Xp = 100
		out.Stats["Brawn"] = 1
		out.Stats["Agility"] = 2
		out.Stats["Intellect"] = 2
		out.Stats["Cunning"] = 3
		out.Stats["Willpower"] = 2
		out.Stats["Presence"] = 2
		out.Skills["Perception"] = 1
	case "Zabrak":
		out.Xp = 100
		out.Stats["Brawn"] = 2
		out.Stats["Agility"] = 2
		out.Stats["Intellect"] = 2
		out.Stats["Cunning"] = 2
		out.Stats["Willpower"] = 3
		out.Stats["Presence"] = 1
	case "Besalisk":
		out.Xp = 85
		out.Stats["Brawn"] = 3
		out.Stats["Agility"] = 1
		out.Stats["Intellect"] = 2
		out.Stats["Cunning"] = 1
		out.Stats["Willpower"] = 2
		out.Stats["Presence"] = 2
		out.Skills["Resilience"] = 1
	case "Dug":
		out.Xp = 90
		out.Stats["Brawn"] = 2
		out.Stats["Agility"] = 3
		out.Stats["Intellect"] = 2
		out.Stats["Cunning"] = 2
		out.Stats["Willpower"] = 2
		out.Stats["Presence"] = 1
		out.Skills["Brawl"] = 1
	case "Mustafarian":
		out.Xp = 100
		out.Stats["Brawn"] = 2
		out.Stats["Agility"] = 2
		out.Stats["Intellect"] = 2
		out.Stats["Cunning"] = 2
		out.Stats["Willpower"] = 2
		out.Stats["Presence"] = 1
	case "Caamasi":
		out.Xp = 100
		out.Stats["Brawn"] = 2
		out.Stats["Agility"] = 2
		out.Stats["Intellect"] = 2
		out.Stats["Cunning"] = 1
		out.Stats["Willpower"] = 3
		out.Stats["Presence"] = 2
	case "Neimoidian":
		out.Xp = 90
		out.Stats["Brawn"] = 1
		out.Stats["Agility"] = 2
		out.Stats["Intellect"] = 3
		out.Stats["Cunning"] = 3
		out.Stats["Willpower"] = 1
		out.Stats["Presence"] = 2
	case "Gossam":
		out.Xp = 100
		out.Stats["Brawn"] = 1
		out.Stats["Agility"] = 2
		out.Stats["Intellect"] = 2
		out.Stats["Cunning"] = 3
		out.Stats["Willpower"] = 2
		out.Stats["Presence"] = 2
		out.Skills["Deception"] = 1
	case "Chadra-Fan":
		out.Xp = 90
		out.Stats["Brawn"] = 1
		out.Stats["Agility"] = 3
		out.Stats["Intellect"] = 3
		out.Stats["Cunning"] = 2
		out.Stats["Willpower"] = 1
		out.Stats["Presence"] = 2
		out.Skills["Mechanics"] = 1
	case "Dressellian":
		out.Xp = 110
		out.Stats["Brawn"] = 2
		out.Stats["Agility"] = 2
		out.Stats["Intellect"] = 2
		out.Stats["Cunning"] = 2
		out.Stats["Willpower"] = 3
		out.Stats["Presence"] = 1
		out.Skills["Survival"] = 1
	case "Xexto":
		out.Xp = 85
		out.Stats["Brawn"] = 1
		out.Stats["Agility"] = 3
		out.Stats["Intellect"] = 2
		out.Stats["Cunning"] = 2
		out.Stats["Willpower"] = 1
		out.Stats["Presence"] = 2
		out.Skills["Confidence"] = 1
	case "Aqualish":
		out.Xp = 90
		out.Stats["Brawn"] = 3
		out.Stats["Agility"] = 2
		out.Stats["Intellect"] = 1
		out.Stats["Cunning"] = 2
		out.Stats["Willpower"] = 2
		out.Stats["Presence"] = 2
		out.Skills["Brawl"] = 1
	case "Klatooinian":
		out.Xp = 100
		out.Stats["Brawn"] = 2
		out.Stats["Agility"] = 3
		out.Stats["Intellect"] = 2
		out.Stats["Cunning"] = 2
		out.Stats["Willpower"] = 1
		out.Stats["Presence"] = 2
	case "Weequay":
		out.Xp = 90
		out.Stats["Brawn"] = 3
		out.Stats["Agility"] = 2
		out.Stats["Intellect"] = 1
		out.Stats["Cunning"] = 3
		out.Stats["Willpower"] = 2
		out.Stats["Presence"] = 1
	case "Chiss":
		out.Xp = 90
		out.Stats["Brawn"] = 2
		out.Stats["Agility"] = 2
		out.Stats["Intellect"] = 3
		out.Stats["Cunning"] = 2
		out.Stats["Willpower"] = 2
		out.Stats["Presence"] = 1
		out.Skills["Cool"] = 1
	case "Toydarian":
		out.Xp = 90
		out.Stats["Brawn"] = 1
		out.Stats["Agility"] = 1
		out.Stats["Intellect"] = 2
		out.Stats["Cunning"] = 2
		out.Stats["Willpower"] = 3
		out.Stats["Presence"] = 3
	case "Arcona":
		out.Xp = 100
		out.Stats["Brawn"] = 1
		out.Stats["Agility"] = 2
		out.Stats["Intellect"] = 2
		out.Stats["Cunning"] = 2
		out.Stats["Willpower"] = 3
		out.Stats["Presence"] = 2
		out.Skills["Vigilance"] = 1
	case "Chevin":
		out.Xp = 80
		out.Stats["Brawn"] = 3
		out.Stats["Agility"] = 1
		out.Stats["Intellect"] = 2
		out.Stats["Cunning"] = 3
		out.Stats["Willpower"] = 2
		out.Stats["Presence"] = 1
		out.Skills["Negotiation"] = 1
	case "Falleen":
		out.Xp = 90
		out.Stats["Brawn"] = 1
		out.Stats["Agility"] = 2
		out.Stats["Intellect"] = 2
		out.Stats["Cunning"] = 2
		out.Stats["Willpower"] = 2
		out.Stats["Presence"] = 3
		out.Skills["Charm"] = 1
	case "Gotal":
		out.Xp = 100
		out.Stats["Brawn"] = 2
		out.Stats["Agility"] = 2
		out.Stats["Intellect"] = 2
		out.Stats["Cunning"] = 3
		out.Stats["Willpower"] = 2
		out.Stats["Presence"] = 1
		out.Skills["Perception"] = 1
	case "Quarren":
		out.Xp = 95
		out.Stats["Brawn"] = 2
		out.Stats["Agility"] = 2
		out.Stats["Intellect"] = 1
		out.Stats["Cunning"] = 2
		out.Stats["Willpower"] = 3
		out.Stats["Presence"] = 2
		out.Skills["Negotiation"] = 1
	case "Hutt":
		out.Xp = 70
		out.Stats["Brawn"] = 3
		out.Stats["Agility"] = 1
		out.Stats["Intellect"] = 2
		out.Stats["Cunning"] = 2
		out.Stats["Willpower"] = 3
		out.Stats["Presence"] = 2
	case "Gank":
		out.Xp = 110
		out.Stats["Brawn"] = 2
		out.Stats["Agility"] = 2
		out.Stats["Intellect"] = 2
		out.Stats["Cunning"] = 2
		out.Stats["Willpower"] = 2
		out.Stats["Presence"] = 1
	case "Nikto":
		out.Xp = 100
		out.Stats["Brawn"] = 3
		out.Stats["Agility"] = 2
		out.Stats["Intellect"] = 2
		out.Stats["Cunning"] = 2
		out.Stats["Willpower"] = 3
		out.Stats["Presence"] = 1
	case "Sakiyan":
		out.Xp = 80
		out.Stats["Brawn"] = 2
		out.Stats["Agility"] = 2
		out.Stats["Intellect"] = 3
		out.Stats["Cunning"] = 2
		out.Stats["Willpower"] = 2
		out.Stats["Presence"] = 2
	case "Drall":
		out.Xp = 90
		out.Stats["Brawn"] = 1
		out.Stats["Agility"] = 1
		out.Stats["Intellect"] = 4
		out.Stats["Cunning"] = 2
		out.Stats["Willpower"] = 2
		out.Stats["Presence"] = 2
		out.Skills["Knowledge (Education)"] = 1
	case "Selonian":
		out.Xp = 80
		out.Stats["Brawn"] = 2
		out.Stats["Agility"] = 3
		out.Stats["Intellect"] = 2
		out.Stats["Cunning"] = 1
		out.Stats["Willpower"] = 3
		out.Stats["Presence"] = 1
		out.Skills["Coordination"] = 1
	}
	return
}

//StartingStats figures out the starting wound threshold, strain threshold, and soak value for the character's species. Run after editing the base 6 stats.
func (c *Character) StartingStats() {
	c.SoakVal = c.Stats["Brawn"]
	switch c.Species {
	case "Bothan":
		c.WoundThreshold = 10 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 11 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Droid":
		c.WoundThreshold = 10 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 10 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Duros":
		c.WoundThreshold = 11 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 10 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Gran":
		c.WoundThreshold = 10 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 9 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Human":
		c.WoundThreshold = 10 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 10 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Ithorian":
		c.WoundThreshold = 9 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 12 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Mon Calamari":
		c.WoundThreshold = 10 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 10 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Sullustan":
		c.WoundThreshold = 10 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 10 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Gand":
		c.WoundThreshold = 10 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 10 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Rodian":
		c.WoundThreshold = 10 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 10 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Trandoshan":
		c.WoundThreshold = 12 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 9 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Twi'lek":
		c.WoundThreshold = 10 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 11 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Wookiee":
		c.WoundThreshold = 14 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 8 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Cerean":
		c.WoundThreshold = 10 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 13 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Kel Dor":
		c.WoundThreshold = 10 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 10 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Mirialan":
		c.WoundThreshold = 11 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 10 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Nautolan":
		c.WoundThreshold = 11 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 9 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Togruta":
		c.WoundThreshold = 10 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 10 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Zabrak":
		c.WoundThreshold = 10 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 10 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Besalisk":
		c.WoundThreshold = 12 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 7 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Dug":
		c.WoundThreshold = 9 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 8 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Mustafarian":
		c.WoundThreshold = 10 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 10 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Caamasi":
		c.WoundThreshold = 10 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 11 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Neimoidian":
		c.WoundThreshold = 11 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 9 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Gossam":
		c.WoundThreshold = 9 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 11 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Chadra-Fan":
		c.WoundThreshold = 9 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 11 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Dressellian":
		c.WoundThreshold = 11 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 10 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Xexto":
		c.WoundThreshold = 9 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 9 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Aqualish":
		c.WoundThreshold = 11 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 8 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Klatooinian":
		c.WoundThreshold = 10 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 10 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Weequay":
		c.WoundThreshold = 10 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 9 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Chiss":
		c.WoundThreshold = 10 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 10 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Toydarian":
		c.WoundThreshold = 11 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 10 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Arcona":
		c.WoundThreshold = 10 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 10 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Chevin":
		c.WoundThreshold = 11 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 11 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Falleen":
		c.WoundThreshold = 10 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 12 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Gotal":
		c.WoundThreshold = 9 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 8 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Quarren":
		c.WoundThreshold = 10 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 8 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Hutt":
		c.WoundThreshold = 13 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 11 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Gank":
		c.WoundThreshold = 10 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 10 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Nikto":
		c.WoundThreshold = 11 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 10 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Sakiyan":
		c.WoundThreshold = 8 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 10 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Drall":
		c.WoundThreshold = 8 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 12 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	case "Selonian":
		c.WoundThreshold = 11 + c.Stats["Brawn"]
		c.WoundCurrent = c.WoundThreshold
		c.StrainThreshold = 10 + c.Stats["Willpower"]
		c.StrainCurrent = c.StrainThreshold
	}
}
