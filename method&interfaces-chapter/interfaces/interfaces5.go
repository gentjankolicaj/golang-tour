package main

import "fmt"

//Using EMPTY INTERFACE(0 methods) as a super/parent of anny type
// interface{}
func main() {

	//Declare a variable of an empty interface
	var i interface{}

	describe2(i)

	i = 42 //Assign a value of type int
	describe2(i)

	i = 10.5 //Assign a value of type float64
	describe2(i)

	i = "Hello empty interfaces 'interface{}' "
	describe2(i)
}

func describe2(i interface{}) {
	fmt.Printf("(%v , %T ) \n", i, i)
}
