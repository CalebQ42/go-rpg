package rpg

//Dice is a collection of die
type Dice []Die

//CreateDice creates dice.
func CreateDice(numDice, diceSides int) (out Dice) {
	for i := 0; i < numDice; i++ {
		out.AddDie(CreateDie(diceSides))
	}
	return
}

//AddDie adds a die to a dice collection
func (d *Dice) AddDie(de Die) {
	*d = append(*d, de)
}

//RollTotal rolls all the dice and adds up the result.
func (d Dice) RollTotal() (out int) {
	for _, v := range d {
		out += v.Roll()
	}
	return
}

//RollIndividual rolls all the dice and returns an array of all the results.
func (d Dice) RollIndividual() (out []int) {
	for _, v := range d {
		out = append(out, v.Roll())
	}
	return
}

//RollBoth returns both the total of each dice's result, and also returns an array with each die's roll
func (d Dice) RollBoth() (tot int, ind []int) {
	for _, v := range d {
		rl := v.Roll()
		tot += rl
		ind = append(ind, rl)
	}
	return
}
