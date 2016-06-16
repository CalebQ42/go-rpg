package rpg

//Player TODO
type Player struct {
	Name          string
	Stats         Stats
	Armor         Armor
	Shield        Shield
	Weapons       []Weapon
	Inventory     []Item
	EquipedWeapon *Weapon
}

//CreatePlayer TODO
func CreatePlayer(name string, stats Stats) (out Player) {
	out.Name = name
	out.Stats = stats
	return
}

//AddWeapon TODO
func (p *Player) AddWeapon(weapon Weapon) {
	p.Weapons = append(p.Weapons, weapon)
}

//SetEquippedWeapon TODO
func (p *Player) SetEquippedWeapon(index int) {
	p.EquipedWeapon = &p.Weapons[index]
}

//AddItem TODO
func (p *Player) AddItem(item Item) {
	p.Inventory = append(p.Inventory, item)
}
