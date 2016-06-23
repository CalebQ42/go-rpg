package rpgxui

import (
	"encoding/gob"
	"os"

	"github.com/CalebQ42/go-rpg"
	"github.com/nelsam/gxui"
	"github.com/nelsam/gxui/math"
)

//WeaponListAdapter TODO
type WeaponListAdapter struct {
	gxui.AdapterBase
	weaps []rpg.Weapon
}

//AddWeapon TODO
func (w *WeaponListAdapter) AddWeapon(wep rpg.Weapon) {
	w.weaps = append(w.weaps, wep)
	w.DataChanged(false)
}

//Weapon TODO
func (w *WeaponListAdapter) Weapon(index int) *rpg.Weapon {
	return &w.weaps[index]
}

//Remove TODO
func (w *WeaponListAdapter) Remove(index int) {
	w.weaps = append(w.weaps[:index], w.weaps[index+1:]...)
	w.DataChanged(false)
}

//Save TODO
func (w *WeaponListAdapter) Save(filename string) {
	os.Remove(filename)
	fi, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	e := gob.NewEncoder(fi)
	err = e.Encode(w.weaps)
	if err != nil {
		panic(err)
	}
	fi.Close()
}

//Load TODO
func (w *WeaponListAdapter) Load(filename string) {
	fi, err := os.Open(filename)
	if err != nil {
		return
	}
	d := gob.NewDecoder(fi)
	err = d.Decode(&w.weaps)
	if err != nil {
		panic(err)
	}
	fi.Close()
	w.DataChanged(false)
}

//Count TODO
func (w *WeaponListAdapter) Count() int {
	return len(w.weaps)
}

//ItemAt TODO
func (w *WeaponListAdapter) ItemAt(index int) gxui.AdapterItem {
	return w.weaps[index].Name
}

//ItemIndex TODO
func (w *WeaponListAdapter) ItemIndex(item gxui.AdapterItem) int {
	nm, ok := item.(string)
	if ok {
		for i, v := range w.weaps {
			if v.Name == nm {
				return i
			}
		}
	}
	return -1
}

//Create TODO
func (w *WeaponListAdapter) Create(th gxui.Theme, index int) gxui.Control {
	lbl := th.CreateLabel()
	lbl.SetText(w.weaps[index].Name)
	return lbl
}

//Size TODO
func (w *WeaponListAdapter) Size(gxui.Theme) math.Size {
	return math.Size{H: 20, W: math.MaxSize.W}
}
