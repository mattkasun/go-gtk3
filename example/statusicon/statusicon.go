package main

import (
	"fmt"
	"github.com/zurek87/go-gtk3/glib"
	"github.com/zurek87/go-gtk3/gtk3"
	"os"
)

func main() {
	gtk3.Init(&os.Args)

	glib.SetApplicationName("go-gtk-statusicon-example")

	mi := gtk3.NewMenuItemWithLabel("Popup!")
	mi.Connect("activate", func() {
		gtk3.MainQuit()
	})
	nm := gtk3.NewMenu()
	nm.Append(mi)
	nm.ShowAll()

	si := gtk3.NewStatusIconFromStock(gtk3.STOCK_FILE)
	si.SetTitle("StatusIcon Example")
	si.SetTooltipMarkup("StatusIcon Example")
	si.Connect("popup-menu", func(cbx *glib.CallbackContext) {
		nm.Popup(nil, nil, gtk3.StatusIconPositionMenu, si, uint(cbx.Args(0)), uint32(cbx.Args(1)))
	})

	fmt.Println(`
Can you see statusicon in systray?
If you don't see it and if you use 'unity', try following.

# gsettings set com.canonical.Unity.Panel systray-whitelist \
  "$(gsettings get com.canonical.Unity.Panel systray-whitelist \|
  sed -e "s/]$/, 'go-gtk-statusicon-example']/")"
`)

	gtk3.Main()
}
