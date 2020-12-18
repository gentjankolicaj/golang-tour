package main

import "fmt"

const e, pi = 1.2, 23 //i kam deklaruar me lowercase pasi po ti deklaroj me uppercase behen exported ne pakete
//Dhe duke qene se jane deklaruar me perpara atehere kemi compile error

var x, y = 0, 1

func main() {

	fmt.Printf("e value ", e)
	fmt.Println("p value ", pi)
	fmt.Println("")
	fmt.Println("x value ", x, " y value ", y)
	fmt.Println("")

	const a int = 1
	const b float64 = 99
	fmt.Println("a value ", a, " b value ", b)

	const f, d float64 = 11.1, 12
	fmt.Println("f value ", f, "d value ", d)

}
