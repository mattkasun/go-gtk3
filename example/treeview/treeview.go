package main

import (
	"github.com/zurek87/go-gtk3/gdkpixbuf"
	"github.com/zurek87/go-gtk3/glib"
	"github.com/zurek87/go-gtk3/gtk3"
	"os"
	"strconv"
)

func main() {
	gtk3.Init(&os.Args)
	window := gtk3.NewWindow(gtk3.WINDOW_TOPLEVEL)
	window.SetTitle("GTK Folder View")
	window.Connect("destroy", gtk3.MainQuit)

	swin := gtk3.NewScrolledWindow(nil, nil)

	store := gtk3.NewTreeStore(gdkpixbuf.GetType(), glib.G_TYPE_STRING)
	treeview := gtk3.NewTreeView()
	swin.Add(treeview)

	treeview.SetModel(store.ToTreeModel())
	treeview.AppendColumn(gtk3.NewTreeViewColumnWithAttributes("pixbuf", gtk3.NewCellRendererPixbuf(), "pixbuf", 0))
	treeview.AppendColumn(gtk3.NewTreeViewColumnWithAttributes("text", gtk3.NewCellRendererText(), "text", 1))

	for n := 1; n <= 10; n++ {
		var iter1, iter2, iter3 gtk3.TreeIter
		store.Append(&iter1, nil)
		store.Set(&iter1, gtk3.NewImage().RenderIcon(gtk3.STOCK_DIRECTORY, gtk3.ICON_SIZE_SMALL_TOOLBAR, "").GPixbuf, "Folder"+strconv.Itoa(n))
		store.Append(&iter2, &iter1)
		store.Set(&iter2, gtk3.NewImage().RenderIcon(gtk3.STOCK_DIRECTORY, gtk3.ICON_SIZE_SMALL_TOOLBAR, "").GPixbuf, "SubFolder"+strconv.Itoa(n))
		store.Append(&iter3, &iter2)
		store.Set(&iter3, gtk3.NewImage().RenderIcon(gtk3.STOCK_FILE, gtk3.ICON_SIZE_SMALL_TOOLBAR, "").GPixbuf, "File"+strconv.Itoa(n))
	}

	treeview.Connect("row_activated", func() {
		var path *gtk3.TreePath
		var column *gtk3.TreeViewColumn
		treeview.GetCursor(&path, &column)
		mes := "TreePath is: " + path.String()
		dialog := gtk3.NewMessageDialog(
			treeview.GetTopLevelAsWindow(),
			gtk3.DIALOG_MODAL,
			gtk3.MESSAGE_INFO,
			gtk3.BUTTONS_OK,
			mes)
		dialog.SetTitle("TreePath")
		dialog.Response(func() {
			dialog.Destroy()
		})
		dialog.Run()
	})

	window.Add(swin)
	window.SetSizeRequest(400, 200)
	window.ShowAll()

	gtk3.Main()
}
