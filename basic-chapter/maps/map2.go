package main

import "fmt"

type Health struct {
	height float64
	weight float64
}

//healthMap => key of type int & value of type Health
//Below is continued map literal (harte e vazhduar e shkruar)
var healthMap = map[int]Health{
	1: {170, 87},
	2: {174, 85},
	3: {180, 97},
	4: {178, 80},
}

type vertex struct {
	lat, lon float64
}

var vertexMap = map[string]vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println("vertex ->", vertexMap, "\n")
	fmt.Println("health ->", healthMap)
}
