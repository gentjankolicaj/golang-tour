package main

import (
	"fmt"
	"math"
)

type Vertex2 struct {
	X, Y float64
}

//Method type : Pointer receiver
func (v *Vertex2) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//Method type : Pointer receiver
func (v *Vertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex2{3, 4}
	fmt.Printf("Before scaling : %+v , Abs: %v \n", v, v.Abs())

	v.Scale(20)
	fmt.Printf("After scaling : %+v , Abs: %v \n", v, v.Abs())
}
