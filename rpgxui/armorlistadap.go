package rpgxui

import (
	"encoding/gob"
	"os"

	"github.com/CalebQ42/go-rpg"
	"github.com/nelsam/gxui"
	"github.com/nelsam/gxui/math"
)

//ArmorListAdapter TODO
type ArmorListAdapter struct {
	gxui.AdapterBase
	arms []rpg.Armor
}

//AddArmor TODO
func (a *ArmorListAdapter) AddArmor(arm rpg.Armor) {
	a.arms = append(a.arms, arm)
	a.DataChanged(false)
}

//Armor TODO
func (a *ArmorListAdapter) Armor(index int) *rpg.Armor {
	return &a.arms[index]
}

//Remove TODO
func (a *ArmorListAdapter) Remove(index int) {
	a.arms = append(a.arms[:index], a.arms[index+1:]...)
	a.DataChanged(false)
}

//Save TODO
func (a *ArmorListAdapter) Save(filename string) {
	os.Remove(filename)
	fi, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	d := gob.NewEncoder(fi)
	err = d.Encode(a.arms)
	if err != nil {
		panic(err)
	}
	fi.Close()
}

//Load TODO
func (a *ArmorListAdapter) Load(filename string) {
	fi, err := os.Open(filename)
	if err != nil {
		return
	}
	e := gob.NewDecoder(fi)
	err = e.Decode(&a.arms)
	if err != nil {
		panic(err)
	}
	fi.Close()
	a.DataChanged(false)
}

//Count TODO
func (a *ArmorListAdapter) Count() int {
	return len(a.arms)
}

//ItemAt TODO
func (a *ArmorListAdapter) ItemAt(index int) gxui.AdapterItem {
	return a.arms[index].Name
}

//ItemIndex TODO
func (a *ArmorListAdapter) ItemIndex(item gxui.AdapterItem) int {
	nm, ok := item.(string)
	if ok {
		for i, v := range a.arms {
			if v.Name == nm {
				return i
			}
		}
	}
	return -1
}

//Create TODO
func (a *ArmorListAdapter) Create(th gxui.Theme, index int) gxui.Control {
	lbl := th.CreateLabel()
	lbl.SetText(a.arms[index].Name)
	return lbl
}

//Size TODO
func (a *ArmorListAdapter) Size(gxui.Theme) math.Size {
	return math.Size{W: math.MaxSize.W, H: 20}
}
