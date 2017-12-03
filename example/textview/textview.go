package main

import (
	"fmt"
	"github.com/zurek87/go-gtk3/gdk3"
	"github.com/zurek87/go-gtk3/glib"
	"github.com/zurek87/go-gtk3/gtk3"
	"unsafe"
)

func main() {
	gtk3.Init(nil)

	window := gtk3.NewWindow(gtk3.WINDOW_TOPLEVEL)
	window.SetPosition(gtk3.WIN_POS_CENTER)
	window.SetTitle("GTK Go!")
	window.SetIconName("textview")
	window.Connect("destroy", gtk3.MainQuit)

	textview := gtk3.NewTextView()
	textview.SetEditable(true)
	textview.SetCursorVisible(true)
	var iter gtk3.TextIter
	buffer := textview.GetBuffer()

	buffer.GetStartIter(&iter)
	buffer.Insert(&iter, "Hello ")

	tag := buffer.CreateTag("bold", map[string]string{"background": "#FF0000", "weight": "700"})
	buffer.InsertWithTag(&iter, "Google!", tag)

	u := "http://www.google.com"
	tag.SetData("tag-name", unsafe.Pointer(&u))
	textview.Connect("event-after", func(ctx *glib.CallbackContext) {
		arg := ctx.Args(0)
		if ev := *(**gdk3.EventAny)(unsafe.Pointer(&arg)); ev.Type != gdk3.BUTTON_RELEASE {
			return
		}
		ev := *(**gdk3.EventButton)(unsafe.Pointer(&arg))
		var iter gtk3.TextIter
		textview.GetIterAtLocation(&iter, int(ev.X), int(ev.Y))
		tags := iter.GetTags()
		for n := uint(0); n < tags.Length(); n++ {
			vv := tags.NthData(n)
			tag := gtk3.NewTextTagFromPointer(vv)
			u := *(*string)(tag.GetData("tag-name"))
			fmt.Println(u)
		}
	})

	window.Add(textview)
	window.SetSizeRequest(600, 600)
	window.ShowAll()
	gtk3.Main()
}
