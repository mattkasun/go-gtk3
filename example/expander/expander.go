package main

import (
	"github.com/zurek87/go-gtk3/gtk3"
	"os"
)

func main() {
	gtk3.Init(&os.Args)
	window := gtk3.NewWindow(gtk3.WINDOW_TOPLEVEL)
	window.SetTitle("We love Expander")
	window.Connect("destroy", gtk3.MainQuit)

	vbox := gtk3.NewVBox(true, 0)
	vbox.SetBorderWidth(5)
	expander := gtk3.NewExpander("dan the ...")
	expander.Add(gtk3.NewLabel("404 contents not found"))
	vbox.PackStart(expander, false, false, 0)

	window.Add(vbox)
	window.ShowAll()

	gtk3.Main()
}
