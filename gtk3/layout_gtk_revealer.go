package gtk3

// #include "gtk3.go.h"
// #cgo pkg-config: gtk+-3.0
import "C"

//-----------------------------------------------------------------------
// GtkRevealer
//  Hide and show with animation.
//
// Functions:
// - GtkWidget * 	gtk_revealer_new ()
// - gboolean 	gtk_revealer_get_reveal_child ()
// - void 	gtk_revealer_set_reveal_child ()
// - gboolean 	gtk_revealer_get_child_revealed ()
// - guint 	gtk_revealer_get_transition_duration ()
// - void 	gtk_revealer_set_transition_duration ()
// - GtkRevealerTransitionType 	gtk_revealer_get_transition_type ()
// - void 	gtk_revealer_set_transition_type ()
//-----------------------------------------------------------------------

