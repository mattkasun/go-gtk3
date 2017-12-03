package gtk3

// #include "gtk3.go.h"
// #cgo pkg-config: gtk+-3.0
import "C"


//-----------------------------------------------------------------------
// GtkFontSelection -> Use GtkFontChooserWidget instead
//
// GtkFontSelectionDialog -> Use GtkFontChooserDialog
//-----------------------------------------------------------------------

// deprecated
type FontSelection struct {
	VBox
}

// deprecated
func NewFontSelection() *FontSelection {
	return &FontSelection{VBox{Box{Container{Widget{C.gtk_font_selection_new()}}}}}
}

// deprecated
func (v *FontSelection) GetFontName() string {
	return gostring(C.gtk_font_selection_get_font_name(FONT_SELECTION(v)))
}

// deprecated
func (v *FontSelection) SetFontName(name string) {
	ptr := C.CString(name)
	defer cfree(ptr)
	C.gtk_font_selection_set_font_name(FONT_SELECTION(v), gstring(ptr))
}


// deprecated
type FontSelectionDialog struct {
	Dialog
}

// deprecated
func NewFontSelectionDialog(title string) *FontSelectionDialog {
	ptitle := C.CString(title)
	defer cfree(ptitle)
	return &FontSelectionDialog{Dialog{Window{Bin{Container{Widget{
		C.gtk_font_selection_dialog_new(gstring(ptitle))}}}}, nil}}
}

// deprecated
func (v *FontSelectionDialog) GetFontName() string {
	return gostring(C.gtk_font_selection_dialog_get_font_name(FONT_SELECTION_DIALOG(v)))
}

// deprecated
func (v *FontSelectionDialog) SetFontName(font string) {
	pfont := C.CString(font)
	defer cfree(pfont)
	C.gtk_font_selection_dialog_set_font_name(FONT_SELECTION_DIALOG(v), gstring(pfont))
}

// deprecated
func (v *FontSelectionDialog) GetPreviewText() string {
	return gostring(C.gtk_font_selection_dialog_get_preview_text(FONT_SELECTION_DIALOG(v)))
}

// deprecated
func (v *FontSelectionDialog) SetPreviewText(text string) {
	ptr := C.CString(text)
	defer cfree(ptr)
	C.gtk_font_selection_dialog_set_preview_text(FONT_SELECTION_DIALOG(v), gstring(ptr))
}

// deprecated
func (v *FontSelectionDialog) GetCancelButton() *Widget {
	return &Widget{C.gtk_font_selection_dialog_get_cancel_button(FONT_SELECTION_DIALOG(v))}
}

// deprecated
func (v *FontSelectionDialog) GetOkButton() *Widget {
	return &Widget{C.gtk_font_selection_dialog_get_ok_button(FONT_SELECTION_DIALOG(v))}
}

// deprecated
func (v *FontSelectionDialog) GetFontSelection() *FontSelection {
	return &FontSelection{VBox{Box{Container{Widget{
		C.gtk_font_selection_dialog_get_font_selection(FONT_SELECTION_DIALOG(v))}}}}}
}