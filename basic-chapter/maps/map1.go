package main

import "fmt"

type Cord struct {
	lat, lon float64
}

var arr = []Cord{Cord{1, 2}, {12, 1}, {28, 2}} //Cord struct array literal

var myMap = map[string]Cord{
	"Bell Labs": Cord{
		40.68433, -74.39967,
	},
	"Google": Cord{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println("Map literal ", myMap)
	fmt.Println("Array literal ", arr)
}
