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

//PlayerListAdapter TODO
type PlayerListAdapter struct {
	gxui.AdapterBase
	plays rpg.InitPlays
}

//CreatePlayerListAdapter TODO
func CreatePlayerListAdapter(plays rpg.InitPlays) (out PlayerListAdapter) {
	out.plays = plays
	out.order()
	return
}

//AddPlayer TODO
func (p *PlayerListAdapter) AddPlayer(play rpg.Player, initiative int) {
	p.AddInitPlayer(rpg.CreateInitPlay(play, initiative))
	p.order()
}

//AddInitPlayer TODO
func (p *PlayerListAdapter) AddInitPlayer(play rpg.InitPlay) {
	p.plays = append(p.plays, play)
	p.order()
}

//SetInitiative TODO
func (p *PlayerListAdapter) SetInitiative(index, init int) {
	p.plays[index].SetInitiative(init)
	p.order()
}

//GetInitiative TODO
func (p *PlayerListAdapter) GetInitiative(index int) int {
	return p.plays[index].Initiative
}

//MoveUp TODO
func (p *PlayerListAdapter) MoveUp(index int) {
	if index == 0 {
		return
	}
	p.plays.Swap(index, index-1)
	p.order()
}

//MoveDown TODO
func (p *PlayerListAdapter) MoveDown(index int) {
	if index == p.plays.Len()-1 {
		return
	}
	p.plays.Swap(index, index+1)
	p.order()
}

//Remove TODO
func (p *PlayerListAdapter) Remove(index int) {
	p.plays = append(p.plays[:index], p.plays[index+1:]...)
	p.DataChanged(false)
}

//Load TODO
func (p *PlayerListAdapter) Load(filename string) {
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
func (p *PlayerListAdapter) Save(filename string) {
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

func (p *PlayerListAdapter) order() {
	sort.Sort((*p).plays)
	p.DataChanged(false)
}

//Count TODO
func (p *PlayerListAdapter) Count() int {
	return len(p.plays)
}

//ItemAt TODO
func (p *PlayerListAdapter) ItemAt(index int) gxui.AdapterItem {
	return p.plays[index].Player().Name
}

//ItemIndex TODO
func (p *PlayerListAdapter) ItemIndex(item gxui.AdapterItem) int {
	for i, v := range p.plays {
		if v.Player().Name == item {
			return i
		}
	}
	return -1
}

//Create TODO
func (p *PlayerListAdapter) Create(th gxui.Theme, index int) gxui.Control {
	lbl := th.CreateLabel()
	lbl.SetText(p.plays[index].Player().Name)
	return lbl
}

//Size TODO
func (p *PlayerListAdapter) Size(gxui.Theme) math.Size {
	return math.Size{W: math.MaxSize.W, H: 20}
}
