package main

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/zurek87/go-gtk3/glib"
	"github.com/zurek87/go-gtk3/gtk3"
)

func CreateWindow() *gtk3.Window {
	window := gtk3.NewWindow(gtk3.WINDOW_TOPLEVEL)
	window.SetDefaultSize(700, 300)
	vbox := gtk3.NewVBox(false, 1)
	CreateMenuAndToolbar(window, vbox)
	CreateActivatableDemo(vbox)
	window.Add(vbox)
	return window
}

func CreateMenuAndToolbar(w *gtk3.Window, vbox *gtk3.VBox) {
	action_group := gtk3.NewActionGroup("my_group")
	ui_manager := CreateUIManager()
	accel_group := ui_manager.GetAccelGroup()
	w.AddAccelGroup(accel_group)
	AddFileMenuActions(action_group)
	AddEditMenuActions(action_group)
	AddChoicesMenuActions(action_group)
	ui_manager.InsertActionGroup(action_group, 0)
	menubar := ui_manager.GetWidget("/MenuBar")
	vbox.PackStart(menubar, false, false, 0)
	toolbar := ui_manager.GetWidget("/ToolBar")
	vbox.PackStart(toolbar, false, false, 0)
	eventbox := gtk3.NewEventBox()
	vbox.PackStart(eventbox, false, false, 0)
	//    label := gtk.NewLabel("Right-click to see the popup menu.")
	//    vbox.PackStart(label, false, false, 0)
}

func CreateActivatableDemo(vbox *gtk3.VBox) {
	action_entry := gtk3.NewAction("ActionEntry",
		"Button attached to Action", "", gtk3.STOCK_INFO)
	action_entry.Connect("activate", func() {
		fmt.Println("Action clicked")
	})
	frame1 := gtk3.NewFrame("GtkActivatable interface demonstration")
	frame1.SetBorderWidth(5)
	hbox2 := gtk3.NewHBox(false, 5)
	hbox2.SetSizeRequest(400, 50)
	hbox2.SetBorderWidth(5)
	button1 := gtk3.NewButton()
	button1.SetSizeRequest(250, 0)
	button1.SetRelatedAction(action_entry)
	hbox2.PackStart(button1, false, false, 0)
	hbox2.PackStart(gtk3.NewVSeparator(), false, false, 0)
	button2 := gtk3.NewButtonWithLabel("Hide Action")
	button2.SetSizeRequest(150, 0)
	button2.Connect("clicked", func() {
		action_entry.SetVisible(false)
		fmt.Println("Hide Action")
	})
	hbox2.PackStart(button2, false, false, 0)
	button3 := gtk3.NewButtonWithLabel("Unhide Action")
	button3.SetSizeRequest(150, 0)
	button3.Connect("clicked", func() {
		action_entry.SetVisible(true)
		fmt.Println("Show Action")
	})
	hbox2.PackStart(button3, false, false, 0)
	frame1.Add(hbox2)
	vbox.PackStart(frame1, false, true, 0)
}

func CreateUIManager() *gtk3.UIManager {
	UI_INFO := `
<ui>
  <menubar name='MenuBar'>
    <menu action='FileMenu'>
      <menu action='FileNew'>
        <menuitem action='FileNewStandard' />
        <menuitem action='FileNewFoo' />
        <menuitem action='FileNewGoo' />
      </menu>
      <separator />
      <menuitem action='FileQuit' />
    </menu>
    <menu action='EditMenu'>
      <menuitem action='EditCopy' />
      <menuitem action='EditPaste' />
      <menuitem action='EditSomething' />
    </menu>
    <menu action='ChoicesMenu'>
      <menuitem action='ChoiceOne'/>
      <menuitem action='ChoiceTwo'/>
      <menuitem action='ChoiceThree'/>
      <separator />
      <menuitem action='ChoiceToggle'/>
    </menu>
  </menubar>
  <toolbar name='ToolBar'>
    <toolitem action='FileNewStandard' />
    <toolitem action='FileQuit' />
  </toolbar>
  <popup name='PopupMenu'>
    <menuitem action='EditCopy' />
    <menuitem action='EditPaste' />
    <menuitem action='EditSomething' />
  </popup>
</ui>
`
	ui_manager := gtk3.NewUIManager()
	ui_manager.AddUIFromString(UI_INFO)
	return ui_manager
}

func OnMenuFileNewGeneric() {
	fmt.Println("A File|New menu item was selected.")
}

