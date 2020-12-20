package main

import (
	"fmt"
	"math"
)

type Vertex1 struct {
	X, Y float64
}

//Declare function
func AbsFunc(v Vertex1) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//Declare method with value receiver
func (v Vertex1) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex1{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex1{6, 8}
	fmt.Println(p.Abs())     //Because p is a pointer we have indirect value receiver (Automatically by go (*p).Abs())
	fmt.Println(AbsFunc(*p)) //Because p is a pointer and we need a value variable
}
