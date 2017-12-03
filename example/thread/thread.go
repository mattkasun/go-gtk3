package main

import (
	"runtime"
	"strconv"
	"time"

	"github.com/zurek87/go-gtk3/gdk3"
	"github.com/zurek87/go-gtk3/glib"
	"github.com/zurek87/go-gtk3/gtk3"
)

func main() {
	runtime.GOMAXPROCS(10)
	glib.ThreadInit(nil)
	gdk3.ThreadsInit()
	gdk3.ThreadsEnter()
	gtk3.Init(nil)
	window := gtk3.NewWindow(gtk3.WINDOW_TOPLEVEL)
	window.Connect("destroy", gtk3.MainQuit)

	vbox := gtk3.NewVBox(false, 1)

	label1 := gtk3.NewLabel("")
	vbox.Add(label1)
	label2 := gtk3.NewLabel("")
	vbox.Add(label2)

	window.Add(vbox)

	window.SetSizeRequest(100, 100)
	window.ShowAll()
	time.Sleep(1000 * 1000 * 100)
	go (func() {
		for i := 0; i < 300000; i++ {
			gdk3.ThreadsEnter()
			label1.SetLabel(strconv.Itoa(i))
			gdk3.ThreadsLeave()
		}
		gtk3.MainQuit()
	})()
	go (func() {
		for i := 300000; i >= 0; i-- {
			gdk3.ThreadsEnter()
			label2.SetLabel(strconv.Itoa(i))
			gdk3.ThreadsLeave()
		}
		gtk3.MainQuit()
	})()
	gtk3.Main()
}
