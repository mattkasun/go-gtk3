package main

import (
	"fmt"
	"github.com/zurek87/go-gtk3/gtk3"
	"os"
)

func main() {
	gtk3.Init(&os.Args)
	window := gtk3.NewWindow(gtk3.WINDOW_TOPLEVEL)
	window.SetTitle("GTK Table")
	window.Connect("destroy", gtk3.MainQuit)

	swin := gtk3.NewScrolledWindow(nil, nil)
	swin.SetPolicy(gtk3.POLICY_AUTOMATIC, gtk3.POLICY_AUTOMATIC)

	table := gtk3.NewTable(5, 5, false)
	for y := uint(0); y < 5; y++ {
		for x := uint(0); x < 5; x++ {
			table.Attach(gtk3.NewButtonWithLabel(fmt.Sprintf("%02d:%02d", x, y)), x, x+1, y, y+1, gtk3.FILL, gtk3.FILL, 5, 5)
		}
	}
	swin.AddWithViewPort(table)

	window.Add(swin)
	window.SetDefaultSize(200, 200)
	window.ShowAll()

	gtk3.Main()
}
