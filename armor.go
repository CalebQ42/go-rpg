package rpg

//Armor TODO
type Armor struct {
	name    string
	rating  int
	penalty int
	strain  int
}

//CreateArmor TODO
func CreateArmor(name string, rating, penalty, strain int) (out Armor) {
	out.name = name
	out.rating = rating
	out.penalty = penalty
	out.strain = strain
	return
}

//Name TODO
func (a *Armor) Name() string {
	return a.name
}

//Rating TODO
func (a *Armor) Rating() int {
	return a.rating
}

//Penalty TODO
func (a *Armor) Penalty() int {
	return a.penalty
}

//Strain TODO
func (a *Armor) Strain() int {
	return a.strain
}
