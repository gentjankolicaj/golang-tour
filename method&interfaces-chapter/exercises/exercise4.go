package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type MyImage struct {
	length int
	width  int
	Rect   image.Rectangle
}

func (i MyImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (i MyImage) Bounds() image.Rectangle {
	i.Rect = image.Rect(0, 0, i.width, i.length)
	return i.Rect
}

func (i MyImage) At(x, y int) color.Color {
	return color.Gray{10}
}

func main() {
	m := MyImage{length: 100, width: 100}
	pic.ShowImage(m)
}
