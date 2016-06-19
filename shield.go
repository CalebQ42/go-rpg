package rpg

//Shield is a basic shield struct
type Shield struct {
	Name  string
	Bonus int
}

//CreateShield creates a shield with the specified name and bonus
func CreateShield(name string, bonus int) (out Shield) {
	out.Name = name
	out.Bonus = bonus
	return
}
