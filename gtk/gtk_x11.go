// +build with-x11

package gtk

import (
	"unsafe"

	"github.com/zurek87/go-gtk3/gdk"
)

func (v *Window) XID() int32 {
	return gdk.WindowFromUnsafe(unsafe.Pointer(v.GWidget.window)).GetNativeWindowID()
}
