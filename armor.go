package rpg

//Armor is a simple instance of a piece of armor
type Armor struct {
	Name    string
	Rating  int
	Penalty int
	Strain  int
}

//CreateArmor creates armor
func CreateArmor(name string, rating, penalty, strain int) (out Armor) {
	out.Name = name
	out.Rating = rating
	out.Penalty = penalty
	out.Strain = strain
	return
}
