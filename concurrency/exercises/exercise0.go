package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

//todo: to implement functions
func Walk(t *tree.Tree, ch chan int) {

}

func Same(t1, t2 *tree.Tree) bool {
	return true
}

func main() {
	//	ch:=make(chan int)

	//go Walk(tree.New(1),ch)

	t := tree.New(10)
	fmt.Println(t)
}
