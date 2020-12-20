package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

//New type originating float64
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//New type originating from struct
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	f := MyFloat(-math.Sqrt2)

	var a Abser

	a = f  //My float implements Abser interface
	a = &v //*Vertex implements Abser interface

	a = v
	fmt.Println(a.Abs())

}
