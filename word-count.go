package wordCount

import "strings"

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordCountMap := make(map[string]int)
	for _, word := range words {
		wordCountMap[word]++
	}
	return wordCountMap
}
