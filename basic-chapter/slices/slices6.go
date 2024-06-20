package main

import "fmt"

func main() {
	var s []int
	fmt.Printf("%v , len=%d cap=%d \n", s, len(s), cap(s))

	if s == nil {
		fmt.Println("S slice is nil.")
	}
}
