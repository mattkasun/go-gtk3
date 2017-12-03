package gtk3

// #include "gtk3.go.h"
// #cgo pkg-config: gtk+-3.0
import "C"

//-----------------------------------------------------------------------
// GtkListBox
//  A list container.
//
// Functions:
// - gboolean 	(*GtkListBoxFilterFunc) ()
// - gint 	(*GtkListBoxSortFunc) ()
// - void 	(*GtkListBoxUpdateHeaderFunc) ()
// - GtkWidget * 	gtk_list_box_new ()
// - void 	gtk_list_box_prepend ()
// - void 	gtk_list_box_insert ()
// - void 	gtk_list_box_select_row ()
// - void 	gtk_list_box_unselect_row ()
// - void 	gtk_list_box_select_all ()
// - void 	gtk_list_box_unselect_all ()
// - GtkListBoxRow * 	gtk_list_box_get_selected_row ()
// - void 	(*GtkListBoxForeachFunc) ()
// - void 	gtk_list_box_selected_foreach ()
// - GList * 	gtk_list_box_get_selected_rows ()
// - void 	gtk_list_box_set_selection_mode ()
// - GtkSelectionMode 	gtk_list_box_get_selection_mode ()
// - void 	gtk_list_box_set_activate_on_single_click ()
// - gboolean 	gtk_list_box_get_activate_on_single_click ()
// - GtkAdjustment * 	gtk_list_box_get_adjustment ()
// - void 	gtk_list_box_set_adjustment ()
// - void 	gtk_list_box_set_placeholder ()
// - GtkListBoxRow * 	gtk_list_box_get_row_at_index ()
// - GtkListBoxRow * 	gtk_list_box_get_row_at_y ()
// - void 	gtk_list_box_invalidate_filter ()
// - void 	gtk_list_box_invalidate_headers ()
// - void 	gtk_list_box_invalidate_sort ()
// - void 	gtk_list_box_set_filter_func ()
// - void 	gtk_list_box_set_header_func ()
// - void 	gtk_list_box_set_sort_func ()
// - void 	gtk_list_box_drag_highlight_row ()
// - void 	gtk_list_box_drag_unhighlight_row ()
// - GtkWidget * 	(*GtkListBoxCreateWidgetFunc) ()
// - void 	gtk_list_box_bind_model ()
// - GtkWidget * 	gtk_list_box_row_new ()
// - void 	gtk_list_box_row_changed ()
// - gboolean 	gtk_list_box_row_is_selected ()
// - GtkWidget * 	gtk_list_box_row_get_header ()
// - void 	gtk_list_box_row_set_header ()
// - gint 	gtk_list_box_row_get_index ()
// - void 	gtk_list_box_row_set_activatable ()
// - gboolean 	gtk_list_box_row_get_activatable ()
// - void 	gtk_list_box_row_set_selectable ()
// - gboolean 	gtk_list_box_row_get_selectable ()
//-----------------------------------------------------------------------

