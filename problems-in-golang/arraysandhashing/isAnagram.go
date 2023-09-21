package arraysandhashing

import "fmt"

func isAnagramSolution(s string, t string) bool {
	m := make(map[rune]int)

	for i := 0; i < len(s); i++ {
		// key - character, value - count
		m[rune(s[i])] += 1
	}

	for i := 0; i < len(t); i++ {
		// key - character, value - count
		m[rune(t[i])] -= 1
	}

	// not using 'v' because it creates a copy which means more memory
	for k := range m {
		if m[k] != 0 {
			return false
		}
	}

	return true
}

func TestIsAnagram() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagramSolution(s, t))
}
