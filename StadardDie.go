package rpg

//CreateNumberDie creates a standard number die with the given amount of sides
func CreateNumberDie(sides int) (out Die) {
	for i := 1; i <= sides; i++ {
		out.AddSide(i)
	}
	return
}

//CreateStringDie creates a standard die with the given sides
func CreateStringDie(sides []string) (out Die) {
	for _, v := range sides {
		out.AddSide(v)
	}
	return
}
