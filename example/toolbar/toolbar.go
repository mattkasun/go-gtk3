package main

import (
	"fmt"
	"github.com/zurek87/go-gtk3/glib"
	"github.com/zurek87/go-gtk3/gtk3"
)

func main() {
	gtk3.Init(nil)
	window := gtk3.NewWindow(gtk3.WINDOW_TOPLEVEL)
	window.SetPosition(gtk3.WIN_POS_CENTER)
	window.SetTitle("GTK Go!")
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		gtk3.MainQuit()
	}, "")

	vbox := gtk3.NewVBox(false, 0)

	toolbar := gtk3.NewToolbar()
	toolbar.SetStyle(gtk3.TOOLBAR_ICONS)
	vbox.PackStart(toolbar, false, false, 5)

	btnnew := gtk3.NewToolButtonFromStock(gtk3.STOCK_NEW)
	btnclose := gtk3.NewToolButtonFromStock(gtk3.STOCK_CLOSE)
	separator := gtk3.NewSeparatorToolItem()
	btncustom := gtk3.NewToolButton(nil, "Custom")
	btnmenu := gtk3.NewMenuToolButtonFromStock("gtk.STOCK_CLOSE")
	btnmenu.SetArrowTooltipText("This is a tool tip")

	btnnew.OnClicked(onToolButtonClicked)
	btnclose.OnClicked(onToolButtonClicked)
	btncustom.OnClicked(onToolButtonClicked)

	toolmenu := gtk3.NewMenu()
	menuitem := gtk3.NewMenuItemWithMnemonic("8")
	menuitem.Show()
	toolmenu.Append(menuitem)
	menuitem = gtk3.NewMenuItemWithMnemonic("16")
	menuitem.Show()
	toolmenu.Append(menuitem)
	menuitem = gtk3.NewMenuItemWithMnemonic("32")
	menuitem.Show()
	toolmenu.Append(menuitem)
	btnmenu.SetMenu(toolmenu)

	toolbar.Insert(btnnew, -1)
	toolbar.Insert(btnclose, -1)
	toolbar.Insert(separator, -1)
	toolbar.Insert(btncustom, -1)
	toolbar.Insert(btnmenu, -1)

	hbox := gtk3.NewHBox(false, 0)
	vbox.PackStart(hbox, true, true, 0)

	toolbar2 := gtk3.NewToolbar()
	toolbar2.SetOrientation(gtk3.ORIENTATION_VERTICAL)
	hbox.PackStart(toolbar2, false, false, 5)
	btnhelp := gtk3.NewToolButtonFromStock(gtk3.STOCK_HELP)
	btnhelp.OnClicked(onToolButtonClicked)
	toolbar2.Insert(btnhelp, -1)

	btntoggle := gtk3.NewToggleToolButton()
	btntoggle2 := gtk3.NewToggleToolButtonFromStock(gtk3.STOCK_ITALIC)
	toolbar2.Insert(btntoggle, -1)
	toolbar2.Insert(btntoggle2, -1)

	for i := 0; i < toolbar.GetNItems(); i++ {
		gti := toolbar.GetNthItem(i)
		switch gti.(type) {
		case *gtk3.ToolButton:
			fmt.Printf("toolbar[%d] is a *gtk.ToolButton\n", i)
			w := gti.(*gtk3.ToolButton).GetIconWidget()
			gti.(*gtk3.ToolButton).SetIconWidget(w)
		case *gtk3.ToggleToolButton:
			fmt.Printf("toolbar[%d] is a *gtk.ToggleToolButton\n", i)
			gti.(*gtk3.ToggleToolButton).SetActive(true)
		case *gtk3.SeparatorToolItem:
			fmt.Printf("toolbar[%d] is a *gtk.SeparatorToolItem\n", i)
		default:
			fmt.Printf("toolbar: Item is of unknown type\n")
		}
	}

	for i := 0; i < toolbar2.GetNItems(); i++ {
		gti := toolbar2.GetNthItem(i)
		switch gti.(type) {
		case *gtk3.ToolButton:
			fmt.Printf("toolbar2[%d] is a *gtk.ToolButton\n", i)
		case *gtk3.ToggleToolButton:
			fmt.Printf("toolbar2[%d] is a *gtk.ToggleToolButton\n", i)
			gti.(*gtk3.ToggleToolButton).SetActive(true)
		case *gtk3.SeparatorToolItem:
			fmt.Printf("toolbar2[%d] is a *gtk.SeparatorToolItem\n", i)
		default:
			fmt.Printf("toolbar2: Item is of unknown type\n")
		}
	}

	window.Add(vbox)
	window.SetSizeRequest(600, 600)
	window.ShowAll()
	gtk3.Main()
}

func onToolButtonClicked() {
	fmt.Println("ToolButton clicked")
}
