package main

import "fmt"

type I1 interface {
	M1()
}

type T1 struct {
	S string
}

func main() {
	var i I1
	describe1(i)
	i.M1() //Causes runtime error because i==nil
	// and  there is no interface tuple (value,type) to indicate concrete method to call
}

func describe1(i I1) {
	fmt.Printf("(%v, %T) ", i, i)
}
