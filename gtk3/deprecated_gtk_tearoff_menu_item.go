package gtk3

// #include "gtk3.go.h"
// #cgo pkg-config: gtk+-3.0
import "C"

//-----------------------------------------------------------------------
// GtkTearoffMenuItem
//
// GtkTearoffMenuItem is deprecated and should not be used in newly written code.
// Menus are not meant to be torn around.
//-----------------------------------------------------------------------


// deprecated
type TearoffMenuItem struct {
	MenuItem
}

// deprecated
func NewTearoffMenuItem() *TearoffMenuItem {
	return &TearoffMenuItem{*newMenuItemInternal(
		C.gtk_tearoff_menu_item_new())}
}

