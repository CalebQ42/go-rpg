package rpg

import (
	"math/rand"
	"time"
)

//Die TODO
type Die struct {
	sides int
	ran   *rand.Rand
}

//CreateDie TODO
func CreateDie(sides int) (out Die) {
	out.sides = sides
	out.ran = rand.New(rand.NewSource(time.Now().UnixNano() - int64(sides)))
	return
}

//Roll TODO
func (d Die) Roll() int {
	return d.ran.Intn(d.sides) + 1
}

//Sides TODO
func (d Die) Sides() int {
	return d.sides
}
