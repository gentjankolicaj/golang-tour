package main

import "fmt"

func varDeclarationBlock() {

	//Var declaration block
	var (
		i int
		j uint
		f float64
		b bool
		z complex128
		s string
	)

	fmt.Printf("i value %v \n", i)
	fmt.Printf("j value %v \n", j)
	fmt.Printf("f value %v \n", f)
	fmt.Printf("b value %v \n", b)
	fmt.Printf("z value %v \n", z)
	fmt.Printf("s value %v \n", s)
}

func main() {
	var i int
	var f float32
	var b bool
	var s string

	fmt.Printf("i value %v \n", i)
	fmt.Printf("f value %v \n", f)
	fmt.Printf("b value %v \n", b)
	fmt.Printf("s value %v \n", s)

	fmt.Println("")
	fmt.Println("")

	varDeclarationBlock()

}
