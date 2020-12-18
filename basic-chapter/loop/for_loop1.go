package main

import "fmt"

func main() {
	var (
		i   int
		sum float64
	)

	for ; i < 30; i++ {
		sum += float64(i)
		fmt.Println("Index ", i)
	}

	fmt.Println(sum)
}
