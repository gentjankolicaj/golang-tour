package main

import "fmt"

func main() {
	const n = 10
	const m = 10

	var arr [n][m]int
	fmt.Println("Starting to populate NxM array", n, m)

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {

			if j == i {
				arr[i][j] = 1
			} else {
				arr[i][j] = 0
			}
		}
	}

	for k := 0; k < n; k++ {
		fmt.Println()
		for l := 0; l < m; l++ {
			fmt.Print(arr[k][l])
		}
	}
	fmt.Println("\nFinished printing 2 dimensional array")
}
