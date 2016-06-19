package rpg

//InitPlay is a player and initiative pair
type InitPlay struct {
	Player     Player
	Initiative int
}

//CreateInitPlay creates an InitPlay instance
func CreateInitPlay(player Player, initiative int) (out InitPlay) {
	out.Player = player
	out.Initiative = initiative
	return
}

//InitPlays is a collection of InitPlay
type InitPlays []InitPlay

func (ip InitPlays) Len() int {
	return len(ip)
}

func (ip InitPlays) Less(i, j int) bool {
	return ip[i].Initiative > ip[j].Initiative
}

func (ip InitPlays) Swap(i, j int) {
	ip[i], ip[j] = ip[j], ip[i]
}
