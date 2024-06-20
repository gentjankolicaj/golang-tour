package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

//Declare a new type of color and impl RGBA method to make it color.Color eligible
type MyColor struct {
	r, g, b, a uint32
}

func (c MyColor) RGBA() (r, g, b, a uint32) {
	r = c.r
	g = c.g
	b = c.b
	a = c.a
	return
}

//Declare a new type of image and implement methods to make it image.Image eligible
type MyImage struct {
	length int
	width  int
	Rect   image.Rectangle
	Color  MyColor
}

func (i MyImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (i MyImage) Bounds() image.Rectangle {
	i.Rect = image.Rect(0, 0, i.width, i.length)
	return i.Rect
}

func (i MyImage) At(x, y int) color.Color {
	return color.Gray{Y: 122}
}

func main() {
	m := MyImage{length: 100, width: 100, Color: MyColor{48, 255, 109, 1}}
	pic.ShowImage(m)
}
