package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go is running on ")

	//Like if statement also in switch we can have short init statements

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Linux os")

	case "darwin":
		fmt.Println("Darwin os")

	case "windows":
		fmt.Println("Windows os")

	case "mac":
		fmt.Println("Mac os")

	default:
		fmt.Printf("%s os", os)

	}

}
