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
	window.SetTitle("GTK Stock Icons")
	window.Connect("destroy", gtk3.MainQuit)

	swin := gtk3.NewScrolledWindow(nil, nil)

	store := gtk3.NewListStore(glib.G_TYPE_STRING, glib.G_TYPE_BOOL, gdkpixbuf.GetType())
	treeview := gtk3.NewTreeView()
	swin.Add(treeview)

	treeview.SetModel(store)
	treeview.AppendColumn(gtk3.NewTreeViewColumnWithAttributes("name", gtk3.NewCellRendererText(), "text", 0))
	treeview.AppendColumn(gtk3.NewTreeViewColumnWithAttributes("check", gtk3.NewCellRendererToggle(), "active", 1))
	treeview.AppendColumn(gtk3.NewTreeViewColumnWithAttributes("icon", gtk3.NewCellRendererPixbuf(), "pixbuf", 2))
	n := 0
	gtk3.StockListIDs().ForEach(func(d unsafe.Pointer, v interface{}) {
		id := glib.GPtrToString(d)
		var iter gtk3.TreeIter
		store.Append(&iter)
		store.Set(&iter,
			0, id,
			1, (n == 1),
			2, gtk3.NewImage().RenderIcon(id, gtk3.ICON_SIZE_SMALL_TOOLBAR, "").GPixbuf,
		)
		n = 1 - n
	})

	window.Add(swin)
	window.SetSizeRequest(400, 200)
	window.ShowAll()

	gtk3.Main()
}
