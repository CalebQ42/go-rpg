package rpg

//Player is a simple player struct
type Player struct {
	Name          string
	Stats         Stats
	Armor         Armor
	Shield        Shield
	Weapons       []Weapon
	Inventory     []Item
	EquipedWeapon *Weapon
}

//CreatePlayer creates a player with the specified name and stats
func CreatePlayer(name string, stats Stats) (out Player) {
	out.Name = name
	out.Stats = stats
	return
}

//AddWeapon adds a weapon into the player's weapon array
func (p *Player) AddWeapon(weapon Weapon) {
	p.Weapons = append(p.Weapons, weapon)
}

//SetEquippedWeapon equips the weapon in the player's weapon array
func (p *Player) SetEquippedWeapon(index int) {
	p.EquipedWeapon = &p.Weapons[index]
}

//AddItem Adds an item into the players inventory
func (p *Player) AddItem(item Item) {
	p.Inventory = append(p.Inventory, item)
}
