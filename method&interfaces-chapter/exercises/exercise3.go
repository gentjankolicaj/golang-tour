package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (n int, err error) {
	var slc = make([]byte, 13)

	tmpNr, tmpErr := r.r.Read(slc)
	n = tmpNr
	err = tmpErr

	//Decrypt slc
	decryptSlice(&slc)

	return

}

//todo: Decryption method to be looked because I happen to not know algorithm ROT13
func decryptSlice(s *[]byte) {
	valueSlice := *s
	for i, v := range valueSlice {
		valueSlice[i] = v - byte(13)
	}
	fmt.Println("Decrypted slice : ", string(valueSlice))
}

func main() {
	//s is pointer variable of type Reader
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{r: s}
	io.Copy(os.Stdout, &r) //&r gives address of value not value because value is going to be duplicated
}
