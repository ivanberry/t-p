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

func main() {
	c := WordCount("word, xafd, list, tag, red, green, green, tag, list")

	for k, d := range c {
		fmt.Printf("单词%v出现字数为 %d\n", k, d)
	}
}