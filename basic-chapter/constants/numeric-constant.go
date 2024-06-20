package main

import "fmt"

const ( ///with ( ) are created blocks in golang

	BigValue = 1 << 100

	SmallValue = 1 >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(y float64) float64 {
	return y * 0.1
}

func main() {

	fmt.Println(needInt(SmallValue))
	fmt.Println(needFloat(SmallValue))

	fmt.Println(needFloat(BigValue))
	//	fmt.Println(needInt(BigValue)) overflows int because max value of int is 2^64 and BigValue=2^100

}
