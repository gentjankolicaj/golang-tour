package main

import "fmt"

//Interface definition
type I interface {
	M()
}

//Struct definition
type T struct {
	S string
}

//Method declaration & interface impl because :
//1. (t T) => value receiver of T
//2. M() => same method sign as interface I

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{S: "Hello world of GO"}
	i.M()
}
