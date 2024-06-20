package variables

import (
	"errors"
	"fmt"
	"log"
)

func BuildHello(name string) (int, string, error) {
	if name == "" {
		return 0, "", errors.New("Parameter 'name' is empty ")
	}
	return 0, name, nil
}

func SayHello(name string) string {
	code, name, err := BuildHello(name)

	msg0 := fmt.Sprintf("Return code %d ", code)
	msg1 := fmt.Sprintf("Return name %s ", name)

	fmt.Println(msg0)
	fmt.Println(msg1)
	if err != nil {
		log.Fatal(err)
	}

	return name
}
