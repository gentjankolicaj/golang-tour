package main

import "fmt"

//Stucts is a collection of fields of different/same types
//Structs have this template
/**
type Struct_Name struct{
   x int
   y int
   w,z float64
}
*/

type vertex0 struct {
	x int
	y int
}

func main() {
	fmt.Println(vertex0{26, 23})
}
