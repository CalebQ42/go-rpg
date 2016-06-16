package rpg

//Shield TODO
type Shield struct {
	Name  string
	Bonus int
}

//CreateShield TODO
func CreateShield(name string, bonus int) (out Shield) {
	out.Name = name
	out.Bonus = bonus
	return
}
