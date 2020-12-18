package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))

	if 10 > 9 {
		fmt.Println("10 is bigger than 9")
	}
	//1>>10  1/2^10 vlera qe rezulton
	//1<<10 1*2^10 pra zhvendos 1-in 10 pozicione ne sistem me baze 2
}
