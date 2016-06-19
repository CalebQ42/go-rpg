package rpg

//AgeDice is 3 die specially for the age system.
type AgeDice [3]Die

//CreateFantasyDice Creates an instance of AgeDice
func CreateFantasyDice() (out AgeDice) {
	for i := range out {
		out[i] = CreateDie(6)
	}
	return
}

//Roll the AgeDice. Returns the total, then the stunt die result, and says if stunt points are generated
func (f AgeDice) Roll() (res, stunt int, genSP bool) {
	var resInd []int
	for _, v := range f {
		tmp := v.Roll()
		res += tmp
		resInd = append(resInd, tmp)
	}
	stunt = resInd[2]
	genSP = resInd[0] == resInd[1] || resInd[0] == resInd[2] || resInd[1] == resInd[2]
	return
}
