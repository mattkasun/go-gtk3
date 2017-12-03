package main

import "github.com/zurek87/go-gtk3/gtk3"
import "github.com/zurek87/go-gtk3/glib"
import "fmt"
import "syscall"

func main() {
	gtk3.SetLocale()

	bs, _, _, err := glib.LocaleFromUtf8("こんにちわ世界\n")
	if err == nil {
		syscall.Write(syscall.Stdout, bs)
	} else {
		fmt.Println(err.Message())
	}
}
