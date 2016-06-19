package rpgxui

import (
	"encoding/gob"
	"os"

	"github.com/CalebQ42/go-rpg"
	"github.com/nelsam/gxui"
	"github.com/nelsam/gxui/math"
)

//PlayerListAdapter TODO
type PlayerListAdapter struct {
	gxui.AdapterBase
	plays []rpg.Player
}

//AddPlayer TODO
func (p *PlayerListAdapter) AddPlayer(pl rpg.Player) {
	p.plays = append(p.plays, pl)
}

//Player TODO
func (p *PlayerListAdapter) Player(index int) *rpg.Player {
	return &p.plays[index]
}

//Remove TODO
func (p *PlayerListAdapter) Remove(index int) {
	p.plays = append(p.plays[:index], p.plays[index+1:]...)
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

//Load TODO
func (p *PlayerListAdapter) Load(filename string) {
	fi, err := os.Open(filename)
	if err != nil {
		return
	}
	d := gob.NewDecoder(fi)
	err = d.Decode(&p.plays)
	if err != nil {
		panic(err)
	}
	fi.Close()
}

//Count TODO
func (p *PlayerListAdapter) Count() int {
	return len(p.plays)
}

//ItemAt TODO
func (p *PlayerListAdapter) ItemAt(index int) gxui.AdapterItem {
	return p.plays[index].Name
}

//ItemIndex TODO
func (p *PlayerListAdapter) ItemIndex(item gxui.AdapterItem) int {
	name, ok := item.(string)
	if ok {
		for i, v := range p.plays {
			if v.Name == name {
				return i
			}
		}
	}
	return -1
}

//Create TODO
func (p *PlayerListAdapter) Create(th gxui.Theme, index int) gxui.Control {
	lbl := th.CreateLabel()
	lbl.SetText(p.plays[index].Name)
	return lbl
}

//Size TODO
func (p *PlayerListAdapter) Size(gxui.Theme) math.Size {
	return math.Size{H: 20, W: math.MaxSize.W}
}
