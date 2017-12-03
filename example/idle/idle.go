package main

import (
	"fmt"
	"time"

	"github.com/zurek87/go-gtk3/glib"
)

func main() {
	glib.IdleAdd(func() bool {
		fmt.Println("start")
		return false
	})
	glib.TimeoutAdd(1000, func() bool {
		fmt.Println(fmt.Sprintf("%v", time.Now()))
		return true
	})

	glib.NewMainLoop(nil, false).Run()
}
