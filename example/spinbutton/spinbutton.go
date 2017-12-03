package main

import (
	"fmt"
	"github.com/zurek87/go-gtk3/glib"
	"github.com/zurek87/go-gtk3/gtk3"
	"strconv"
)

func main() {
	gtk3.Init(nil)
	window := gtk3.NewWindow(gtk3.WINDOW_TOPLEVEL)
	window.SetPosition(gtk3.WIN_POS_CENTER)
	window.SetTitle("GTK Go!")
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		fmt.Println("got destroy!", ctx.Data().(string))
		gtk3.MainQuit()
	}, "foo")

	//--------------------------------------------------------
	// GtkHBox
	//--------------------------------------------------------
	fixed := gtk3.NewFixed()

	//--------------------------------------------------------
	// GtkSpinButton
	//--------------------------------------------------------
	spinbutton1 := gtk3.NewSpinButtonWithRange(1.0, 10.0, 1.0)
	spinbutton1.SetDigits(3)
	spinbutton1.Spin(gtk3.SPIN_STEP_FORWARD, 7.0)
	fixed.Put(spinbutton1, 40, 50)

	spinbutton1.OnValueChanged(func() {
		val := spinbutton1.GetValueAsInt()
		fval := spinbutton1.GetValue()
		fmt.Println("SpinButton changed, new value: " + strconv.Itoa(val) + " | " + strconv.FormatFloat(fval, 'f', 2, 64))
		min, max := spinbutton1.GetRange()
		fmt.Println("Range: " + strconv.FormatFloat(min, 'f', 2, 64) + " " + strconv.FormatFloat(max, 'f', 2, 64))
		fmt.Println("Digits: " + strconv.Itoa(int(spinbutton1.GetDigits())))
	})

	adjustment := gtk3.NewAdjustment(2.0, 1.0, 8.0, 2.0, 0.0, 0.0)
	spinbutton2 := gtk3.NewSpinButton(adjustment, 1.0, 1)
	spinbutton2.SetRange(0.0, 20.0)
	spinbutton2.SetValue(18.0)
	spinbutton2.SetIncrements(2.0, 4.0)
	fixed.Put(spinbutton2, 150, 50)

	//--------------------------------------------------------
	// Event
	//--------------------------------------------------------
	window.Add(fixed)
	window.SetSizeRequest(600, 600)
	window.ShowAll()
	gtk3.Main()
}
