package main

import "fmt"

type Vertex0 struct {
	X, Y float64
}

func (v *Vertex0) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex0, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex0{3, 4} //Variable v holds value
	v.Scale(2)         //Pointer indirection
	ScaleFunc(&v, 10)  //&v because v is of type Vertex value

	p := &Vertex0{4, 3} //Variable p holds pointer(mem address)
	p.Scale(3)
	ScaleFunc(p, 8) //p and not &p because p is of type Vertex pointer

	fmt.Println(v, p)
}
