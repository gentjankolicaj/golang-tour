package main

import "fmt"

func main() {
	var (
		f = -10.2
		i = int(f)
		u = uint(f)
	)

	fmt.Println("f ", f)
	fmt.Println("i ", i)
	fmt.Println("u ", u)
}
