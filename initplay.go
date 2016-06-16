package rpg

//InitPlay TODO
type InitPlay struct {
	Play       Player
	Initiative int
}

//CreateInitPlay TODO
func CreateInitPlay(player Player, initiative int) (out InitPlay) {
	out.Play = player
	out.Initiative = initiative
	return
}

//Player TODO
func (i *InitPlay) Player() *Player {
	return &i.Play
}

//SetInitiative TODO
func (i *InitPlay) SetInitiative(init int) {
	i.Initiative = init
}

//InitPlays TODO
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
