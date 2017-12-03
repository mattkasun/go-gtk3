package gtk3

// #include "gtk3.go.h"
// #cgo pkg-config: gtk+-3.0
import "C"

//-----------------------------------------------------------------------
// GtkFixed
//  A container which allows you to position widgets at fixed coordinates.
//
// Functions:
// GtkWidget * 	gtk_fixed_new ()
// void 	gtk_fixed_put ()
// void 	gtk_fixed_move ()
//-----------------------------------------------------------------------

type Fixed struct {
	Container
}

func NewFixed() *Fixed {
	return &Fixed{Container{Widget{C.gtk_fixed_new()}}}
}

func (v *Fixed) Put(w IWidget, x, y int) {
	C.gtk_fixed_put(FIXED(v), ToNative(w), gint(x), gint(y))
}

func (v *Fixed) Move(w IWidget, x, y int) {
	C.gtk_fixed_move(FIXED(v), ToNative(w), gint(x), gint(y))
}