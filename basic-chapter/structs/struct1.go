package main

import "fmt"

type vertex struct {
	x    int
	y    int
	w, z float64
}

func main() {

	fmt.Println("Strukture me fusha pjeserisht te initicializuara ", vertex{x: 1, y: 1})
	fmt.Println("Strukture me fusha te pa inicializuara ", vertex{})

}
