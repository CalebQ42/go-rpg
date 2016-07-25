package rpg

//CreateNumberDie creates a standard number die with the given amount of sides
func CreateNumberDie(sides int) (out Die) {
	var sds []int
	for i := 1; i <= sides; i++ {
		sds = append(sds, i)
	}
	out.SetSides(sds)
	return
}

//CreateStringDie creates a standard die with the given sides
func CreateStringDie(sides []string) (out Die) {
	out.SetSides(sides)
	return
}
