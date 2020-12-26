package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		c <- x

		x, y = y, x+y //=> x=y & y=x+y
	}

	close(c) //close channel to inform receiver that there are no more values coming
}
func main() {
	ch := make(chan int, 100)

	go fibonacci(cap(ch), ch) //start execution of fibonacci in another goroutine

	for i := range ch { //Read values from channel
		fmt.Println(i)
	}
}
