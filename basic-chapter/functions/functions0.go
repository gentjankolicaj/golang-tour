package main

import (
	"fmt"
	"math"
)

func compute(myFunction func(float64, float64) float64) float64 {
	return myFunction(3, 4)
}

func main() {

	//Declare an anonymous function impl //NOTE my term
	//This function has one difference from normal functions : has no name
	functionVariable := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	//By using function variable I call the function
	fmt.Println(functionVariable(5, 12))

	//Pas function variable as some normal parameter
	fmt.Println(compute(functionVariable))
	fmt.Println(compute(math.Pow))

}
