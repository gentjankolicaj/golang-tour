package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	//Declare a byte slice []byte
	b := make([]byte, 8)

	//for=>forever
	for {
		n, err := r.Read(b) //Populate b slice with 8 bytes
		fmt.Printf("n=%v , err=%v , b=%v \n", n, err, b)
		fmt.Printf("b[:n]=%q \n", b[:n])

		if err == io.EOF {
			break
		}
	}
}
