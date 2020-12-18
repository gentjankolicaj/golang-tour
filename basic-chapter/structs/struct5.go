package main

import "fmt"

type Wall struct {
	x, y int
}

//Package-Level variable declaration
//Note: When we have var keyword we can't use := short variable declaration
var (
	a = Wall{1, 2}
	b = Wall{x: 15}
	c = Wall{}
	d = &Wall{212, 121}
)

func main() {
	fmt.Println(a, b, c, d)
}
