package main

import "fmt"

type fieldSet struct {
	a bool
	b int
	c uint
	d float64
	e complex64
}

func getStructSlice(nr int) []fieldSet {
	var slice []fieldSet = make([]fieldSet, nr)
	for i := 0; i < nr; i++ {
		var tmpFieldSet fieldSet = fieldSet{a: true, b: i, c: uint(i), d: float64(i), e: complex(1.6, 2*float32(i))}
		slice[i] = tmpFieldSet
	}
	return slice
}

func printSliceDetails(slice []fieldSet) {
	fmt.Println("Slice details : ")
	for i, v := range slice {
		fmt.Printf("Idx: %d | ptr value: %v | value: %v \n", i, &v, v)
	}
}

func main() {
	structSlice := getStructSlice(20)

	printSliceDetails(structSlice)

}
