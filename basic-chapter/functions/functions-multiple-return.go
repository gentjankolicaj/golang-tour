package main

import "fmt"

func process(x int, y int, z float32) (int, string) {
	fmt.Println("This is float argument value ", z)
	return x + y, "func process second return value !"
}

func main() {

	value, message := process(12, 12, 123)

	fmt.Println("This is x+y value ", value)
	fmt.Println("Message value ", message)

}
