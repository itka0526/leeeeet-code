package simple

import "fmt"

func findWordsContaining(words []string, x byte) (ans []int) {
	dupl := make(map[int]bool)
	for i := range words {
		for j := range words[i] {
			ch := words[i][j]
			if _, ok := dupl[i]; !ok && ch == x {
				dupl[i] = true
				ans = append(ans, i)
			}
		}
	}
	return ans
}

func TestFindWordsContaining() {
	s := []string{"leet", "code"}
	fmt.Println(findWordsContaining(s, 'e'))
}
