package main

import (
	"fmt"
	"math"
)

type MyInterface interface {
	M()
}

type MyType float64

type MyStruct struct {
	S string
}

func (t MyType) M() {
	fmt.Println("MyFloat impl of M() : ", t)
}

func (s MyStruct) M() {
	fmt.Println("MyStruct impl of M() : ", s)
}

func main() {
	var i MyInterface

	i = MyStruct{"Hello world "}
	Describe(&i)
	i.M() //Calls method func (s MyStruct) M()

	i = MyType(math.Pi)
	Describe(&i)
	i.M() //Calls method func (t MyType) M()

}

func Describe(i *MyInterface) {
	fmt.Printf("(%v , %T) \n", i, i)
}
