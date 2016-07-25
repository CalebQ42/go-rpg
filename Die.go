package rpg

import (
	"crypto/rand"
	"math/big"
	"reflect"
)

//Die is a basic multi-sided die
type Die struct {
	sides []interface{}
}

//SetSides can take in both an array or single value, but either way it overwrites
//The sides already added. If array adds each element of the array.
func (d *Die) SetSides(array interface{}) {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		var sds []interface{}
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			sds = append(sds, s.Index(i))
		}
		d.sides = sds
	default:
		var empt []interface{}
		empt = append(empt, array)
		d.sides = empt
	}
}

//AddSide adds a side to the die. Doesn't check input value is slice
func (d *Die) AddSide(side interface{}) {
	d.sides = append(d.sides, side)
}

//AddSides should take in an array of values and adds each element of the array.
//Checks if it's an array first.
func (d *Die) AddSides(array interface{}) {
	if reflect.TypeOf(array).Kind() == reflect.Slice {
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			d.sides = append(d.sides, s.Index(i))
		}
	}
}

//Roll returns a random side of the die
func (d Die) Roll() interface{} {
	res, _ := rand.Int(rand.Reader, big.NewInt(int64(len(d.sides))))
	return d.sides[res.Int64()]
}
