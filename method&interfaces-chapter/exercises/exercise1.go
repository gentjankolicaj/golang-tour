package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

//Make ErrNegativeSqrt  "error" eligible
//Note: A call to fmt.Sprint(e) inside the Error method will send the program into an infinite loop.
//You can avoid this by converting e first: fmt.Sprint(float64(e))
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(Sqrt(2))

	fmt.Println(Sqrt(-2))
}
