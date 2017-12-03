package gtk3

// #include "gtk3.go.h"
// #cgo pkg-config: gtk+-3.0
import "C"

//-----------------------------------------------------------------------
// Resource Files
//
// In GTK+ 3.0, resource files have been deprecated and replaced by:
//     CSS-like style sheets, which are understood by GtkCssProvider.
//-----------------------------------------------------------------------

// deprecated
func RCParse(file string) {
	ptr := C.CString(file)
	defer cfree(ptr)
	C.gtk_rc_parse((*C.gchar)(ptr))
}

// deprecated
func RCParseString(style string) {
	ptr := C.CString(style)
	defer cfree(ptr)
	C.gtk_rc_parse_string((*C.gchar)(ptr))
}

// deprecated
func RCReparseAll() bool {
	return gobool(C.gtk_rc_reparse_all())
}


// deprecated
func RCResetStyles(settings *Settings) {
	C.gtk_rc_reset_styles(settings.GSettings)
}

// gtk_rc_reparse_all_for_settings
// gtk_rc_add_default_file
// gtk_rc_get_default_files
// gtk_rc_set_default_files
// gtk_rc_parse_color
// gtk_rc_parse_color_full
// gtk_rc_parse_state
// gtk_rc_parse_priority
// gtk_rc_find_module_in_path
// gtk_rc_find_pixmap_in_path
// gtk_rc_get_module_dir
// gtk_rc_get_im_module_path
// gtk_rc_get_im_module_file
// gtk_rc_get_theme_dir
// gtk_rc_style_new
// gtk_rc_style_copy
// gtk_rc_style_ref
// gtk_rc_style_unref
// gtk_rc_scanner_new
// gtk_rc_get_style
// gtk_rc_get_style_by_paths
// gtk_rc_add_widget_name_style
// gtk_rc_add_widget_class_style
// gtk_rc_add_class_style