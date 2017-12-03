package gtk3

// #include "gtk3.go.h"
// #cgo pkg-config: gtk+-3.0
import "C"

//-----------------------------------------------------------------------
// GtkStack
//  A stacking container.
//
// Functions:
// - GtkWidget * 	gtk_stack_new ()
// - void 	gtk_stack_add_named ()
// - void 	gtk_stack_add_titled ()
// - GtkWidget * 	gtk_stack_get_child_by_name ()
// - void 	gtk_stack_set_visible_child ()
// - GtkWidget * 	gtk_stack_get_visible_child ()
// - void 	gtk_stack_set_visible_child_name ()
// - const gchar * 	gtk_stack_get_visible_child_name ()
// - void 	gtk_stack_set_visible_child_full ()
// - void 	gtk_stack_set_homogeneous ()
// - gboolean 	gtk_stack_get_homogeneous ()
// - void 	gtk_stack_set_hhomogeneous ()
// - gboolean 	gtk_stack_get_hhomogeneous ()
// - void 	gtk_stack_set_vhomogeneous ()
// - gboolean 	gtk_stack_get_vhomogeneous ()
// - void 	gtk_stack_set_transition_duration ()
// - guint 	gtk_stack_get_transition_duration ()
// - void 	gtk_stack_set_transition_type ()
// - GtkStackTransitionType 	gtk_stack_get_transition_type ()
// - gboolean 	gtk_stack_get_transition_running ()
// - gboolean 	gtk_stack_get_interpolate_size ()
// - void 	gtk_stack_set_interpolate_size ()
//-----------------------------------------------------------------------

