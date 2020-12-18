package main

import "fmt"

func main() {
	slice := make([]int, 10) //len=10 and cap=10

	for i := range slice {
		slice[i] = 1 << uint(i) // == 2**i ,for every i it moves 1 i positions to the left=> 1*2^i
	}

	for _, v := range slice {
		fmt.Printf("%d \n", v)
	}

	//Other range forms are
	//1. for i,v:=range slice index+value
	//2. for _,v:=range slice no-index+value
	//3. for i,_:=range slice index+no-value
	//4. for i:=range slice   index
}
