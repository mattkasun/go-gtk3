package main

import (
	"fmt"
	"github.com/zurek87/go-gtk3/glib"
	"github.com/zurek87/go-gtk3/gtk3"
)

func main() {
	gtk3.Init(nil)
	window := gtk3.NewWindow(gtk3.WINDOW_TOPLEVEL)
	window.SetPosition(gtk3.WIN_POS_CENTER)
	window.SetTitle("GTK Go!")
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		gtk3.MainQuit()
	}, "")

	box := gtk3.NewHPaned()

	palette := gtk3.NewToolPalette()
	group := gtk3.NewToolItemGroup("Tools")
	b := gtk3.NewToolButtonFromStock(gtk3.STOCK_NEW)
	b.OnClicked(func() { fmt.Println("You clicked new!") })
	group.Insert(b, -1)
	b = gtk3.NewToolButtonFromStock(gtk3.STOCK_CLOSE)
	group.Insert(b, -1)
	b = gtk3.NewToolButtonFromStock(gtk3.STOCK_REDO)
	group.Insert(b, -1)
	b = gtk3.NewToolButtonFromStock(gtk3.STOCK_REFRESH)
	group.Insert(b, -1)
	b = gtk3.NewToolButtonFromStock(gtk3.STOCK_QUIT)
	group.Insert(b, -1)
	b = gtk3.NewToolButtonFromStock(gtk3.STOCK_YES)
	group.Insert(b, -1)
	b = gtk3.NewToolButtonFromStock(gtk3.STOCK_NO)
	group.Insert(b, -1)
	b = gtk3.NewToolButtonFromStock(gtk3.STOCK_PRINT)
	group.Insert(b, -1)
	b = gtk3.NewToolButtonFromStock(gtk3.STOCK_NETWORK)
	group.Insert(b, -1)
	b = gtk3.NewToolButtonFromStock(gtk3.STOCK_INFO)
	group.Insert(b, -1)
	b = gtk3.NewToolButtonFromStock(gtk3.STOCK_HOME)
	group.Insert(b, -1)
	b = gtk3.NewToolButtonFromStock(gtk3.STOCK_INDEX)
	group.Insert(b, -1)
	b = gtk3.NewToolButtonFromStock(gtk3.STOCK_FIND)
	group.Insert(b, -1)
	b = gtk3.NewToolButtonFromStock(gtk3.STOCK_FILE)
	group.Insert(b, -1)
	b = gtk3.NewToolButtonFromStock(gtk3.STOCK_EXECUTE)
	group.Insert(b, -1)
	palette.Add(group)

	bcopy := gtk3.NewToolButtonFromStock(gtk3.STOCK_COPY)
	bcut := gtk3.NewToolButtonFromStock(gtk3.STOCK_CUT)
	bdelete := gtk3.NewToolButtonFromStock(gtk3.STOCK_DELETE)

	group = gtk3.NewToolItemGroup("Stuff")
	group.Insert(bcopy, -1)
	group.Insert(bcut, -1)
	group.Insert(bdelete, -1)
	palette.Add(group)

	frame := gtk3.NewVBox(false, 1)
	align := gtk3.NewAlignment(0, 0, 0, 0)
	image := gtk3.NewImageFromFile("./turkey.jpg")
	align.Add(image)
	frame.Add(align)

	box.Pack1(palette, true, false)
	box.Pack2(frame, false, false)

	window.Add(box)
	window.SetSizeRequest(600, 600)
	window.ShowAll()
	gtk3.Main()
}
