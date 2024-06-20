package main

import (
	"fmt"
	"math"
)

//NOTE: below I have written some templates of how functions me appear in go

//Names
func Exposed()   {}
func unExposed() {}

//Params + returns
func a(x int, y int) int           { return x + y }
func b(x, y int) float64           { return float64(x + y) }
func c(x, y, z int) (float64, int) { return float64(x + z), y }

//Function values
func d(fx func(int, int) (int, int)) float64 {
	i, j := fx(90, 11)    //Get both returned values
	return float64(j + i) //Add both values and covert to float64
}

func e() float64 {

	//Function declaration and variable assignment
	fv := func(x int, y int) float64 {
		return math.Sqrt(float64(x*x) + float64(y*y))
	}

	var value float64
	value = fv(3, 4)

	fmt.Println("Value of square root is ", value)
	return value
}

//Function closures //returning a functions

func f() func(float64) float64 {
	sum := 0.0
	return func(x float64) float64 {
		sum = sum + x
		return sum
	}
}
