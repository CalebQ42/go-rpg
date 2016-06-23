package main

import (
	"strconv"
	"strings"

	"github.com/CalebQ42/go-rpg"
	"github.com/CalebQ42/go-rpg/rpgxui"
	"github.com/nelsam/gxui"
	"github.com/nelsam/gxui/drivers/gl"
	"github.com/nelsam/gxui/math"
	"github.com/nelsam/gxui/themes/dark"
)

var (
	bk    = gxui.CreateBrush(gxui.ColorFromHex(0xFF455A64))
	peeps = &rpgxui.InitiativeListAdapter{}
)

func main() {
	peeps.Load("players.tada")
	gl.StartDriver(uiMain)
}

func uiMain(dr gxui.Driver) {
	th := dark.CreateTheme(dr)
	win := th.CreateWindow(500, 500, "Initiative")
	win.SetBackgroundBrush(bk)
	list := th.CreateList()
	list.SetAdapter(peeps)
	up := th.CreateButton()
	up.SetText("Move Up")
	up.OnClick(func(gxui.MouseEvent) {
		it := peeps.ItemIndex(list.Selected())
		if it != -1 {
			peeps.MoveUp(it)
		}
	})
	dw := th.CreateButton()
	dw.SetText("Move Down")
	dw.OnClick(func(gxui.MouseEvent) {
		it := peeps.ItemIndex(list.Selected())
		if it != -1 {
			peeps.MoveDown(it)
		}
	})
	rm := th.CreateButton()
	rm.SetText("Remove Player")
	rm.OnClick(func(gxui.MouseEvent) {
		it := peeps.ItemIndex(list.Selected())
		if it != -1 {
			peeps.Remove(it)
		}
	})
	ad := th.CreateButton()
	ad.SetText("Add Player")
	ad.OnClick(func(gxui.MouseEvent) {
		dr.Call(func() {
			adplay(th, win)
		})
	})
	st := th.CreateButton()
	st.SetText("Set Initiative")
	st.OnClick(func(gxui.MouseEvent) {
		it := peeps.ItemIndex(list.Selected())
		if it != -1 {
			dr.Call(func() {
				stini(th, win, it)
			})
		}
	})
	butLay := th.CreateLinearLayout()
	butLay.SetDirection(gxui.TopToBottom)
	butLay.AddChild(up)
	butLay.AddChild(dw)
	butLay.AddChild(rm)
	butLay.AddChild(ad)
	butLay.AddChild(st)
	spl := th.CreateSplitterLayout()
	spl.SetOrientation(gxui.Horizontal)
	spl.AddChild(list)
	spl.AddChild(butLay)
	win.AddChild(spl)
	win.OnClose(func() {
		peeps.Save("players.tada")
		dr.Terminate()
	})
}

func adplay(th gxui.Theme, win gxui.Window) {
	win.Hide()
	w := th.CreateWindow(300, 80, "Adding Player")
	w.SetBackgroundBrush(bk)
	lbl := th.CreateLabel()
	lbl.SetText("Player Name:")
	txt := th.CreateTextBox()
	txt.SetDesiredWidth(math.MaxSize.W)
	but := th.CreateButton()
	but.SetText("Add Player")
	but.OnClick(func(gxui.MouseEvent) {
		w.Close()
	})
	lyt := th.CreateLinearLayout()
	lyt.SetDirection(gxui.TopToBottom)
	lyt.AddChild(lbl)
	lyt.AddChild(txt)
	lyt.AddChild(but)
	w.AddChild(lyt)
	w.Redraw()
	w.OnClose(func() {
		if strings.TrimSpace(txt.Text()) != "" {
			if peeps.ItemIndex(txt.Text()) != -1 {
				for num := 0; num < 100; num++ {
					if peeps.ItemIndex(txt.Text()+strconv.Itoa(num)) == -1 {
						play := rpg.CreatePlayer(txt.Text()+strconv.Itoa(num), rpg.Stats{})
						peeps.AddPlayer(play, 0)
						break
					}
				}
			} else {
				play := rpg.CreatePlayer(txt.Text(), rpg.Stats{})
				peeps.AddPlayer(play, 0)
			}
		}
		win.Show()
	})
}

func stini(th gxui.Theme, win gxui.Window, index int) {
	win.Hide()
	w := th.CreateWindow(300, 80, peeps.ItemAt(index).(string))
	w.SetBackgroundBrush(bk)
	lbl := th.CreateLabel()
	lbl.SetText("Initiative:")
	txt := th.CreateTextBox()
	txt.SetDesiredWidth(math.MaxSize.W)
	txt.SetText(strconv.Itoa(peeps.GetInitiative(index)))
	but := th.CreateButton()
	but.SetText("Set Initiative")
	but.OnClick(func(gxui.MouseEvent) {
		w.Close()
	})
	lyt := th.CreateLinearLayout()
	lyt.SetDirection(gxui.TopToBottom)
	lyt.AddChild(lbl)
	lyt.AddChild(txt)
	lyt.AddChild(but)
	w.AddChild(lyt)
	w.OnClose(func() {
		if strings.TrimSpace(txt.Text()) != "" {
			num, err := strconv.Atoi(strings.TrimSpace(txt.Text()))
			if err == nil {
				peeps.SetInitiative(index, num)
			}
		}
		win.Show()
	})
}
