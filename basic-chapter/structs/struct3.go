package main

import "fmt"

type Vertex struct {
	x int
	y int
}

func main() {

	directFieldValueChange()
	indirectFieldValueChange()

}

func directFieldValueChange() {
	var v Vertex = Vertex{1, 2}
	v.x = 4
	fmt.Println("Direct value change of struct field x ", v.x, " - ", v)
}

func indirectFieldValueChange() {
	v := Vertex{10, 11}
	var p *int = &v.x
	fmt.Println("Value of x field through pointer is ", *p)

	//changing the value of x field of struct through indirect referencing
	*p = 99
	fmt.Println("Indirect value change of struct field x ", v.x, " - ", v)
}
