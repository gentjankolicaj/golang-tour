package main

import "fmt"

func add0(x int, y int) int {
	return x + y
}

func add1(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("function add0 with all data-types : ", add0(0, 9))
	fmt.Println("Function add1 with continued form : ", add1(100, 1))

}
