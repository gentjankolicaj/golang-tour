package main

import (
	"fmt"
	"math/rand"
)

type Salary struct {
	empId int
	base  float64
	bonus float64
}

//Declare a method to raise salary
func (s *Salary) raise(percentage float64) {
	s.base = s.base + (s.base * percentage / 100)

}

func (s *Salary) assignBonus(bonus float64) {
	s.bonus = s.bonus + bonus
}

//Declare a function to print all details
func printBeforeRaise(s *Salary) {
	fmt.Printf("Before: EmpId : %d and salary details %v \n", s.empId, s)
}
func printAfterRaise(s *Salary) {
	fmt.Printf("After: EmpId : %d and salary details %v \n", s.empId, s)
}

func generateRandomSalary(nr int) *[]Salary {
	var slice = make([]Salary, nr)
	for i := 0; i < nr; i++ {
		slice[i] = Salary{empId: i, base: float64(rand.Intn(1000)), bonus: 0}
	}
	return &slice
}

func main() {
	nr := 15
	fmt.Println("Generating 15 random salaries")

	//Declare slice pointer
	var slicePointer *[]Salary
	slicePointer = generateRandomSalary(nr)

	//Declare slice value because I couldn't iterate over slice pointer
	var sliceValue []Salary
	sliceValue = *slicePointer

	for _, v := range sliceValue {
		printBeforeRaise(&v)
		v.raise(10)       //same as &v.raise(10) because of pointer indirection
		v.assignBonus(10) //same as &v.assignBonus(10) because of pointer indirection
		printAfterRaise(&v)

	}
}
