package main

import (
	"github.com/zurek87/go-gtk3/gdk3"
	"github.com/zurek87/go-gtk3/glib"
	"github.com/zurek87/go-gtk3/gtk3"
	"os"
	"unsafe"
)

type point struct {
	x int
	y int
}

func main() {
	gtk3.Init(&os.Args)
	window := gtk3.NewWindow(gtk3.WINDOW_TOPLEVEL)
	window.SetTitle("GTK DrawingArea")
	window.Connect("destroy", gtk3.MainQuit)

	vbox := gtk3.NewVBox(true, 0)
	vbox.SetBorderWidth(5)
	drawingarea := gtk3.NewDrawingArea()

	var p1, p2 point
	var gdkwin *gdk3.Window
	var pixmap *gdk3.Pixmap
	var gc *gdk3.GC
	p1.x = -1
	p1.y = -1
	colors := []string{
		"black",
		"gray",
		"blue",
		"purple",
		"red",
		"orange",
		"yellow",
		"green",
		"darkgreen",
	}

	drawingarea.Connect("configure-event", func() {
		if pixmap != nil {
			pixmap.Unref()
		}
		allocation := drawingarea.GetAllocation()
		pixmap = gdk3.NewPixmap(drawingarea.GetWindow().GetDrawable(), allocation.Width, allocation.Height, 24)
		gc = gdk3.NewGC(pixmap.GetDrawable())
		gc.SetRgbFgColor(gdk3.NewColor("white"))
		pixmap.GetDrawable().DrawRectangle(gc, true, 0, 0, -1, -1)
		gc.SetRgbFgColor(gdk3.NewColor(colors[0]))
		gc.SetRgbBgColor(gdk3.NewColor("white"))
	})

	drawingarea.Connect("motion-notify-event", func(ctx *glib.CallbackContext) {
		arg := ctx.Args(0)
		mev := *(**gdk3.EventMotion)(unsafe.Pointer(&arg))
		var mt gdk3.ModifierType
		if mev.IsHint != 0 {
			gdkwin.GetPointer(&p2.x, &p2.y, &mt)
		} else {
			p2.x, p2.y = int(mev.X), int(mev.Y)
		}
		if p1.x != -1 && p2.x != -1 && (gdk3.EventMask(mt)&gdk3.BUTTON_PRESS_MASK) != 0 {
			pixmap.GetDrawable().DrawLine(gc, p1.x, p1.y, p2.x, p2.y)
			gdkwin.Invalidate(nil, false)
		}
		colors = append(colors[1:], colors[0])
		gc.SetRgbFgColor(gdk3.NewColor(colors[0]))
		p1 = p2
	})

	drawingarea.Connect("expose-event", func() {
		if pixmap == nil {
			return
		}
		gdkwin.GetDrawable().DrawDrawable(gc, pixmap.GetDrawable(), 0, 0, 0, 0, -1, -1)
	})

	drawingarea.SetEvents(int(gdk3.POINTER_MOTION_MASK | gdk3.POINTER_MOTION_HINT_MASK | gdk3.BUTTON_PRESS_MASK))
	vbox.Add(drawingarea)

	window.Add(vbox)
	window.SetSizeRequest(400, 400)
	window.ShowAll()

	gdkwin = drawingarea.GetWindow()

	gtk3.Main()
}
