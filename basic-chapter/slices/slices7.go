package main

import "fmt"

func main() {

	//Create a slice with length 5 and capacity 5
	//
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2] //0,1 indexes retrieved
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

}

func printSlice(t string, s []int) {
	fmt.Printf("%s len=%d cap=%d %v \n", t, len(s), cap(s), s)
}
