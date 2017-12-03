package main

import (
	"github.com/zurek87/go-gtk3/gtk3"
	"os"
)

func main() {
	gtk3.Init(&os.Args)
	window := gtk3.NewWindow(gtk3.WINDOW_TOPLEVEL)
	window.SetTitle("Alignment")
	window.Connect("destroy", gtk3.MainQuit)

	notebook := gtk3.NewNotebook()
	window.Add(notebook)

	align := gtk3.NewAlignment(0.5, 0.5, 0.5, 0.5)
	notebook.AppendPage(align, gtk3.NewLabel("Alignment"))

	button := gtk3.NewButtonWithLabel("Hello World!")
	align.Add(button)

	fixed := gtk3.NewFixed()
	notebook.AppendPage(fixed, gtk3.NewLabel("Fixed"))

	button2 := gtk3.NewButtonWithLabel("Pulse")
	fixed.Put(button2, 30, 30)

	progress := gtk3.NewProgressBar()
	fixed.Put(progress, 30, 70)

	button.Connect("clicked", func() {
		progress.SetFraction(0.1 + 0.9*progress.GetFraction()) //easter egg
	})
	button2.Connect("clicked", func() {
		progress.Pulse()
	})

	window.ShowAll()
	window.SetSizeRequest(200, 200)

	gtk3.Main()
}
