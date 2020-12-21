package main

import "fmt"

//Declare a type-switch inside this function
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Variable i holds a value: %v of type: %T \n", v, v)
	case string:
		fmt.Printf("Variable i holds a value: %v of type: %T \n", v, v)
	case float64:
		fmt.Printf("Variable i holds a value: %v of type: %T \n", v, v)
	default:
		fmt.Printf("Variable i holds a value type of which is not validated in type-switch \n")
	}
}

func main() {
	do(10)
	do("Hello world")
	do(true)
	do(2323.23)
}
