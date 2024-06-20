package main

import "fmt"

//Declare a new type IPAddr of [4]byte array
type IPAddr [4]byte

//In order to satisfy fmt.Stringer interface i have to impl String() string method
func (ip IPAddr) String() string {
	return fmt.Sprintf(" \"%v.%v.%v.%v\"", ip[0], ip[1], ip[2], ip[3])
}

func main() {

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for key, value := range hosts {
		fmt.Printf("Key : %v ,Value %v \n", key, value)
	}
}
