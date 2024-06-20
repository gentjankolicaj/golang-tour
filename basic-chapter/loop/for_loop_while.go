package main

import "fmt"

func main() {
	sum := 1

	//For-while loop in go
	for sum < 1000 {
		sum = sum + sum
	}

	fmt.Println("Sum value ", sum)

}
