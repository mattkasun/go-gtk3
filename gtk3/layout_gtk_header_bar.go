package gtk3

// #include "gtk3.go.h"
// #cgo pkg-config: gtk+-3.0
import "C"

//-----------------------------------------------------------------------
// GtkHeaderBar
//  A box with a centered child.
//
// Functions:
// - GtkWidget * 	gtk_header_bar_new ()
// - void 	gtk_header_bar_set_title ()
// - const gchar * 	gtk_header_bar_get_title ()
// - void 	gtk_header_bar_set_subtitle ()
// - const gchar * 	gtk_header_bar_get_subtitle ()
// - void 	gtk_header_bar_set_has_subtitle ()
// - gboolean 	gtk_header_bar_get_has_subtitle ()
// - void 	gtk_header_bar_set_custom_title ()
// - GtkWidget * 	gtk_header_bar_get_custom_title ()
// - void 	gtk_header_bar_pack_start ()
// - void 	gtk_header_bar_pack_end ()
// - void 	gtk_header_bar_set_show_close_button ()
// - gboolean 	gtk_header_bar_get_show_close_button ()
// - void 	gtk_header_bar_set_decoration_layout ()
// - const gchar * 	gtk_header_bar_get_decoration_layout ()
//-----------------------------------------------------------------------

