package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var u uint = uint(f)

	fmt.Println("x : ", x)
	fmt.Println("y : ", y)
	fmt.Println("f : ", f)
	fmt.Println("u : ", u)
}
