package main

import "fmt"

//Note: template for type assertion is :
// t:=i.(Type)
// var t=i.(Type)
// var t Type=i.(Type)
// t,flag=i.(Type) ->This template is more safe because it avoids panic because of second variable 'flag'
func main() {
	var i interface{} = "Hello world and type assertion"

	s := i.(string) //type assertion
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64) //return : zero of type float64 and false
	fmt.Println(f, ok)

	f = i.(float64) //Causes panic because concrete type of i is string not float64
	fmt.Println(f)
}
