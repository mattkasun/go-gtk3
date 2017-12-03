package main

import (
	"os"

	"github.com/zurek87/go-gtk3/example/builder/callback"
	"github.com/zurek87/go-gtk3/gtk3"
)

//"github.com/zurek87/go-gtk3/example/builder/callback"
func main() {
	gtk3.Init(&os.Args)

	builder := gtk3.NewBuilder()

	builder.AddFromFile("hello.ui")
	obj := builder.GetObject("window1")

	window := gtk3.WidgetFromObject(obj)
	window.Show()
	window.Connect("destroy", gtk3.MainQuit)

	callback.Init(builder)

	gtk3.Main()
}
