package main

import (
	"os"
	"strings"
	"unsafe"

	"github.com/zurek87/go-gtk3/gdk3"
	"github.com/zurek87/go-gtk3/glib"
	"github.com/zurek87/go-gtk3/gtk3"
)

func main() {
	gtk3.Init(&os.Args)
	window := gtk3.NewWindow(gtk3.WINDOW_TOPLEVEL)
	window.SetTitle("GTK DND")
	window.Connect("destroy", gtk3.MainQuit)

	vbox := gtk3.NewVBox(true, 0)
	vbox.SetBorderWidth(5)

	targets := []gtk3.TargetEntry{
		{"text/uri-list", 0, 0},
		{"STRING", 0, 1},
		{"text/plain", 0, 2},
	}
	dest := gtk3.NewLabel("drop me file")
	dest.DragDestSet(
		gtk3.DEST_DEFAULT_MOTION|
			gtk3.DEST_DEFAULT_HIGHLIGHT|
			gtk3.DEST_DEFAULT_DROP,
		targets,
		gdk3.ACTION_COPY)
	dest.DragDestAddUriTargets()
	dest.Connect("drag-data-received", func(ctx *glib.CallbackContext) {
		sdata := gtk3.NewSelectionDataFromNative(unsafe.Pointer(ctx.Args(3)))
		if sdata != nil {
			a := (*[2000]uint8)(sdata.GetData())
			files := strings.Split(string(a[0:sdata.GetLength()-1]), "\n")
			for i := range files {
				filename, _, _ := glib.FilenameFromUri(files[i])
				files[i] = filename
			}
			dialog := gtk3.NewMessageDialog(
				window,
				gtk3.DIALOG_MODAL,
				gtk3.MESSAGE_INFO,
				gtk3.BUTTONS_OK,
				strings.Join(files, "\n"))
			dialog.SetTitle("D&D")
			dialog.Response(func() {
				dialog.Destroy()
			})
			dialog.Run()
		}
	})
	vbox.Add(dest)

	window.Add(vbox)

	window.SetSizeRequest(300, 100)
	window.ShowAll()
	gtk3.Main()
}
