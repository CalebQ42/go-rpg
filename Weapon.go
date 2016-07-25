package rpg

import (
	"encoding/gob"
	"os"
)

//Weapon is a basic weapon. Damage is a string that is parsed when the weapon is used.
//Attack is a function that takes in the attacker and returns the amount of damage delt.
//
//Like Items, the name is the main way weapons are diferentiated so make sure
//two weapons don't have the same name.
type Weapon struct {
	Name   string
	Melee  bool
	Attack func(attacker Character) int
}

//CreateWeapon creates a weapon with the given name, damage, and melee bool
func CreateWeapon(name string, melee bool) (out Weapon) {
	out.Name = name
	out.Melee = melee
	return
}

//Save saves the weapon to a file
func (w Weapon) Save(filename string) error {
	os.Remove(filename)
	fi, err := os.Create(filename)
	if err != nil {
		return err
	}
	err = gob.NewEncoder(fi).Encode(w)
	return err
}

//Load loads the weapon from a file
func (w *Weapon) Load(filename string) error {
	fi, err := os.Open(filename)
	if err != nil {
		return err
	}
	err = gob.NewDecoder(fi).Decode(w)
	return err
}
