package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	var value int

	value = add(10000, 101)

	fmt.Println(value)
}
