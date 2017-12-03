package gtk3

// #include "gtk3.go.h"
// #cgo pkg-config: gtk+-3.0
import "C"
import (
	"unsafe"
	"github.com/zurek87/go-gtk3/gdk3"
)

//-----------------------------------------------------------------------
// GtkStyle
//
// Use GtkStyleContext
//-----------------------------------------------------------------------

// deprecated
type Style struct {
	GStyle *C.GtkStyle
}

// deprecated
func NewStyle() *Style {
	return &Style{C.gtk_style_new()}
}

// deprecated
func (v *Style) Copy() *Style {
	return &Style{C.gtk_style_copy(STYLE(v))}
}

// deprecated
func (v *Style) Attach(window *Window) *Style {
	return &Style{C.gtk_style_attach(STYLE(v), C.toGdkWindow(unsafe.Pointer(window)))}
}

// deprecated
func (v *Style) Detach() {
	C.gtk_style_detach(STYLE(v))
}

// deprecated
func (v *Style) SetBackground(window *Window, state StateType) {
	C.gtk_style_set_background(STYLE(v), C.toGdkWindow(unsafe.Pointer(window)), C.GtkStateType(state))
}


// deprecated
func (v *Style) LookupColor(colorName string) (*gdk3.Color, bool) {
	color_name := C.CString(colorName)
	defer cfree(color_name)
	color := new(gdk3.Color)
	b := C.gtk_style_lookup_color(v.GStyle, gstring(color_name), (*C.GdkColor)(unsafe.Pointer(&color.GColor)))
	return color, gobool(b)
}

// gtk_style_apply_default_background
// gtk_style_lookup_icon_set
// gtk_style_render_icon
// gtk_style_get_style_property
// gtk_style_get_valist
// gtk_style_get
// gtk_paint_arrow
// gtk_paint_box
// gtk_paint_box_gap
// gtk_paint_check
// gtk_paint_diamond
// gtk_paint_extension
// gtk_paint_flat_box
// gtk_paint_focus
// gtk_paint_handle
// gtk_paint_hline
// gtk_paint_option
// gtk_paint_polygon
// gtk_paint_shadow
// gtk_paint_shadow_gap
// gtk_paint_slider
// gtk_paint_spinner
// gtk_paint_tab
// gtk_paint_vline
// gtk_paint_expander
// gtk_paint_layout
// gtk_paint_resize_grip
// gtk_draw_insertion_cursor
// gtk_border_new
// gtk_border_copy
// gtk_border_free