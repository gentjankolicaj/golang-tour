package main

import "fmt"

type Rectangle struct {
	a, b int
}

func main() {
	var r Rectangle = Rectangle{b: 99, a: 30}

	//Declare  a variable of type rectangle pointer
	//Pointer variables are declared with
	//1. var a *Datatype
	//2. a=&Datatype
	var p *Rectangle
	p = &r
	fmt.Println("Initial struct values through pointer var", *p)

	//I change value of struct field with pointer
	p.b = 44
	fmt.Println("Actual struct value after change ", r)

	//Short declaration of another pointer
	//1.Direct value of m variable is value of p
	//2.Value of p is memory address of v
	//3.To print values of v we need *m to point directly at value
	m := p
	fmt.Println("Short declaration of another pointer var 'm' ", m)

}
