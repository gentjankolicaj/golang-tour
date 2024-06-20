package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu sync.Mutex
	mp map[string]int
}

//Implement SafeCounter methods with pointer receiver
func (s *SafeCounter) Inc(key string) {
	s.mu.Lock() //Lock access to s structure
	s.mp[key]++
	s.mu.Unlock() //Unlock access to s structure
}

func (s *SafeCounter) Value(key string) int {
	s.mu.Lock() //Lock access to s structure

	defer s.mu.Unlock() //To execute after Value method finishes executing

	return s.mp[key]
}

func main() {
	s := SafeCounter{mp: make(map[string]int)}

	//Create 1000 goroutines
	for i := 0; i < 1000; i++ {
		go s.Inc("somekey") //1000 goroutines,each of routines increases map value by 1
	}

	time.Sleep(time.Second)
	fmt.Println(s.mp["somekey"])
}
