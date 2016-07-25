package rpg

import (
	"encoding/gob"
	"os"
)

//Character is a basic struct to represent a character.
type Character struct {
	Name            string
	Stats           BaseStats
	Skills          Skills
	Inventory       Inventory
	EquippedItems   []Item
	Weapons         []Weapon
	EquippedWeapons []Weapon
}

//CreateCharacter creates a basic character
func CreateCharacter(name string, stats BaseStats) (out Character) {
	out.Name = name
	out.Stats = stats
	return
}

//HasSkill return if the character has the given skill
func (c Character) HasSkill(skill string) bool {
	_, ok := c.Skills[skill]
	return ok
}

//EquipItem equips the item with the given name (if the player has that item in their
//inventory).
func (c *Character) EquipItem(name string) {
	for i, v := range c.Inventory {
		if v.Name == name {
			v.Equip(c)
			c.EquippedItems = append(c.EquippedItems, v)
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return
		}
	}
}

//UnequipItem unequips the item with the given name and places it into inventory (if the item is currently equipped).
func (c *Character) UnequipItem(name string) {
	for i, v := range c.EquippedItems {
		if v.Name == name {
			v.Unequip(c)
			c.Inventory = append(c.Inventory, v)
			c.EquippedItems = append(c.EquippedItems[:i], c.EquippedItems[i+1:]...)
			return
		}
	}
}

//Use quickly equips then unequips an item, consuming the item if consumed is set
//to true.
func (c *Character) Use(name string, consumed bool) {
	for i, v := range c.Inventory {
		if v.Name == name {
			v.Equip(c)
			v.Unequip(c)
			if consumed {
				c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			}
		}
	}
}

//GiveWeapon gives the character a weapon
func (c *Character) GiveWeapon(we Weapon) {
	c.Weapons = append(c.Weapons, we)
}

//EquipWeapon equips the weapon with the given name (given that the Character
//has the weapon)
func (c *Character) EquipWeapon(name string) {
	for i, v := range c.Weapons {
		if v.Name == name {
			c.EquippedWeapons = append(c.EquippedWeapons, v)
			c.Weapons = append(c.Weapons[:i], c.Weapons[i+1:]...)
			return
		}
	}
}

//UnequipWeapon unequips the weapon with the given name (given the weapon is equipped).
func (c *Character) UnequipWeapon(name string) {
	for i, v := range c.Weapons {
		if v.Name == name {
			c.Weapons = append(c.Weapons, v)
			c.EquippedWeapons = append(c.EquippedWeapons[:i], c.EquippedWeapons[i+1:]...)
			return
		}
	}
}

//Save saves a character to a file
func (c Character) Save(filename string) error {
	os.Remove(filename)
	fi, err := os.Create(filename)
	if err != nil {
		return err
	}
	err = gob.NewEncoder(fi).Encode(c)
	return err
}

//Load loads a character from a file
func (c *Character) Load(filename string) error {
	fi, err := os.Open(filename)
	if err != nil {
		return err
	}
	err = gob.NewDecoder(fi).Decode(c)
	return err
}
