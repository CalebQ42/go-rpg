package rpg

//Dice TODO
type Dice []Die

//CreateDice TODO
func CreateDice(numDice, diceSides int) (out Dice) {
	for i := 0; i < numDice; i++ {
		out.AddDie(CreateDie(diceSides))
	}
	return
}

//AddDie TODO
func (d *Dice) AddDie(de Die) {
	*d = append(*d, de)
}

//RollTotal TODO
func (d Dice) RollTotal() (out int) {
	for _, v := range d {
		out += v.Roll()
	}
	return
}

//RollIndividual TODO
func (d Dice) RollIndividual() (out []int) {
	for _, v := range d {
		out = append(out, v.Roll())
	}
	return
}

//RollBoth TODO
func (d Dice) RollBoth() (tot int, ind []int) {
	tot = d.RollTotal()
	ind = d.RollIndividual()
	return
}
