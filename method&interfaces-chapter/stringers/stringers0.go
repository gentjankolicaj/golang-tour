package main

import "fmt"

//Declare a set of fields (struct)
type Person struct {
	Name string
	Age  int
}

//To satisfy Stringer interface i have to implement String() string method
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"John Doe", 42}
	b := Person{"John Wick", 40}

	fmt.Println(a, b)
}
