package rpg

//Dice is a collection of Die
type Dice struct {
	pool []Die
}

//SetPool overwrites the dice pool with the new Die array
func (d *Dice) SetPool(pool []Die) {
	d.pool = pool
}

//AddDie Adds a die into the dice pool
func (d *Dice) AddDie(die Die) {
	d.pool = append(d.pool, die)
}

//Roll rolls all the dice and returns them as an interface array
func (d Dice) Roll() (out []interface{}) {
	for _, v := range d.pool {
		out = append(out, v.Roll())
	}
	return
}
