package rpg

import (
	"strconv"
	"strings"
)

//Weapon is a universal weapon
type Weapon struct {
	Name    string
	Damage  string
	Ranged  bool
	SRange  int
	LRange  int
	Loaded  bool
	MaxAmmo int
	Ammo    int
}

//CreateWeapon creates a weapon with the specified name and damage. Damage is in
//The form of 2d6+3. The + or - has to come after the die
func CreateWeapon(name string, damage string) (out Weapon) {
	out.Name = name
	out.Damage = damage
	out.Ranged = false
	out.Loaded = true
	return
}

//SetRange Sets the range of the weapon. Also sets the weapon as ranged.
func (w *Weapon) SetRange(short, long int) {
	w.Ranged = true
	w.SRange = short
	w.LRange = long
}

//RollDamage returns the damage dealt by the weapon.
func (w *Weapon) RollDamage() (out int) {
	lwr := strings.ToLower(w.Damage)
	ind := strings.IndexAny(lwr, "+-")
	indd := strings.Index(lwr, "d")
	if ind != -1 {
		ad, _ := strconv.Atoi(lwr[ind:])
		out += ad
		dieNum, _ := strconv.Atoi(lwr[:indd])
		dieSides, _ := strconv.Atoi(lwr[indd+1 : ind])
		di := CreateDice(dieNum, dieSides)
		out += di.RollTotal()
	} else {
		dieNum, _ := strconv.Atoi(lwr[:indd])
		dieSides, _ := strconv.Atoi(lwr[indd+1:])
		di := CreateDice(dieNum, dieSides)
		out += di.RollTotal()
	}
	return
}
