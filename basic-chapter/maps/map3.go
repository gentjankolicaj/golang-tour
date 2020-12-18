package main

import "fmt"

func main() {
	m := make(map[string]int)

	//Insert
	m["Answer"] = 42
	fmt.Println("After insert value ", m["Answer"])

	//Update
	m["Answer"] = 50
	fmt.Println("After update value ", m["Answer"])

	//Read
	var el = m["Answer"]
	fmt.Println("After read value ", el)

	//Delete
	delete(m, "Answer")
	fmt.Println("After delete value ", m["Answer"])

	//Read with 2 variables
	elem, flag := m["Answer"]
	fmt.Println("The element :", elem, "Is present ?", flag)
}
