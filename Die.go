package rpg

import (
	"crypto/rand"
	"math/big"
)

//Die is a basic multi-sided die
type Die struct {
	sides []interface{}
}

//AddSide adds a side to the die. Doesn't check input value is slice
func (d *Die) AddSide(side interface{}) {
	d.sides = append(d.sides, side)
}

//Roll returns a random side of the die
func (d Die) Roll() interface{} {
	res, _ := rand.Int(rand.Reader, big.NewInt(int64(len(d.sides))))
	return d.sides[res.Int64()]
}
