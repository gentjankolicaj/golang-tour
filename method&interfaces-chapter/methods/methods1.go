package main

import (
	"fmt"
	"math"
)

type RightTriangle struct {
	a, b, c float64
}

//Note when I declare a method of a type T in GO,that means that each instance of T has that method
//And i can call it from variables of type T.
func (t RightTriangle) calculateHypotenuse() float64 {
	var hypotenuse = math.Sqrt(t.a*t.a + t.b*t.b)
	t.c = hypotenuse
	return hypotenuse
}

func main() {

	t0 := RightTriangle{a: 3, b: 4}
	t1 := RightTriangle{a: 6, b: 8}

	fmt.Println("t0 variable ", t0.calculateHypotenuse())
	fmt.Println("t1 variable ", t1.calculateHypotenuse())
}
