package main

import "fmt"

func fibonacci() func() int {

	return func() int {
		return 0
	}
}

func main() {
	fv := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(fv())
	}

}
