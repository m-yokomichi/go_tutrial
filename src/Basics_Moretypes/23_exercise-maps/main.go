package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func main() {
	wc.Test(WordCount)
}

func WordCount(s string) map[string]int {
	// カウントするために単語単位で区切る
	words := strings.Fields(s)
	wordCount := make(map[string]int)
	
	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}