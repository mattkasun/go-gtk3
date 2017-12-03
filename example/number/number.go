package main

import (
	"fmt"
	"github.com/zurek87/go-gtk3/glib"
	"github.com/zurek87/go-gtk3/gtk3"
	"os"
	"strconv"
	"unsafe"
)

func main() {
	gtk3.Init(&os.Args)

	dialog := gtk3.NewDialog()
	dialog.SetTitle("number input")

	vbox := dialog.GetVBox()

	label := gtk3.NewLabel("Numnber:")
	vbox.Add(label)

	input := gtk3.NewEntry()
	input.SetEditable(true)
	vbox.Add(input)

	input.Connect("insert-text", func(ctx *glib.CallbackContext) {
		a := (*[2000]uint8)(unsafe.Pointer(ctx.Args(0)))
		p := (*int)(unsafe.Pointer(ctx.Args(2)))
		i := 0
		for a[i] != 0 {
			i++
		}
		s := string(a[0:i])
		if s == "." {
			if *p == 0 {
				input.StopEmission("insert-text")
			}
		} else {
			_, err := strconv.ParseFloat(s, 64)
			if err != nil {
				input.StopEmission("insert-text")
			}
		}
	})

	button := gtk3.NewButtonWithLabel("OK")
	button.Connect("clicked", func() {
		fmt.Println(input.GetText())
		gtk3.MainQuit()
	})
	vbox.Add(button)

	dialog.ShowAll()
	gtk3.Main()
}
