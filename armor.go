package rpg

//Armor TODO
type Armor struct {
	Name    string
	Rating  int
	Penalty int
	Strain  int
}

//CreateArmor TODO
func CreateArmor(name string, rating, penalty, strain int) (out Armor) {
	out.Name = name
	out.Rating = rating
	out.Penalty = penalty
	out.Strain = strain
	return
}
