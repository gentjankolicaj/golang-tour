package main

import "fmt"

//Declaration
//Var a=type
//Var a type
//a:=type short declaration

var slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func main() {

	//Note: range keyword is applied on {arrays,slices,maps} and returns 2 values {index & element}
	for i, v := range slice {
		fmt.Println("Index=", i, " - ", " element=", v)
	}

}
