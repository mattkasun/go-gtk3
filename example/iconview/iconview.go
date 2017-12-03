package main

import (
	"os"
	"unsafe"

	"github.com/zurek87/go-gtk3/gdkpixbuf"
	"github.com/zurek87/go-gtk3/glib"
	"github.com/zurek87/go-gtk3/gtk3"
)

func main() {
	gtk3.Init(&os.Args)
	window := gtk3.NewWindow(gtk3.WINDOW_TOPLEVEL)
	window.SetTitle("GTK Icon View")
	window.Connect("destroy", gtk3.MainQuit)

	swin := gtk3.NewScrolledWindow(nil, nil)

	store := gtk3.NewListStore(gdkpixbuf.GetType(), glib.G_TYPE_STRING)
	iconview := gtk3.NewIconViewWithModel(store)
	iconview.SetPixbufColumn(0)
	iconview.SetTextColumn(1)
	swin.Add(iconview)

	gtk3.StockListIDs().ForEach(func(d unsafe.Pointer, v interface{}) {
		id := glib.GPtrToString(d)
		var iter gtk3.TreeIter
		store.Append(&iter)
		store.Set(&iter,
			0, gtk3.NewImage().RenderIcon(id, gtk3.ICON_SIZE_SMALL_TOOLBAR, "").GPixbuf,
			1, id)
	})

	window.Add(swin)
	window.SetSizeRequest(500, 200)
	window.ShowAll()

	gtk3.Main()
}
