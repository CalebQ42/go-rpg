package rpg

//Player TODO
type Player struct {
	Name          string
	Stats         Stats
	Armor         Armor
	Shield        Shield
	Weapons       []Weapon
	Items         []Item
	EquipedWeapon *Weapon
}

//CreatePlayer TODO
func CreatePlayer(name string, stats Stats) (out Player) {
	out.Name = name
	out.Stats = stats
	return
}

//SetArmor TODO
func (p *Player) SetArmor(armor Armor) {
	p.Armor = armor
}

//SetShield TODO
func (p *Player) SetShield(shield Shield) {
	p.Shield = shield
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
	p.Items = append(p.Items, item)
}

//Item TODO
func (p Player) Item(index int) Item {
	return p.Items[index]
}

//Inventory TODO
func (p *Player) Inventory() *[]Item {
	return &p.Items
}
