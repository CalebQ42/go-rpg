package rpg

//Weapon TODO
type Weapon struct {
	name    string
	atk     func() int
	ranged  bool
	sRange  int
	lRange  int
	loaded  bool
	maxAmmo int
	ammo    int
}

//CreateWeapon TODO
func CreateWeapon(name string, attack func() int) (out Weapon) {
	out.name = name
	out.atk = attack
	out.ranged = false
	out.loaded = true
	return
}

//SetRange TODO
func (w *Weapon) SetRange(short, long int) {
	w.ranged = true
	w.sRange = short
	w.lRange = long
}

//SetMaxAmmo TODO
func (w *Weapon) SetMaxAmmo(max int) {
	w.maxAmmo = max
}

//MaxAmmo TODO
func (w Weapon) MaxAmmo() int {
	return w.maxAmmo
}

//SetAmmo TODO
func (w *Weapon) SetAmmo(ammo int) {
	w.ammo = ammo
}

//Ammo TODO
func (w Weapon) Ammo() int {
	return w.ammo
}

//Loaded TODO
func (w Weapon) Loaded() bool {
	return w.loaded
}

//Attack TODO
func (w Weapon) Attack() int {
	return w.atk()
}
