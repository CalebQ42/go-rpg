package rpg

//Stats TODO
type Stats struct {
	stats map[string]int
}

//CreateStats TODO
func CreateStats(stats []string, defaultValue int) (out Stats) {
	for _, v := range stats {
		out.stats[v] = defaultValue
	}
	return
}

//StatValue TODO
func (s *Stats) StatValue(stat string) int {
	return s.stats[stat]
}

//SetStat TODO
func (s *Stats) SetStat(stat string, value int) {
	s.stats[stat] = value
}

//Map TODO
func (s *Stats) Map() *map[string]int {
	return &s.stats
}
