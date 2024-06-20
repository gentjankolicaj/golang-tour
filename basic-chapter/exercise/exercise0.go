package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	var z0 float64 = 0
	var z1 float64 = 1
	var counter = 0

	for canLoop(z0, z1, counter) {
		counter++
		z0 = z1
		z1 = z0 - ((z0*z0 - x) / (2 * z0))
		fmt.Println(counter, " - ", z1)
	}

	return z1
}

func canLoop(z0 float64, z1 float64, counter int) bool {
	if counter != 0 {
		if (z1 - z0) > 0.00001 {
			return true
		} else {
			return false
		}
	} else {
		return true
	}

}

func main() {

	Sqrt(4)
}
