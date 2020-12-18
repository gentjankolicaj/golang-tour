package main

import "fmt"

//When we declare variables in package level those variables are visible
//to all files is same directory (package level)
var g = 19         //type int implicit type placement by compiler
var k = float64(g) //type conversion from int=>float64
var d = 12.23      //type float32/64 based on system
var u = uint(d)    //type conversion from float to unit(unsigned integer 0-2^32-1)

func main() {
	fmt.Println("g ", g)
	fmt.Println("k ", k)
	fmt.Println("d ", d)
	fmt.Println("u ", u)

}
