package rpg

//Shield TODO
type Shield struct {
	name  string
	bonus int
}

//CreateShield TODO
func CreateShield(name string, bonus int) (out Shield) {
	out.name = name
	out.bonus = bonus
	return
}

//Name TODO
func (s Shield) Name() string {
	return s.name
}

//Bonus TODO
func (s Shield) Bonus() int {
	return s.bonus
}
