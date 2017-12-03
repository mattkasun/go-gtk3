package gtk3

// #include "gtk3.go.h"
// #cgo pkg-config: gtk+-3.0
import "C"

//-----------------------------------------------------------------------
// GtkHBox
//  A horizontal container box.
//
// Recommendation is to switch to GtkGrid.
//
//-----------------------------------------------------------------------

// deprecated
type HBox struct {
	Box
}

// deprecated
func NewHBox(homogeneous bool, spacing int) *HBox {
	return &HBox{Box{Container{Widget{C.gtk_hbox_new(gbool(homogeneous), gint(spacing))}}}}
}