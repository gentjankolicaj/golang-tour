package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 24 //vlera e pointerit p=24 do te thote qe i===24
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

}
