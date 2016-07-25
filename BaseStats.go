package rpg

//BaseStats holds a players base stats. Is map[string]int type.
type BaseStats map[string]int

//CreateBaseStats initializes the BaseStats map and assigns the startingVal to each
//string in the stats array.
func CreateBaseStats(stats []string, startingVal int) (out BaseStats) {
	out = make(map[string]int)
	for _, v := range stats {
		out[v] = startingVal
	}
	return
}
