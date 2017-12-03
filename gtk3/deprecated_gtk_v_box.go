package gtk3

// #include "gtk3.go.h"
// #cgo pkg-config: gtk+-3.0
import "C"

//-----------------------------------------------------------------------
// GtkVBox
//  A vertical container box.
//
// Recommendation is to switch to GtkGrid.
//
//-----------------------------------------------------------------------

// deprecated
type VBox struct {
	Box
}

// deprecated
func NewVBox(homogeneous bool, spacing int) *VBox {
	return &VBox{Box{Container{Widget{C.gtk_vbox_new(gbool(homogeneous), gint(spacing))}}}}
}
