package main

import "fmt"

func main() {
	names := [4]string{"Hello ", "World", "Author", "Version"}

	fmt.Println(names)

	//Declare 2 slice variables
	//var a[] type or var a,b,c,d []type

	var a []string
	var b []string

	a = names[0:2]
	b = names[1:3]
	fmt.Println(a, b)

	b[0] = "Tim-doe-testing"
	fmt.Println("Slices a and b", a, b)
	fmt.Println("Underlying array 'names' of slices ", names)
}
