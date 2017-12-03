package gtk3

// #include "gtk3.go.h"
// #cgo pkg-config: gtk+-3.0
import "C"

//-----------------------------------------------------------------------
// GtkPaned
//  A widget with two adjustable panes.
//
// Functions:
// - GtkWidget * gtk_paned_new () -
// void 		gtk_paned_add1 ()
// void 		gtk_paned_add2 ()
// void 		gtk_paned_pack1 ()
// void 		gtk_paned_pack2 ()
// GtkWidget * 	gtk_paned_get_child1 ()
// GtkWidget * 	gtk_paned_get_child2 ()
// void 		gtk_paned_set_position ()
// gint 		gtk_paned_get_position ()
// - GdkWindow * 	gtk_paned_get_handle_window () -
// - void 		gtk_paned_set_wide_handle () -
// - gboolean 	gtk_paned_get_wide_handle () -
//-----------------------------------------------------------------------

type IPaned interface {
	IContainer
	Add1(child IWidget)
	Add2(child IWidget)
	Pack1(child IWidget, expand bool, fill bool)
	Pack2(child IWidget, expand bool, fill bool)
}
type Paned struct {
	Container
}

func (v *Paned) Add1(child IWidget) {
	C.gtk_paned_add1(PANED(v), ToNative(child))
}

func (v *Paned) Add2(child IWidget) {
	C.gtk_paned_add2(PANED(v), ToNative(child))
}

func (v *Paned) Pack1(child IWidget, resize bool, shrink bool) {
	C.gtk_paned_pack1(PANED(v), ToNative(child), gbool(resize), gbool(shrink))
}

func (v *Paned) Pack2(child IWidget, resize bool, shrink bool) {
	C.gtk_paned_pack2(PANED(v), ToNative(child), gbool(resize), gbool(shrink))
}

func (v *Paned) GetChild1() *Widget {
	return &Widget{C.gtk_paned_get_child1(PANED(v))}
}

func (v *Paned) GetChild2() *Widget {
	return &Widget{C.gtk_paned_get_child2(PANED(v))}
}

func (v *Paned) SetPosition(position int) {
	C.gtk_paned_set_position(PANED(v), gint(position))
}

func (v *Paned) GetPosition() int {
	return int(C.gtk_paned_get_position(PANED(v)))
}
