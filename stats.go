package rpg

//Stats is a basic way of holding base stats
type Stats struct {
	Stats map[string]int
}

//CreateStats creates a stats instance with the stats in the string array and sets it's value to defaultValue
func CreateStats(stats []string, defaultValue int) (out Stats) {
	out.Stats = make(map[string]int)
	for _, v := range stats {
		out.Stats[v] = defaultValue
	}
	return
}

//StatValue returns a stat's value
func (s *Stats) StatValue(stat string) int {
	return s.Stats[stat]
}

//SetStat sets a stat's value
func (s *Stats) SetStat(stat string, value int) {
	s.Stats[stat] = value
}
