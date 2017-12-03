package gtk3

// #include "gtk3.go.h"
// #cgo pkg-config: gtk+-3.0
import "C"

//-----------------------------------------------------------------------
// GtkFlowBox
//  A container that allows reflowing its children.
//
// Functions:
// - GtkWidget * 	gtk_flow_box_new ()
// - void 	gtk_flow_box_insert ()
// - GtkFlowBoxChild * 	gtk_flow_box_get_child_at_index ()
// - GtkFlowBoxChild * 	gtk_flow_box_get_child_at_pos ()
// - void 	gtk_flow_box_set_hadjustment ()
// - void 	gtk_flow_box_set_vadjustment ()
// - void 	gtk_flow_box_set_homogeneous ()
// - gboolean 	gtk_flow_box_get_homogeneous ()
// - void 	gtk_flow_box_set_row_spacing ()
// - guint 	gtk_flow_box_get_row_spacing ()
// - void 	gtk_flow_box_set_column_spacing ()
// - guint 	gtk_flow_box_get_column_spacing ()
// - void 	gtk_flow_box_set_min_children_per_line ()
// - guint 	gtk_flow_box_get_min_children_per_line ()
// - void 	gtk_flow_box_set_max_children_per_line ()
// - guint 	gtk_flow_box_get_max_children_per_line ()
// - void 	gtk_flow_box_set_activate_on_single_click ()
// - gboolean 	gtk_flow_box_get_activate_on_single_click ()
// - void 	(*GtkFlowBoxForeachFunc) ()
// - void 	gtk_flow_box_selected_foreach ()
// - GList * 	gtk_flow_box_get_selected_children ()
// - void 	gtk_flow_box_select_child ()
// - void 	gtk_flow_box_unselect_child ()
// - void 	gtk_flow_box_select_all ()
// - void 	gtk_flow_box_unselect_all ()
// - void 	gtk_flow_box_set_selection_mode ()
// - GtkSelectionMode 	gtk_flow_box_get_selection_mode ()
// - gboolean 	(*GtkFlowBoxFilterFunc) ()
// - void 	gtk_flow_box_set_filter_func ()
// - void 	gtk_flow_box_invalidate_filter ()
// - gint 	(*GtkFlowBoxSortFunc) ()
// - void 	gtk_flow_box_set_sort_func ()
// - void 	gtk_flow_box_invalidate_sort ()
// - GtkWidget * 	(*GtkFlowBoxCreateWidgetFunc) ()
// - void 	gtk_flow_box_bind_model ()
// - GtkWidget * 	gtk_flow_box_child_new ()
// - gint 	gtk_flow_box_child_get_index ()
// - gboolean 	gtk_flow_box_child_is_selected ()
// - void 	gtk_flow_box_child_changed ()
//-----------------------------------------------------------------------

