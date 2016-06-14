package rpg

//AgeDice TODO
type AgeDice [3]Die

//CreateFantasyDice TODO
func CreateFantasyDice() (out AgeDice) {
	for i := range out {
		out[i] = CreateDie(6)
	}
	return
}

//Roll TODO
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
