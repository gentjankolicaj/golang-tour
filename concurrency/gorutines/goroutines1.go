package main

import (
	"errors"
	"fmt"
	"time"
)

type student struct {
	id            int64
	name, surname string
	age           int
}

func (s *student) String() string {
	return fmt.Sprintf("type:student | id=%v,name=%v,surname=%v,age=%v", s.id, s.name, s.surname, s.age)
}

//Student dao
type studentDao interface {
	getAllStudents() (*[]student, error)
	getStudent(id int64) (*student, error)
	editStudent(new, old *student) (*student, error)
	deleteStudent(s *student) (bool, error)
}

type studentDaoImpl struct{}

func (sdi studentDaoImpl) getAllStudents() (*[]student, error) {
	var amount int = 21
	var slice []student = make([]student, amount)
	for i := 1; i <= 21; i++ {
		sp, err := sdi.getStudent(int64(i))

		if err != nil {
			return nil, errors.New(err.Error())
		}
		slice[i-1] = *sp

	}
	return &slice, nil
}

func (sdi studentDaoImpl) getStudent(id int64) (*student, error) {
	var s *student = &student{id: id, name: "Name ", surname: "Surname ", age: int(id + 10)}
	return s, nil
}

func (sdi studentDaoImpl) editStudent(new, old *student) (*student, error) {
	old.id = new.id
	old.name = new.name
	old.surname = new.surname
	old.age = new.age
	return old, nil
}

func (sdi studentDaoImpl) deleteStudent(s *student) (bool, error) {
	s = nil
	if s == nil {
		return true, nil
	}
	return false, errors.New("Couldn't make pointer student nil")
}

func testStudentDao(amount int) {
	var studentDao studentDao
	studentDao = studentDaoImpl{}

	//Read all students
	//Return type is slice pointer and error value
	slicePointer, err := studentDao.getAllStudents()
	if err == nil {
		for _, v := range *slicePointer {
			fmt.Println(v)
		}
	}

	//Get specific student
	st, err := studentDao.getStudent(int64(99))
	fmt.Println(st)

	//Edit specific student
	var newStudent student
	newStudent = student{id: 101, name: "Name 101", surname: "New student ", age: int(-101)}
	ust, err := studentDao.editStudent(&newStudent, st)

	if err == nil {
		fmt.Println(ust)
	}

	//Delete student
	flag, err := studentDao.deleteStudent(ust)
	if err == nil {
		fmt.Println("Student : ", *ust, " deleted status ", flag)
	}
}

func main() {

	//Calling method in a goroutine
	fmt.Println("Starting additional goroutine : ")
	go testStudentDao(13)

	//Set a sleep duration in main goroutine
	fmt.Println("Main-goroutine set to sleep 30 seconds.")
	time.Sleep(30 * time.Second)
	fmt.Println("Starting main goroutine : ")
	testStudentDao(7)
}
