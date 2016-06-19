package rpgxui

import (
	"encoding/gob"
	"fmt"
	"os"
	"sort"

	"github.com/CalebQ42/go-rpg"
	"github.com/nelsam/gxui"
	"github.com/nelsam/gxui/math"
)

//InitiativeListAdapter TODO
type InitiativeListAdapter struct {
	gxui.AdapterBase
	plays rpg.InitPlays
}

//CreatePlayerListAdapter TODO
func CreatePlayerListAdapter(plays rpg.InitPlays) (out InitiativeListAdapter) {
	out.plays = plays
	out.order()
	return
}

//InitPlay TODO
func (p *InitiativeListAdapter) InitPlay(index int) *rpg.InitPlay {
	return &p.plays[index]
}

//AddPlayer TODO
func (p *InitiativeListAdapter) AddPlayer(play rpg.Player, initiative int) {
	p.AddInitPlayer(rpg.CreateInitPlay(play, initiative))
	p.order()
}

//AddInitPlayer TODO
func (p *InitiativeListAdapter) AddInitPlayer(play rpg.InitPlay) {
	p.plays = append(p.plays, play)
	p.order()
}

//SetInitiative TODO
func (p *InitiativeListAdapter) SetInitiative(index, init int) {
	p.plays[index].Initiative = init
	p.order()
}

//GetInitiative TODO
func (p *InitiativeListAdapter) GetInitiative(index int) int {
	return p.plays[index].Initiative
}

//MoveUp TODO
func (p *InitiativeListAdapter) MoveUp(index int) {
	if index == 0 {
		return
	}
	p.plays[index].Initiative = p.plays[index-1].Initiative
	p.plays.Swap(index, index-1)
	p.order()
}

//MoveDown TODO
func (p *InitiativeListAdapter) MoveDown(index int) {
	if index == p.plays.Len()-1 {
		return
	}
	p.plays[index].Initiative = p.plays[index+1].Initiative
	p.plays.Swap(index, index+1)
	p.order()
}

//Remove TODO
func (p *InitiativeListAdapter) Remove(index int) {
	p.plays = append(p.plays[:index], p.plays[index+1:]...)
	p.DataChanged(false)
}

//Load TODO
func (p *InitiativeListAdapter) Load(filename string) {
	fi, err := os.Open(filename)
	if err != nil {
		return
	}
	de := gob.NewDecoder(fi)
	err = de.Decode(&p.plays)
	if err != nil {
		fmt.Println(err)
	}
	fi.Close()
	p.order()
}

//Save TODO
func (p *InitiativeListAdapter) Save(filename string) {
	os.Remove(filename)
	fi, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	e := gob.NewEncoder(fi)
	err = e.Encode(p.plays)
	if err != nil {
		panic(err)
	}
	fi.Close()
}

func (p *InitiativeListAdapter) order() {
	sort.Sort((*p).plays)
	p.DataChanged(true)
}

//Count TODO
func (p *InitiativeListAdapter) Count() int {
	return len(p.plays)
}

//ItemAt TODO
func (p *InitiativeListAdapter) ItemAt(index int) gxui.AdapterItem {
	return p.plays[index].Player.Name
}

//ItemIndex TODO
func (p *InitiativeListAdapter) ItemIndex(item gxui.AdapterItem) int {
	for i, v := range p.plays {
		if v.Player.Name == item {
			return i
		}
	}
	return -1
}

//Create TODO
func (p *InitiativeListAdapter) Create(th gxui.Theme, index int) gxui.Control {
	lbl := th.CreateLabel()
	lbl.SetText(p.plays[index].Player.Name)
	return lbl
}

//Size TODO
func (p *InitiativeListAdapter) Size(gxui.Theme) math.Size {
	return math.Size{W: math.MaxSize.W, H: 20}
}
