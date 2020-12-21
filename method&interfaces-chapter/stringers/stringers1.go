package main

import "fmt"

//Declare new custom types
type Plane2D [2]byte
type Space3D [3]Plane2D

//Make new custom types eligible for Stringer interface
func (p Plane2D) String() string {
	return fmt.Sprintf("P-unit : ( %v , %v )", p[0], p[1])
}

func (s Space3D) String() string {
	return fmt.Sprintf("S-Unit : [ %v , %v , %v ]", s[0], s[1], s[2])
}

func main() {
	var p1 = Plane2D{1, 1} //Because Plane2D is array of bytes
	var p2 = Plane2D{2, 2}
	var p3 = Plane2D{1, 2}

	var s = Space3D{p1, p2, p3}

	//Call exported function Println(a ...interface{})
	fmt.Println(s)
}
