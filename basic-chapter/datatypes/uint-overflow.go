package main

import (
	"fmt"
)

func main() {

	var i, j uint
	i = 10
	j = -10 //Uint overflow because j accepts 0 and positive

	fmt.Printf("I value %v", i)
	fmt.Printf("J value %v", j)

}
