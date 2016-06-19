package rpgxui

import (
	"encoding/gob"
	"os"

	"github.com/CalebQ42/go-rpg"
	"github.com/nelsam/gxui"
	"github.com/nelsam/gxui/math"
)

//ArmorListAdapter is a list adapter for an armor array
type ArmorListAdapter struct {
	gxui.AdapterBase
	arms []rpg.Armor
}

//AddArmor adds an armor onto the array
func (a *ArmorListAdapter) AddArmor(arm rpg.Armor) {
	a.arms = append(a.arms, arm)
	a.DataChanged(false)
}

//Armor returns a pointer to the armor at the index
func (a *ArmorListAdapter) Armor(index int) *rpg.Armor {
	return &a.arms[index]
}

//Remove removes armor from the array
func (a *ArmorListAdapter) Remove(index int) {
	a.arms = append(a.arms[:index], a.arms[index+1:]...)
	a.DataChanged(false)
}

//Save saves the array to the filename using gob. First deletes the file (if it exists)
//Then saves to the file
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

//Load loads the array from the specified filename using gob. If there is an error
//while opening the file then it simply returns
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

//Count returns the length of the armor array
func (a *ArmorListAdapter) Count() int {
	return len(a.arms)
}

//ItemAt returns the name of the armor a the specified index
func (a *ArmorListAdapter) ItemAt(index int) gxui.AdapterItem {
	return a.arms[index].Name
}

//ItemIndex returns the index of the armor with the specified name
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

//Create creates a label with the armor at the index's name
func (a *ArmorListAdapter) Create(th gxui.Theme, index int) gxui.Control {
	lbl := th.CreateLabel()
	lbl.SetText(a.arms[index].Name)
	return lbl
}

//Size figure it out
func (a *ArmorListAdapter) Size(gxui.Theme) math.Size {
	return math.Size{W: math.MaxSize.W, H: 20}
}
