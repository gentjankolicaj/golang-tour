package main

import "fmt"

func main() {

	var i int //Note : var keyword is necessary when we declare a variable but whe we have
	//short assignment operator we can omit var keyword
	j := i //j int

	fmt.Printf("j is of type %T \n", j)

	u := 132
	z := 1 + i

	fmt.Printf("u is of type %T \n", u)
	fmt.Printf("z is of type %T \n", z)

}
