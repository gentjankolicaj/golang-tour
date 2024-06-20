package main

import "fmt"

func main() {

	var sum float64
	var counter int
	var total int = 50

	for counter < total { //in this for loop init an post statement are not present
		counter = counter + 1
		sum = sum + float64(counter)
		fmt.Println("idx ", counter)
	}

	fmt.Println("Total sum ", sum)
}
