package gtk3

// #include "gtk3.go.h"
// #cgo pkg-config: gtk+-3.0
import "C"

//-----------------------------------------------------------------------
// GtkGrid
//  Pack widgets in rows and columns.
//
// Functions:
// - GtkWidget * gtk_grid_new ()
// - void 		gtk_grid_attach ()
// - void  		gtk_grid_attach_next_to ()
// - GtkWidget * gtk_grid_get_child_at ()
// - void  		gtk_grid_insert_row ()
// - void  		gtk_grid_insert_column ()
// - void  		gtk_grid_remove_row ()
// - void  		gtk_grid_remove_column ()
// - void  		gtk_grid_insert_next_to ()
// - void  		gtk_grid_set_row_homogeneous ()
// - gboolean 	gtk_grid_get_row_homogeneous ()
// - void  		gtk_grid_set_row_spacing ()
// - guint 		gtk_grid_get_row_spacing ()
// - void  		gtk_grid_set_column_homogeneous ()
// - gboolean 	gtk_grid_get_column_homogeneous ()
// - void 		gtk_grid_set_column_spacing ()
// - guint 		gtk_grid_get_column_spacing ()
// - gint 		gtk_grid_get_baseline_row ()
// - void 		gtk_grid_set_baseline_row ()
// - GtkBaselinePosition 	gtk_grid_get_row_baseline_position ()
// - void 		gtk_grid_set_row_baseline_position ()
//-----------------------------------------------------------------------

