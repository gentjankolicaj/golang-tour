package main

import "fmt"

var (
	i = 1
	f = float64(i)
	z = uint(f)
)

func main() {
	fmt.Println("Package variable declaration block and type-conversion")
	fmt.Println("i : ", i)
	fmt.Println("f : ", f)
	fmt.Println("z : ", z)

}
