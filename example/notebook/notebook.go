package main

import (
	"github.com/zurek87/go-gtk3/gtk3"
	"os"
	"strconv"
)

func main() {
	gtk3.Init(&os.Args)
	window := gtk3.NewWindow(gtk3.WINDOW_TOPLEVEL)
	window.SetTitle("GTK Notebook")
	window.Connect("destroy", gtk3.MainQuit)

	notebook := gtk3.NewNotebook()
	for n := 1; n <= 10; n++ {
		page := gtk3.NewFrame("demo" + strconv.Itoa(n))
		notebook.AppendPage(page, gtk3.NewLabel("demo"+strconv.Itoa(n)))

		vbox := gtk3.NewHBox(false, 1)

		prev := gtk3.NewButtonWithLabel("go prev")
		prev.Clicked(func() {
			notebook.PrevPage()
		})
		vbox.Add(prev)

		next := gtk3.NewButtonWithLabel("go next")
		next.Clicked(func() {
			notebook.NextPage()
		})
		vbox.Add(next)

		page.Add(vbox)
	}

	window.Add(notebook)
	window.SetSizeRequest(400, 200)
	window.ShowAll()

	gtk3.Main()
}
