package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

//Method type : VALUE RECEIVER
func (v Vertex) abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

//Method type : POINTER RECEIVER
func (v *Vertex) scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func main() {
	v := Vertex{3, 4}

	v.scale(10) //Call method Scale on variable v

	fmt.Println(v.abs()) //Call method Abs() on a copy value of v variable
}
