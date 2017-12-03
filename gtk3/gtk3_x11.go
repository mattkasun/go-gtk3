// +build with-x11

package gtk3

import (
	"unsafe"

	"github.com/zurek87/go-gtk3/gdk3"
)

func (v *Window) XID() int32 {
	return gdk3.WindowFromUnsafe(unsafe.Pointer(v.GWidget.window)).GetNativeWindowID()
}
