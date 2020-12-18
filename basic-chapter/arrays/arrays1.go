package main

import "fmt"

type object struct {
	address uint
	name    string
}

var (
	a [10]int
	b [10]float64
	c [10]uint
	d [10]complex128
	e [10]string
	o [10]object
	p [10]*object // variable p,array of 10 elements,element types : object pointer
)

func main() {

	//For loop to populate arrays
	for i := 0; i < 10; i++ {
		a[i] = i
		b[i] = float64(i) + 1.2
		c[i] = uint(i) + 5
		d[i] = complex(float64(i), 3.3)
		e[i] = string(i + 70)
		o[i] = object{address: uint(i) + 200, name: "Testing "}
		p[i] = &o[i]

	}

	fmt.Println("Int array ", a)
	fmt.Println("float array ", b)
	fmt.Println("uint array ", c)
	fmt.Println("complex array ", d)
	fmt.Println("string array ", e)
	fmt.Println("struct array ", o)
	fmt.Println("struct pointer array ", p) //Array that contain pointers ,specifically structure pointers

}
