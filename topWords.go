package hw4

import (
	"sort"
	"strings"
)
type wordWithCount struct {
	word string
	count int
}

func topWords(s string, n int) []string {
	stringMap := countWords(s)

	var words []wordWithCount
	for key, value := range stringMap {
		words = append(words, wordWithCount{key, value})
	}
	sort.Slice(words, func(i, j int) bool {
		if words[i].count == words[j].count {
			return strings.Compare(words[i].word, words[j].word) == -1
		}
		return words[i].count > words[j].count
	})
	
	var result []string
	cnt := 0
	for i := 0; i < len(words) - 1; i++ {
		if words[i].count != words[i+1].count {
			cnt++
		}
		result = append(result, words[i].word)
		if cnt == n {
			break
		}
	}
	return result
}


func countWords(s string) map[string]int {
	replacer := strings.NewReplacer(",", "",
		".", "",
		":", "",
		"'", "")
	strArr := strings.Fields(replacer.Replace(s))
	stringMap := make(map[string]int)
	for _, v := range strArr {
		stringMap[v]++
	}
	return stringMap
}
