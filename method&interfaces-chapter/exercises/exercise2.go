package main

import (
	"fmt"
	"golang.org/x/tour/reader"
)

//ASCII american standard code for information interchange
type MyReader struct{}

//Implement method Read(b [] byte) (n int,err error)
func (r MyReader) Read(b []byte) (n int, err error) {
	var aAscii = uint8('A')
	var slcLen, slcCap = len(b), cap(b)

	if slcLen == 0 || slcCap == 0 {
		n = 0
		err = &MyError{"Slice length or capacity is 0"}
		return
	}
	for i := range b {
		b[i] = aAscii
		n = i
	}

	n = n + 1
	err = nil
	return
}

//Custom error
type MyError struct {
	message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("MyError: message { %s } ", e.message)
}

func main() {
	reader.Validate(MyReader{})
}
