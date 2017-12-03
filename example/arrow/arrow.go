package main

import (
	"github.com/zurek87/go-gtk3/gtk3"
	"os"
)

func createArrowButton(at gtk3.ArrowType, st gtk3.ShadowType) *gtk3.Button {
	b := gtk3.NewButton()
	a := gtk3.NewArrow(at, st)

	b.Add(a)

	b.Show()
	a.Show()

	return b
}

func main() {
	gtk3.Init(&os.Args)

	window := gtk3.NewWindow(gtk3.WINDOW_TOPLEVEL)
	window.SetTitle("Arrow Buttons")
	window.Connect("destroy", gtk3.MainQuit)

	box := gtk3.NewHBox(false, 0)
	box.Show()
	window.Add(box)

	up := createArrowButton(gtk3.ARROW_UP, gtk3.SHADOW_IN)
	down := createArrowButton(gtk3.ARROW_DOWN, gtk3.SHADOW_OUT)
	left := createArrowButton(gtk3.ARROW_LEFT, gtk3.SHADOW_ETCHED_IN)
	right := createArrowButton(gtk3.ARROW_RIGHT, gtk3.SHADOW_ETCHED_OUT)

	box.PackStart(up, false, false, 3)
	box.PackStart(down, false, false, 3)
	box.PackStart(left, false, false, 3)
	box.PackStart(right, false, false, 3)

	up.Clicked(func() { println("↑") })
	down.Clicked(func() { println("↓") })
	left.Clicked(func() { println("←") })
	right.Clicked(func() { println("→") })

	window.Show()
	gtk3.Main()
}
