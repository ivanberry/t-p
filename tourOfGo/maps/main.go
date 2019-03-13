package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int  {
	wordCountMap := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		if _, ok := wordCountMap[word]; ok {
			wordCountMap[word]++
		} else {
			wordCountMap[word] = 1
		}
	}
	return wordCountMap
}

func fibonacci() func() int {
	a := 1
	b := 0
	return func() int {
		if b < 0 {
			a, b = 1, 0
		}
		a, b = b, a + b
		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 100; i++ {
		fmt.Println(f())
	}
}