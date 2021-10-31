package hw4

import (
	"fmt"
	"sort"
	"strings"
)

func topWords(s string, n int) []string {
	stringMap := countWords(s)

	var values []int
	for _, v := range stringMap {
		values = append(values, v)
	}
	sort.Ints(values)

	var result []string
	cnt := 0
	for i := len(values) - 1; i >= 0; i-- {
		if cnt == n {
			break
		}
		if i == 0 || values[i] != values[i-1] {
			var buffer []string
			for key, val := range stringMap {
				if val == values[i] {
					buffer = append(buffer, key)
				}
			}
			sort.Strings(buffer)
			result = append(result, buffer...)
			cnt++
		}
	}
	fmt.Println(stringMap)
	return result
}


func countWords(s string) map[string]int {
	var str strings.Builder
	stringMap := make(map[string]int)
	for _, v := range s {
		if v != ' '  {
			if v == ',' || v == '.' || v == ':' || v == '"' { continue }
			str.WriteRune(v)
		} else {
			stringMap[str.String()]++
			str.Reset()
		}
	}
	stringMap[str.String()]++
	return stringMap
}
