package rpg

import (
	"encoding/gob"
	"os"
)

//Item is a basic inventory item. If the item is equippable it is important to set
//Equippable to true, OnEquip, and OnUnequip. Because the name is the main way
//items are differentiated, make sure no two items have the same name if they are different.
//
//If equipping the item makes a permanent change set the OnEquip but not the OnUnequip.
//
//If an item is "used" (such as eating food or drinking a potion) it is quickly equipped
//then unequipped.
//
//Armor should be an Item.
type Item struct {
	Name        string
	Description string
	Equippable  bool
	OnEquip     func(*Character)
	OnUnequip   func(*Character)
}

//CreateItem is an easy way to create an item
func CreateItem(name, description string, equippable bool) (out Item) {
	out.Name = name
	out.Description = description
	out.Equippable = equippable
	return
}

//Equip runs the OnEquip function with the given character
func (i Item) Equip(char *Character) {
	i.OnEquip(char)
}

//Unequip runs the OnUnequip function with the given character
func (i Item) Unequip(char *Character) {
	i.OnUnequip(char)
}

//Save saves the item to a file
func (i Item) Save(filename string) error {
	os.Remove(filename)
	fi, err := os.Create(filename)
	if err != nil {
		return err
	}
	err = gob.NewEncoder(fi).Encode(i)
	return err
}

//Load loads an item from a file
func (i *Item) Load(filename string) error {
	fi, err := os.Open(filename)
	if err != nil {
		return err
	}
	err = gob.NewDecoder(fi).Decode(i)
	return err
}
