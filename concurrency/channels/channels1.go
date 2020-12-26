package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1 //channel receives 1
	ch <- 2 //channel receives 2
	ch <- 3 //channel receives 3

	fmt.Println(<-ch) //channel send
	fmt.Println(<-ch) //channel send
}
