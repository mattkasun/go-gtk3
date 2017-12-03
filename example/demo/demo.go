package main

import (
	"fmt"
	"github.com/zurek87/go-gtk3/gdkpixbuf"
	"github.com/zurek87/go-gtk3/glib"
	"github.com/zurek87/go-gtk3/gtk3"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

func uniq(strings []string) (ret []string) {
	return
}

func authors() []string {
	if b, err := exec.Command("git", "log").Output(); err == nil {
		lines := strings.Split(string(b), "\n")

		var a []string
		r := regexp.MustCompile(`^Author:\s*([^ <]+).*$`)
		for _, e := range lines {
			ms := r.FindStringSubmatch(e)
			if ms == nil {
				continue
			}
			a = append(a, ms[1])
		}
		sort.Strings(a)
		var p string
		lines = []string{}
		for _, e := range a {
			if p == e {
				continue
			}
			lines = append(lines, e)
			p = e
		}
		return lines
	}
	return []string{"Yasuhiro Matsumoto <mattn.jp@gmail.com>"}
}

func main() {
	gtk3.Init(&os.Args)

	window := gtk3.NewWindow(gtk3.WINDOW_TOPLEVEL)
	window.SetPosition(gtk3.WIN_POS_CENTER)
	window.SetTitle("GTK Go!")
	window.SetIconName("gtk-dialog-info")
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		fmt.Println("got destroy!", ctx.Data().(string))
		gtk3.MainQuit()
	}, "foo")

	//--------------------------------------------------------
	// GtkVBox
	//--------------------------------------------------------
	vbox := gtk3.NewVBox(false, 1)

	//--------------------------------------------------------
	// GtkMenuBar
	//--------------------------------------------------------
	menubar := gtk3.NewMenuBar()
	vbox.PackStart(menubar, false, false, 0)

	//--------------------------------------------------------
	// GtkVPaned
	//--------------------------------------------------------
	vpaned := gtk3.NewVPaned()
	vbox.Add(vpaned)

	//--------------------------------------------------------
	// GtkFrame
	//--------------------------------------------------------
	frame1 := gtk3.NewFrame("Demo")
	framebox1 := gtk3.NewVBox(false, 1)
	frame1.Add(framebox1)

	frame2 := gtk3.NewFrame("Demo")
	framebox2 := gtk3.NewVBox(false, 1)
	frame2.Add(framebox2)

	vpaned.Pack1(frame1, false, false)
	vpaned.Pack2(frame2, false, false)

	//--------------------------------------------------------
	// GtkImage
	//--------------------------------------------------------
	dir, _ := filepath.Split(os.Args[0])
	imagefile := filepath.Join(dir, "github.com/zurek87/go-gtk3/data/go-gtk-logo.png")

	label := gtk3.NewLabel("Go Binding for GTK")
	label.ModifyFontEasy("DejaVu Serif 15")
	framebox1.PackStart(label, false, true, 0)

	//--------------------------------------------------------
	// GtkEntry
	//--------------------------------------------------------
	entry := gtk3.NewEntry()
	entry.SetText("Hello world")
	framebox1.Add(entry)

	image := gtk3.NewImageFromFile(imagefile)
	framebox1.Add(image)

	//--------------------------------------------------------
	// GtkScale
	//--------------------------------------------------------
	scale := gtk3.NewHScaleWithRange(0, 100, 1)
	scale.Connect("value-changed", func() {
		fmt.Println("scale:", int(scale.GetValue()))
	})
	framebox2.Add(scale)

	//--------------------------------------------------------
	// GtkHBox
	//--------------------------------------------------------
	buttons := gtk3.NewHBox(false, 1)

	//--------------------------------------------------------
	// GtkButton
	//--------------------------------------------------------
	button := gtk3.NewButtonWithLabel("Button with label")
	button.Clicked(func() {
		fmt.Println("button clicked:", button.GetLabel())
		messagedialog := gtk3.NewMessageDialog(
			button.GetTopLevelAsWindow(),
			gtk3.DIALOG_MODAL,
			gtk3.MESSAGE_INFO,
			gtk3.BUTTONS_OK,
			entry.GetText())
		messagedialog.Response(func() {
			fmt.Println("Dialog OK!")

			//--------------------------------------------------------
			// GtkFileChooserDialog
			//--------------------------------------------------------
			filechooserdialog := gtk3.NewFileChooserDialog(
				"Choose File...",
				button.GetTopLevelAsWindow(),
				gtk3.FILE_CHOOSER_ACTION_OPEN,
				gtk3.STOCK_OK,
				gtk3.RESPONSE_ACCEPT)
			filter := gtk3.NewFileFilter()
			filter.AddPattern("*.go")
			filechooserdialog.AddFilter(filter)
			filechooserdialog.Response(func() {
				fmt.Println(filechooserdialog.GetFilename())
				filechooserdialog.Destroy()
			})
			filechooserdialog.Run()
			messagedialog.Destroy()
		})
		messagedialog.Run()
	})
	buttons.Add(button)

	//--------------------------------------------------------
	// GtkFontButton
	//--------------------------------------------------------
	fontbutton := gtk3.NewFontButton()
	fontbutton.Connect("font-set", func() {
		fmt.Println("title:", fontbutton.GetTitle())
		fmt.Println("fontname:", fontbutton.GetFontName())
		fmt.Println("use_size:", fontbutton.GetUseSize())
		fmt.Println("show_size:", fontbutton.GetShowSize())
	})
	buttons.Add(fontbutton)
	framebox2.PackStart(buttons, false, false, 0)

	buttons = gtk3.NewHBox(false, 1)

	//--------------------------------------------------------
	// GtkToggleButton
	//--------------------------------------------------------
	togglebutton := gtk3.NewToggleButtonWithLabel("ToggleButton with label")
	togglebutton.Connect("toggled", func() {
		if togglebutton.GetActive() {
			togglebutton.SetLabel("ToggleButton ON!")
		} else {
			togglebutton.SetLabel("ToggleButton OFF!")
		}
	})
	buttons.Add(togglebutton)

	//--------------------------------------------------------
	// GtkCheckButton
	//--------------------------------------------------------
	checkbutton := gtk3.NewCheckButtonWithLabel("CheckButton with label")
	checkbutton.Connect("toggled", func() {
		if checkbutton.GetActive() {
			checkbutton.SetLabel("CheckButton CHECKED!")
		} else {
			checkbutton.SetLabel("CheckButton UNCHECKED!")
		}
	})
	buttons.Add(checkbutton)

	//--------------------------------------------------------
	// GtkRadioButton
	//--------------------------------------------------------
	buttonbox := gtk3.NewVBox(false, 1)
	radiofirst := gtk3.NewRadioButtonWithLabel(nil, "Radio1")
	buttonbox.Add(radiofirst)
	buttonbox.Add(gtk3.NewRadioButtonWithLabel(radiofirst.GetGroup(), "Radio2"))
	buttonbox.Add(gtk3.NewRadioButtonWithLabel(radiofirst.GetGroup(), "Radio3"))
	buttons.Add(buttonbox)
	//radiobutton.SetMode(false);
	radiofirst.SetActive(true)

	framebox2.PackStart(buttons, false, false, 0)

	//--------------------------------------------------------
	// GtkVSeparator
	//--------------------------------------------------------
	vsep := gtk3.NewVSeparator()
	framebox2.PackStart(vsep, false, false, 0)

	//--------------------------------------------------------
	// GtkComboBoxEntry
	//--------------------------------------------------------
	combos := gtk3.NewHBox(false, 1)
	comboboxentry := gtk3.NewComboBoxText()
	comboboxentry.AppendText("Monkey")
	comboboxentry.AppendText("Tiger")
	comboboxentry.AppendText("Elephant")
	comboboxentry.Connect("changed", func() {
		fmt.Println("value:", comboboxentry.GetActiveText())
	})
	combos.Add(comboboxentry)

	//--------------------------------------------------------
	// GtkComboBox
	//--------------------------------------------------------
	combobox := gtk3.NewComboBoxText()
	combobox.AppendText("Peach")
	combobox.AppendText("Banana")
	combobox.AppendText("Apple")
	combobox.SetActive(1)
	combobox.Connect("changed", func() {
		fmt.Println("value:", combobox.GetActiveText())
	})
	combos.Add(combobox)

	framebox2.PackStart(combos, false, false, 0)

	//--------------------------------------------------------
	// GtkTextView
	//--------------------------------------------------------
	swin := gtk3.NewScrolledWindow(nil, nil)
	swin.SetPolicy(gtk3.POLICY_AUTOMATIC, gtk3.POLICY_AUTOMATIC)
	swin.SetShadowType(gtk3.SHADOW_IN)
	textview := gtk3.NewTextView()
	var start, end gtk3.TextIter
	buffer := textview.GetBuffer()
	buffer.GetStartIter(&start)
	buffer.Insert(&start, "Hello ")
	buffer.GetEndIter(&end)
	buffer.Insert(&end, "World!")
	tag := buffer.CreateTag("bold", map[string]string{
		"background": "#FF0000", "weight": "700"})
	buffer.GetStartIter(&start)
	buffer.GetEndIter(&end)
	buffer.ApplyTag(tag, &start, &end)
	swin.Add(textview)
	framebox2.Add(swin)

	buffer.Connect("changed", func() {
		fmt.Println("changed")
	})

	//--------------------------------------------------------
	// GtkMenuItem
	//--------------------------------------------------------
	cascademenu := gtk3.NewMenuItemWithMnemonic("_File")
	menubar.Append(cascademenu)
	submenu := gtk3.NewMenu()
	cascademenu.SetSubmenu(submenu)

	var menuitem *gtk3.MenuItem
	menuitem = gtk3.NewMenuItemWithMnemonic("E_xit")
	menuitem.Connect("activate", func() {
		gtk3.MainQuit()
	})
	submenu.Append(menuitem)

	cascademenu = gtk3.NewMenuItemWithMnemonic("_View")
	menubar.Append(cascademenu)
	submenu = gtk3.NewMenu()
	cascademenu.SetSubmenu(submenu)

	checkmenuitem := gtk3.NewCheckMenuItemWithMnemonic("_Disable")
	checkmenuitem.Connect("activate", func() {
		vpaned.SetSensitive(!checkmenuitem.GetActive())
	})
	submenu.Append(checkmenuitem)

	menuitem = gtk3.NewMenuItemWithMnemonic("_Font")
	menuitem.Connect("activate", func() {
		fsd := gtk3.NewFontSelectionDialog("Font")
		fsd.SetFontName(fontbutton.GetFontName())
		fsd.Response(func() {
			fmt.Println(fsd.GetFontName())
			fontbutton.SetFontName(fsd.GetFontName())
			fsd.Destroy()
		})
		fsd.SetTransientFor(window)
		fsd.Run()
	})
	submenu.Append(menuitem)

	cascademenu = gtk3.NewMenuItemWithMnemonic("_Help")
	menubar.Append(cascademenu)
	submenu = gtk3.NewMenu()
	cascademenu.SetSubmenu(submenu)

	menuitem = gtk3.NewMenuItemWithMnemonic("_About")
	menuitem.Connect("activate", func() {
		dialog := gtk3.NewAboutDialog()
		dialog.SetName("Go-Gtk Demo!")
		dialog.SetProgramName("demo")
		dialog.SetAuthors(authors())
		dir, _ := filepath.Split(os.Args[0])
		imagefile := filepath.Join(dir, "../../data/mattn-logo.png")
		pixbuf, _ := gdkpixbuf.NewPixbufFromFile(imagefile)
		dialog.SetLogo(pixbuf)
		dialog.SetLicense("The library is available under the same terms and conditions as the Go, the BSD style license, and the LGPL (Lesser GNU Public License). The idea is that if you can use Go (and Gtk) in a project, you should also be able to use go-gtk.")
		dialog.SetWrapLicense(true)
		dialog.Run()
		dialog.Destroy()
	})
	submenu.Append(menuitem)

	//--------------------------------------------------------
	// GtkStatusbar
	//--------------------------------------------------------
	statusbar := gtk3.NewStatusbar()
	context_id := statusbar.GetContextId("go-gtk")
	statusbar.Push(context_id, "GTK binding for Go!")

	framebox2.PackStart(statusbar, false, false, 0)

	//--------------------------------------------------------
	// Event
	//--------------------------------------------------------
	window.Add(vbox)
	window.SetSizeRequest(600, 600)
	window.ShowAll()
	gtk3.Main()
}
