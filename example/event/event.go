package main

import (
	"fmt"
	"github.com/zurek87/go-gtk3/gdk3"
	"github.com/zurek87/go-gtk3/glib"
	"github.com/zurek87/go-gtk3/gtk3"
	"os"
	"unsafe"
)

func main() {
	gtk3.Init(&os.Args)
	window := gtk3.NewWindow(gtk3.WINDOW_TOPLEVEL)
	window.SetTitle("GTK Events")
	window.Connect("destroy", gtk3.MainQuit)

	event := make(chan interface{})

	window.Connect("key-press-event", func(ctx *glib.CallbackContext) {
		arg := ctx.Args(0)
		event <- *(**gdk3.EventKey)(unsafe.Pointer(&arg))
	})
	window.Connect("motion-notify-event", func(ctx *glib.CallbackContext) {
		arg := ctx.Args(0)
		event <- *(**gdk3.EventMotion)(unsafe.Pointer(&arg))
	})

	go func() {
		for {
			e := <-event
			switch ev := e.(type) {
			case *gdk3.EventKey:
				fmt.Println("key-press-event:", ev.Keyval)
				break
			case *gdk3.EventMotion:
				fmt.Println("motion-notify-event:", int(ev.X), int(ev.Y))
				break
			}
		}
	}()

	window.SetEvents(int(gdk3.POINTER_MOTION_MASK | gdk3.POINTER_MOTION_HINT_MASK | gdk3.BUTTON_PRESS_MASK))
	window.SetSizeRequest(400, 400)
	window.ShowAll()

	gtk3.Main()
}
