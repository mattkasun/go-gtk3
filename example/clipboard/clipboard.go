package main

import (
	"fmt"
	"github.com/zurek87/go-gtk3/gdk3"
	"github.com/zurek87/go-gtk3/gtk3"
)

func main() {
	gtk3.Init(nil)
	clipboard := gtk3.NewClipboardGetForDisplay(
		gdk3.DisplayGetDefault(),
		gdk3.SELECTION_CLIPBOARD)
	fmt.Println(clipboard.WaitForText())
	clipboard.SetText("helloworld")
	gtk3.MainIterationDo(true)
	clipboard.Store()
	gtk3.MainIterationDo(true)
}
