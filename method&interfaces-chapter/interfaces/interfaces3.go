package main

import "fmt"

type I0 interface {
	M0()
}

type T0 struct {
	S string
}

func (t *T0) M0() {
	if t == nil {
		fmt.Println("<nil")
		return
	}
	fmt.Println(t.S)
}

func describe0(i I0) {
	fmt.Printf("(%v , %T ", i, i)
}

func main() {
	var i I0
	var t *T0

	i = t
	describe0(i)
	i.M0() //Doesn't throw NPE even though t is nil but i isn't nil because it has pointer of t

	i = &T0{"Hello world"} //C.E for i=T0 because only &T0 has implemented method
	describe0(i)
	i.M0()

}
