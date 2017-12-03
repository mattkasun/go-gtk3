package gtk3

// #include "gtk3.go.h"
// #cgo pkg-config: gtk+-3.0
import "C"

//-----------------------------------------------------------------------
// GtkVScale
//
// Use gtk_scale_new() with GTK_ORIENTATION_VERTICAL instead
//-----------------------------------------------------------------------

// deprecated
func NewVScale(a *Adjustment) *Scale {
	return &Scale{Range{Widget{C.gtk_vscale_new(a.GAdjustment)}}}
}

// deprecated
func NewVScaleWithRange(min, max, step float64) *Scale {
	return &Scale{Range{Widget{C.gtk_vscale_new_with_range(gdouble(min), gdouble(max), gdouble(step))}}}
}
