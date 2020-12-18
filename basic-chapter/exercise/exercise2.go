package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	//space according to utf-8 is 20
	//space according to ascii is 32
	var myMap = make(map[string]int)
	fields := strings.Fields(s)

	for _, val := range fields {
		value, flag := myMap[val]
		if flag == true {
			myMap[val] = value + 1
		} else {
			myMap[val] = 1
		}

	}
	return myMap
}

//Mine impl without using strings.Fields
func WordCount2(s string) map[string]int {
	var myMap = make(map[string]int)
	var spaceChar int32 = 32
	var spaceIdx int = -1
	var charIdx int = 0

	for i, v := range s {
		if spaceChar == v {
			if charIdx != 0 {
				var tmpStr string = s[spaceIdx+1 : charIdx]
				fmt.Println(tmpStr)
				spaceIdx = i
			}
		} else {
			charIdx = i
		}
	}

	return myMap
}
func main() {
	//wc.Test(WordCount)
	//	fmt.Println("Word count my custom function :")
	wc.Test(WordCount2)
}