func OnMenuFileQuit() {
	fmt.Println("quit app...")
	gtk3.MainQuit()
}

func OnMenuOther(ctx *glib.CallbackContext) {
	v := reflect.ValueOf(ctx.Target())
	if v.Kind() == reflect.Ptr {
		fmt.Printf("Item %s(%p) was selected", v.Elem(), v.Interface())
		fmt.Println()
		if w, ok := v.Elem().Interface().(gtk3.IWidget); ok {
			v := reflect.ValueOf(ctx.Target())
			v2 := v.Elem()
			fmt.Println(v.Kind(), v2.Kind())
			fmt.Println("Menu item ", w.GetName(), " was selected")
		}
	}
}

func OnMenuOther2(widget unsafe.Pointer, event unsafe.Pointer,
	data unsafe.Pointer) {
	fmt.Println("Menu item ", "", " was selected")
}

func AddFileMenuActions(action_group *gtk3.ActionGroup) {
	action_group.AddAction(gtk3.NewAction("FileMenu", "File", "", ""))

	action_filenewmenu := gtk3.NewAction("FileNew", "", "", gtk3.STOCK_NEW)
	action_group.AddAction(action_filenewmenu)

	action_new := gtk3.NewAction("FileNewStandard", "_New",
		"Create a new file", gtk3.STOCK_NEW)
	action_new.Connect("activate", OnMenuFileNewGeneric)
	action_group.AddActionWithAccel(action_new, "")

	action_new_foo := gtk3.NewAction("FileNewFoo", "New Foo",
		"Create new foo", gtk3.STOCK_NEW)
	action_new_foo.Connect("activate", OnMenuFileNewGeneric)
	action_group.AddAction(action_new_foo)

	action_new_goo := gtk3.NewAction("FileNewGoo", "_New Goo",
		"Create new goo", gtk3.STOCK_NEW)
	action_new_goo.Connect("activate", OnMenuFileNewGeneric)
	action_group.AddAction(action_new_goo)

	action_filequit := gtk3.NewAction("FileQuit", "", "", gtk3.STOCK_QUIT)
	action_filequit.Connect("activate", OnMenuFileQuit)
	action_group.AddActionWithAccel(action_filequit, "")
}

func AddEditMenuActions(action_group *gtk3.ActionGroup) {
	action_group.AddAction(gtk3.NewAction("EditMenu", "Edit", "", ""))

	action_editcopy := gtk3.NewAction("EditCopy", "", "", gtk3.STOCK_COPY)
	action_editcopy.Connect("activate", OnMenuOther)
	action_group.AddActionWithAccel(action_editcopy, "")

	action_editpaste := gtk3.NewAction("EditPaste", "", "", gtk3.STOCK_PASTE)
	action_editpaste.Connect("activate", OnMenuOther)
	action_group.AddActionWithAccel(action_editpaste, "")

	action_editsomething := gtk3.NewAction("EditSomething", "Something", "", "")
	action_editsomething.Connect("activate", OnMenuOther)
	action_group.AddActionWithAccel(action_editsomething, "<control><alt>S")
}

func AddChoicesMenuActions(action_group *gtk3.ActionGroup) {
	action_group.AddAction(gtk3.NewAction("ChoicesMenu", "Choices", "", ""))

	var ra_list []*gtk3.RadioAction
	ra_one := gtk3.NewRadioAction("ChoiceOne", "One", "", "", 1)
	ra_list = append(ra_list, ra_one)

	ra_two := gtk3.NewRadioAction("ChoiceTwo", "Two", "", "", 2)
	ra_list = append(ra_list, ra_two)

	ra_three := gtk3.NewRadioAction("ChoiceThree", "Three", "", "", 2)
	ra_list = append(ra_list, ra_three)

	var sl *glib.SList
	for _, ra := range ra_list {
		ra.SetGroup(sl)
		sl = ra.GetGroup()
		action_group.AddAction(ra)
	}

	ra_last := gtk3.NewToggleAction("ChoiceToggle", "Toggle", "", "")
	ra_last.SetActive(true)
	action_group.AddAction(ra_last)
}

func main() {
	gtk3.Init(nil)
	window := CreateWindow()
	window.SetPosition(gtk3.WIN_POS_CENTER)
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		fmt.Println("destroy pending...")
		gtk3.MainQuit()
	}, "foo")
	window.ShowAll()
	gtk3.Main()
}
