package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

//Implement method Error() string to make MyError eligible to error interface
func (e *MyError) Error() string {
	return fmt.Sprintf(" at %v , %s ", e.When, e.What)
}

//A kind of factory function
func run() error {
	return &MyError{When: time.Now(), What: "Something didn't work"}
}

func main() {

	if err := run(); err != nil {
		fmt.Println(err)
	}

}
