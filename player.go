package rpg

//Player TODO
type Player struct {
	name          string
	stats         Stats
	armor         Armor
	shield        Shield
	weapons       []Weapon
	items         []Item
	equipedWeapon *Weapon
}

//CreatePlayer TODO
func CreatePlayer(name string, stats Stats) (out Player) {
	out.name = name
	out.stats = stats
	return
}

//Name TODO
func (p Player) Name() string {
	return p.name
}

//Stats TODO
func (p *Player) Stats() *Stats {
	return &p.stats
}

//SetArmor TODO
func (p *Player) SetArmor(armor Armor) {
	p.armor = armor
}

//Armor TODO
func (p *Player) Armor() *Armor {
	return &p.armor
}

//SetShield TODO
func (p *Player) SetShield(shield Shield) {
	p.shield = shield
}

//Shield TODO
func (p Player) Shield() *Shield {
	return &p.shield
}

//AddWeapon TODO
func (p *Player) AddWeapon(weapon Weapon) {
	p.weapons = append(p.weapons, weapon)
}

//Weapons TODO
func (p *Player) Weapons() *[]Weapon {
	return &p.weapons
}

//SetEquippedWeapon TODO
func (p *Player) SetEquippedWeapon(index int) {
	p.equipedWeapon = &p.weapons[index]
}

//AddItem TODO
func (p *Player) AddItem(item Item) {
	p.items = append(p.items, item)
}

//Item TODO
func (p Player) Item(index int) Item {
	return p.items[index]
}

//Inventory TODO
func (p *Player) Inventory() *[]Item {
	return &p.items
}
