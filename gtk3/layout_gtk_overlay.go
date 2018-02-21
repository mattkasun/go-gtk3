package gtk3

// #include "gtk3.go.h"
// #cgo pkg-config: gtk+-3.0
import "C"

//-----------------------------------------------------------------------
// GtkStackSwitcher
//  A controller for GtkStack.
//

type Overlay struct {
	Container
}

func OVERLAY(p *Overlay) *C.GtkOverlay {return C.toGOverlay(p.GWidget)}

// Functions:
// - GtkWidget * 	gtk_overlay_new ()
func NewOverlay() *Overlay {
	return &Overlay{Container{Widget{C.gtk_overlay_new()}}}
	}

// - void 	gtk_overlay_add_overlay ()
func (v *Overlay) AddOverlay(w IWidget) {
	C.gtk_overlay_add_overlay(OVERLAY(v), ToNative(w))
}
// - void 	gtk_overlay_reorder_overlay ()
func (v *Overlay) ReorderOverly (w IWidget, p int) {
	C.gtk_overlay_reorder_overlay(OVERLAY(v), ToNative(w), C.gint(p))
}

// - gboolean 	gtk_overlay_get_overlay_pass_through ()
// - void 	gtk_overlay_set_overlay_pass_through ()
//-----------------------------------------------------------------------

