package main

import "fmt"

var (
	a int       = 1
	b uint      = 1
	c float64   = 1.0
	z complex64 = 1 + 2i
)

func main() {
	fmt.Println("Printing actual values ", a, b, c, z)

	fmt.Println("")
	fmt.Println("Printing memory addresses ", &a, &b, &c, &z)

}
