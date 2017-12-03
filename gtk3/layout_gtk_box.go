package gtk3

// #include "gtk3.go.h"
// #cgo pkg-config: gtk+-3.0
import "C"

//-----------------------------------------------------------------------
// GtkBox
//  A container box.
//
// Functions:
// - GtkWidget * gtk_box_new ()
//	void 		gtk_box_pack_start ()
//	void 		gtk_box_pack_end ()
//	gboolean 	gtk_box_get_homogeneous ()
//	void 		gtk_box_set_homogeneous ()
//	gint 		gtk_box_get_spacing ()
//	void 		gtk_box_set_spacing ()
//	void 		gtk_box_reorder_child ()
//	void 		gtk_box_query_child_packing ()
//	void 		gtk_box_set_child_packing ()
// - GtkBaselinePosition 	gtk_box_get_baseline_position ()
// - void 		gtk_box_set_baseline_position ()
// - GtkWidget * gtk_box_get_center_widget ()
// - void 		gtk_box_set_center_widget ()
//-----------------------------------------------------------------------
type PackType int

const (
	PACK_START PackType = iota
	PACK_END
)

type IBox interface {
	IContainer
	PackStart(child IWidget, expand bool, fill bool, padding uint)
	PackEnd(child IWidget, expand bool, fill bool, padding uint)
}
type Box struct {
	Container
}

func (v *Box) PackStart(child IWidget, expand bool, fill bool, padding uint) {
	C.gtk_box_pack_start(BOX(v), ToNative(child), gbool(expand),
		gbool(fill), guint(padding))
}

func (v *Box) PackEnd(child IWidget, expand bool, fill bool, padding uint) {
	C.gtk_box_pack_end(BOX(v), ToNative(child), gbool(expand), gbool(fill), guint(padding))
}

func (v *Box) GetHomogeneous() bool {
	return gobool(C.gtk_box_get_homogeneous(BOX(v)))
}

func (v *Box) SetHomogeneous(homogeneous bool) {
	C.gtk_box_set_homogeneous(BOX(v), gbool(homogeneous))
}

func (v *Box) GetSpacing() int {
	return int(C.gtk_box_get_spacing(BOX(v)))
}

func (v *Box) SetSpacing(spacing int) {
	C.gtk_box_set_spacing(BOX(v), gint(spacing))
}

func (v *Box) ReorderChild(child IWidget, position int) {
	C.gtk_box_reorder_child(BOX(v), ToNative(child), C.gint(position))
}

func (v *Box) QueryChildPacking(child IWidget, expand *bool, fill *bool, padding *uint, pack_type *PackType) {
	var cexpand, cfill C.gboolean
	var cpadding C.guint
	var cpack_type C.GtkPackType
	C.gtk_box_query_child_packing(BOX(v), ToNative(child), &cexpand, &cfill, &cpadding, &cpack_type)
	*expand = gobool(cexpand)
	*fill = gobool(cfill)
	*padding = uint(cpadding)
	*pack_type = PackType(cpack_type)
}

func (v *Box) SetChildPacking(child IWidget, expand, fill bool, padding uint, pack_type PackType) {
	C.gtk_box_set_child_packing(BOX(v), ToNative(child), gbool(expand), gbool(fill), guint(padding), C.GtkPackType(pack_type))
}

