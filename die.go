package rpg

import (
	"math/rand"
	"time"
)

//Die is a single die
type Die struct {
	Sides int
	ran   *rand.Rand
}

//CreateDie creates a single die with the specified number of sides
func CreateDie(sides int) (out Die) {
	out.Sides = sides
	out.ran = rand.New(rand.NewSource(time.Now().UnixNano() - int64(sides)))
	return
}

//Roll rolls the die and returns the result
func (d Die) Roll() int {
	return d.ran.Intn(d.Sides) + 1
}
