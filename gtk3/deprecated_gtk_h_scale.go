package gtk3

// #include "gtk3.go.h"
// #cgo pkg-config: gtk+-3.0
import "C"

//-----------------------------------------------------------------------
// GtkHScale
//
// Use gtk_scale_new() with GTK_ORIENTATION_HORIZONTAL instead
//-----------------------------------------------------------------------

// deprecated
func NewHScale(adjustment *Adjustment) *Scale {
	return &Scale{Range{Widget{C.gtk_hscale_new(adjustment.GAdjustment)}}}
}

// deprecated
func NewHScaleWithRange(min, max, step float64) *Scale {
	return &Scale{Range{Widget{
		C.gtk_hscale_new_with_range(gdouble(min), gdouble(max), gdouble(step))}}}
}

