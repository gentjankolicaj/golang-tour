package main

import "fmt"

func main() {
	var s []int //Nil slice => len(s)==0 && cap(s)==0
	printMySlice(s)

	s = append(s, 23, 12, 32)
	printMySlice(s)

	s = append(s, 1, 0, 9)
	printMySlice(s)

	s = append(s, 8, 8, 8)
	printMySlice(s)
}

func printMySlice(s []int) {
	fmt.Printf("len=%d cap=%d -- %v \n", len(s), cap(s), s)
}
