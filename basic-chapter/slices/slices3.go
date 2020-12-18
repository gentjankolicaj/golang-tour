package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("Slice q ", q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println("Slice r ", r)

	//Short struct variable declaration
	//Anonymous/local struct declaration
	s := []struct {
		i int
		x bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println("Slice of structs ", s)
}
