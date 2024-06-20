package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//Method has RECEIVER ARGUMENT
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//Function doesn't have RECEIVER ARGUMENT
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v0 := Vertex{3, 4} //short-declaration

	fmt.Println("Value of Abs() method ", v0.Abs())
}
