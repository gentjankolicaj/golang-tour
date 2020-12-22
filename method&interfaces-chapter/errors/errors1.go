package main

import "fmt"

type RuntimeError struct {
	code    int
	message string
}

//Make RuntimeError eligible for "error interface" by implementing  Error()string  method
func (r *RuntimeError) Error() string {
	return fmt.Sprintf("RuntimeError: code { %v } , message { %s} ", r.code, r.message)
}

//Create a factory function
func getError(code int, message string) *RuntimeError {
	return &RuntimeError{code: code, message: message}
}

func calculate(x, y int) (float64, error) {
	if x == 0 || y == 0 {
		return 0, getError(0, "One of params has value 0.")
	}
	return float64(x + y), nil
}

func main() {
	_, err1 := calculate(1, 5)

	if err1 != nil {
		fmt.Println("Calling of method calculate(1,5) caused error : ", err1)
	}

	_, err2 := calculate(2, 0)
	if err2 != nil {
		fmt.Println("Calling of method calculate(2,0) caused error : ", err2)
	}
}
