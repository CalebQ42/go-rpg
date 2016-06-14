package rpg

//InitPlay TODO
type InitPlay struct {
	play       *Player
	initiative int
}

//CreateInitPlay TODO
func CreateInitPlay(player *Player, initiative int) (out InitPlay) {
	out.play = player
	out.initiative = initiative
	return
}

//Player TODO
func (i *InitPlay) Player() *Player {
	return i.play
}

//Initiative TODO
func (i InitPlay) Initiative() int {
	return i.initiative
}

//SetInitiative TODO
func (i *InitPlay) SetInitiative(init int) {
	i.initiative = init
}

//InitPlays TODO
type InitPlays []InitPlay

func (ip InitPlays) Len() int {
	return len(ip)
}

func (ip InitPlays) Less(i, j int) bool {
	return ip[i].initiative > ip[j].initiative
}

func (ip InitPlays) Swap(i, j int) {
	ip[i], ip[j] = ip[j], ip[i]
}
