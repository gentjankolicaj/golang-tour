package main

import (
	"fmt"
	"math"
)

type Vertex1 struct {
	x, y float64
}

//Declare functions
func abs(v Vertex1) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

//Note functions below are different
//function 0 changes the values of fields where v is pointing
//function 1 changes the value of fields of v
func scale0(v *Vertex1, f float64) {
	v.y = v.y * f
	v.x = v.x * f
}

func scale1(v Vertex1, f float64) {
	v.y = v.y * f
	v.x = v.x * f
}

func main() {
	v := Vertex1{3, 4}

	scale1(v, 10) //Because a copy of v is passed, values of actual v are not changed
	fmt.Println("After scale1 function call ", v)

	scale0(&v, 10) //Because pointer of v is passed ,values of actual v are changed to {3*10,4*10}
	fmt.Println("After scale0 function call ", v)

}
