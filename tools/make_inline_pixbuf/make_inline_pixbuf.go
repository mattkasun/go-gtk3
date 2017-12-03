package main

import (
	"fmt"
	"os"

	"github.com/zurek87/go-gtk3/gdkpixbuf"
	"github.com/zurek87/go-gtk3/gtk3"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: make-inline-pixbuf resourceName input > output\n")
		os.Exit(1)
	}

	image := gtk3.NewImageFromFile(os.Args[2])
	pb := image.GetPixbuf()
	if pb.GetWidth() == -1 {
		fmt.Fprintf(os.Stderr, "ERROR: invalid pixbuf image\n")
		os.Exit(2)
	}

	var pbd gdkpixbuf.PixbufData
	pbd.Data = pb.GetPixelsWithLength()
	pbd.Width, pbd.Height, pbd.RowStride, pbd.HasAlpha = pb.GetWidth(), pb.GetHeight(), pb.GetRowstride(), pb.GetHasAlpha()
	pbd.Colorspace, pbd.BitsPerSample = pb.GetColorspace(), pb.GetBitsPerSample()

	fmt.Printf("package main \n\nimport \"github.com/mattn/go-gtk/gdkpixbuf\"\n\nvar (\n")
	fmt.Printf("\t%s = %#v\n", os.Args[1], pbd)

	fmt.Println(")\n")
}
