package arraysandhashing

import (
	"fmt"
	"sort"
)

func groupAnagramsSolution(strs []string) [][]string {
	// key is combination of sorted rune, value is the strings that sum up to this value
	// oh fuck I cannot use combinations because some words have same value damn thought I was being clever
	runesStrs := make(map[string][]string)
	results := [][]string{}

	for i := range strs {
		runes := []rune(strs[i])

		sort.Slice(runes, func(i, j int) bool {
			return runes[i] > runes[j]
		})

		s := string(runes)
		runesStrs[s] = append(runesStrs[s], strs[i])
	}

	for k := range runesStrs {
		results = append(results, runesStrs[k])
	}

	return results
}

func TestGroupAnagrams() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagramsSolution(strs))
}
