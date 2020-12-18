package main

import "fmt"

//NOTE: when we have multiple deferred functions than these functions are placed into a stack
//After the container function is returned then all deferred functions are executed in LIFO order

func main() {
	fmt.Println("Counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done")
}
