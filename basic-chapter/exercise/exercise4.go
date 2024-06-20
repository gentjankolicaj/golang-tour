package main

import (
	"golang.org/x/tour/wc"
)

//My implementation without using string.Fields
//This methods extracts characters & words from string without spaces
func wordCount(s string) map[string]int {
	var myMap = make(map[string]int)
	var spaceChar int32 = 32
	var spaceIdx = -1
	var charIdx = -1

	for i, v := range s {
		if spaceChar == v {
			if charIdx != -1 {
				var tmpStr = s[spaceIdx+1 : charIdx+1]
				counter, flag := myMap[tmpStr]
				if flag {
					myMap[tmpStr] = counter + 1
				} else {
					myMap[tmpStr] = 1
				}
				spaceIdx = i
			}
		} else {
			charIdx = i
		}

		if i == len(s)-1 {
			var tmpStr = s[spaceIdx+1 : charIdx+1]
			counter, flag := myMap[tmpStr]
			if flag {
				myMap[tmpStr] = counter + 1
			} else {
				myMap[tmpStr] = 1
			}
		}
	}

	return myMap
}
func main() {
	wc.Test(wordCount)
}
