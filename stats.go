package rpg

//Stats TODO
type Stats struct {
	Stats map[string]int
}

//CreateStats TODO
func CreateStats(stats []string, defaultValue int) (out Stats) {
	out.Stats = make(map[string]int)
	for _, v := range stats {
		out.Stats[v] = defaultValue
	}
	return
}

//StatValue TODO
func (s *Stats) StatValue(stat string) int {
	return s.Stats[stat]
}

//SetStat TODO
func (s *Stats) SetStat(stat string, value int) {
	s.Stats[stat] = value
}
