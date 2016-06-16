package rpg

//Weapon TODO
type Weapon struct {
	Name    string
	Atk     func() int
	Ranged  bool
	SRange  int
	LRange  int
	Loaded  bool
	MaxAmmo int
	Ammo    int
}

//CreateWeapon TODO
func CreateWeapon(name string, attack func() int) (out Weapon) {
	out.Name = name
	out.Atk = attack
	out.Ranged = false
	out.Loaded = true
	return
}

//SetRange TODO
func (w *Weapon) SetRange(short, long int) {
	w.Ranged = true
	w.SRange = short
	w.LRange = long
}

//Attack TODO
func (w Weapon) Attack() int {
	return w.Atk()
}
