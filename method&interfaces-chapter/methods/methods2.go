package main

import (
	"fmt"
	"math"
)

type MyFloat float64
type MyFlag bool

func (f MyFloat) Abs() float64 {
	return 3.145
}

func (f MyFlag) waveFlag() {
	if f == true {
		fmt.Println("Raise flag !!!")
	} else {
		fmt.Println("Lower flag...")
	}
}

func main() {
	fl := MyFloat(-math.Sqrt2)
	fmt.Println(fl.Abs())

	fg := MyFlag(true)
	fmt.Printf("Other type %T \n", fg)
	fg.waveFlag()
}
