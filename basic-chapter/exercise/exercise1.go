package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var slice2D = make([][]uint8, dx)

	for i := range slice2D {
		slice2D[i] = make([]uint8, dy)
	}

	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			slice2D[i][j] = applyFunc(i, j, 1)
		}
	}

	return slice2D

}

func main() {
	pic.Show(Pic)
}

func applyFunc(x, y, funcNr int) uint8 {
	if funcNr == 1 {
		return uint8(x * y)
	} else if funcNr == 2 {
		return uint8((x + y) / 2)
	} else {
		return uint8(x ^ y)
	}
}
