package gtk3

// #include "gtk3.go.h"
// #cgo pkg-config: gtk+-3.0
import "C"

//-----------------------------------------------------------------------
// GtkExpander
//  A container which can hide its child.
//
// Functions:
// GtkWidget * 	gtk_expander_new ()
// GtkWidget * 	gtk_expander_new_with_mnemonic ()
// void 		gtk_expander_set_expanded ()
// gboolean 	gtk_expander_get_expanded ()
// void 		gtk_expander_set_spacing ()
// gint 		gtk_expander_get_spacing ()
// void 		gtk_expander_set_label ()
// const gchar * 	gtk_expander_get_label ()
// void 		gtk_expander_set_use_underline ()
// gboolean 	gtk_expander_get_use_underline ()
// void 		gtk_expander_set_use_markup ()
// gboolean 	gtk_expander_get_use_markup ()
// void 		gtk_expander_set_label_widget ()
// GtkWidget * 	gtk_expander_get_label_widget ()
// - void 		gtk_expander_set_label_fill ()		-
// - gboolean 	gtk_expander_get_label_fill ()		-
// - void 		gtk_expander_set_resize_toplevel ()	-
// - gboolean 	gtk_expander_get_resize_toplevel () -
//-----------------------------------------------------------------------

type Expander struct {
	Bin
}

func NewExpander(label string) *Expander {
	var ptr *C.char
	if len(label) > 0 {
		ptr = C.CString(label)
		defer cfree(ptr)
	}
	return &Expander{Bin{Container{Widget{C.gtk_expander_new(gstring(ptr))}}}}
}

func NewExpanderWithMnemonic(label string) *Expander {
	var ptr *C.char
	if len(label) > 0 {
		ptr = C.CString(label)
		defer cfree(ptr)
	}
	return &Expander{Bin{Container{Widget{C.gtk_expander_new_with_mnemonic(gstring(ptr))}}}}
}

func (v *Expander) SetExpanded(expanded bool) {
	C.gtk_expander_set_expanded(EXPANDER(v), gbool(expanded))
}

func (v *Expander) GetExpanded() bool {
	return gobool(C.gtk_expander_get_expanded(EXPANDER(v)))
}

func (v *Expander) SetSpacing(spacing int) {
	C.gtk_expander_set_spacing(EXPANDER(v), gint(spacing))
}

func (v *Expander) GetSpacing() int {
	return int(C.gtk_expander_get_spacing(EXPANDER(v)))
}

func (v *Expander) SetLabel(label string) {
	var ptr *C.char
	if len(label) > 0 {
		ptr = C.CString(label)
		defer cfree(ptr)
	}
	C.gtk_expander_set_label(EXPANDER(v), gstring(ptr))
}

func (v *Expander) GetLabel() string {
	return gostring(C.gtk_expander_get_label(EXPANDER(v)))
}

func (v *Expander) SetUseUnderline(setting bool) {
	C.gtk_expander_set_use_underline(EXPANDER(v), gbool(setting))
}

func (v *Expander) GetUseUnderline() bool {
	return gobool(C.gtk_expander_get_use_underline(EXPANDER(v)))
}

func (v *Expander) SetUseMarkup(setting bool) {
	C.gtk_expander_set_use_markup(EXPANDER(v), gbool(setting))
}

func (v *Expander) GetUseMarkup() bool {
	return gobool(C.gtk_expander_get_use_markup(EXPANDER(v)))
}

func (v *Expander) SetLabelWidget(label_widget ILabel) {
	C.gtk_expander_set_label_widget(EXPANDER(v), ToNative(label_widget))
}

func (v *Expander) GetLabelWidget() ILabel {
	return &Label{Misc{Widget{
		C.gtk_expander_get_label_widget(EXPANDER(v))}}}
}

