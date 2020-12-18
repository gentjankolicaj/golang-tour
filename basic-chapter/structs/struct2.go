package main

import "fmt"

//Composed structure => structures inside other structures
type student struct {
	personal_info personal_info
	adress_info   adress_info
}

type personal_info struct {
	firstname string
	lastname  string
}

type adress_info struct {
	street, city, country string
}

func main() {
	var personal personal_info = personal_info{"Gentjan", "Kolicaj"}
	var adress adress_info = adress_info{street: "Kadri kullakshi", city: "Tirane", country: "Albania"}

	var firstStudent = buildStudent(personal, adress)
	fmt.Println("First composed structure ", firstStudent)

	//Long way of creating an composed struct

	var secondStudent student = student{personal_info: personal_info{"Tim", "Doe"}, adress_info: adress_info{street: "Time square", country: "New york"}}
	fmt.Println("Second composed structure ", secondStudent)

	thirdStudent := student{personal_info: personal_info{"John", "Die"}, adress_info: adress_info{street: "Time square", country: "California"}}
	fmt.Println("third composed structure ", thirdStudent)

}

func buildStudent(personal_info personal_info, adress_info adress_info) student {
	//declaration of student type var and initialization
	var localStudent student = student{personal_info: personal_info, adress_info: adress_info}
	return localStudent
}
