package gtk3

// #include "gtk3.go.h"
// #cgo pkg-config: gtk+-3.0
import "C"
import "unsafe"

//-----------------------------------------------------------------------
// GtkLayout
//  Infinite scrollable area containing child widgets and/or custom drawing.
//
// Functions:
//	GtkWidget * 	gtk_layout_new ()
//	void 	gtk_layout_put ()
//	void 	gtk_layout_move ()
//	void 	gtk_layout_set_size ()
//	void 	gtk_layout_get_size ()
//	GtkAdjustment * 	gtk_layout_get_hadjustment ()
//	GtkAdjustment * 	gtk_layout_get_vadjustment ()
//	void 	gtk_layout_set_hadjustment ()
//	void 	gtk_layout_set_vadjustment ()
//	GdkWindow * 	gtk_layout_get_bin_window ()
//-----------------------------------------------------------------------

type Layout struct {
	Container
}

func NewLayout(hadjustment *Adjustment, vadjustment *Adjustment) *Layout {
	var had, vad *C.GtkAdjustment
	if hadjustment != nil {
		had = hadjustment.GAdjustment
	}
	if vadjustment != nil {
		vad = vadjustment.GAdjustment
	}
	return &Layout{Container{Widget{
		C.gtk_layout_new(had, vad)}}}
}

func (v *Layout) GetHAdjustment() *Adjustment {
	return &Adjustment{
		C.gtk_layout_get_hadjustment(LAYOUT(v))}
}

func (v *Layout) GetVAdjustment() *Adjustment {
	return &Adjustment{
		C.gtk_layout_get_vadjustment(LAYOUT(v))}
}

func (v *Layout) SetHAdjustment(hadjustment *Adjustment) {
	C.gtk_layout_set_hadjustment(LAYOUT(v), hadjustment.GAdjustment)
}

func (v *Layout) SetVAdjustment(vadjustment *Adjustment) {
	C.gtk_layout_set_vadjustment(LAYOUT(v), vadjustment.GAdjustment)
}

func (v *Layout) Move(child IWidget, x int, y int) {
	C.gtk_layout_move(LAYOUT(v), ToNative(child), C.gint(x), C.gint(y))
}

func (v *Layout) Put(child IWidget, x int, y int) {
	C.gtk_layout_put(LAYOUT(v), ToNative(child), C.gint(x), C.gint(y))
}

func (v *Layout) SetSize(width int, height int) {
	C.gtk_layout_set_size(LAYOUT(v), C.guint(width), C.guint(height))
}

func (v *Layout) GetSize(width *int, height *int) {
	var cwidth, cheight C.guint
	C.gtk_layout_get_size(LAYOUT(v), &cwidth, &cheight)
	*width = int(cwidth)
	*height = int(cheight)
}

func (v *Layout) GetBinWindow() *Window {
	return &Window{Bin{Container{Widget{
		C.toGWidget(unsafe.Pointer(C.gtk_layout_get_bin_window(LAYOUT(v))))}}}}
}
