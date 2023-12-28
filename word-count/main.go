package main

import (
	"fmt"
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	subStrings := strings.Fields(s)

	wordMap := make(map[string]int)

	for i := 0; i < len(subStrings); i++ {
		key := subStrings[i]
		count, ok := wordMap[key]
		
		if !ok {
			wordMap[key] = 1
		} else {
			wordMap[key] = count + 1
		}
	}
	
	fmt.Println(wordMap)
	
	return wordMap
}

func main() {
	wc.Test(WordCount)

	// WordCount("I love love Go!")
}
