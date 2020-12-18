package main

import "fmt"

//j int=1 => int i=1 & int j=2
var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no"

	fmt.Println("Package level var initializers => ", i, j)
	fmt.Println("")
	fmt.Println("Function level var initializers =>", c, python, java)

}
