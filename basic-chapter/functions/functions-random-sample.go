package main

import (
	"fmt"
	"math/rand"
	"syscall"
)

func main() {

	fmt.Println("My favourite number is ", rand.Intn(10))
	rand.Seed(syscall.SYS_TIME)
	fmt.Println("Random number with system calls ", rand.Intn(10))
}
