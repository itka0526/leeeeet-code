package slidingwindow

import (
	"github.com/itka0526/problems-in-golang/functions"
)

// Given a string s, find the length of the longest
// substring without repeating characters.

func lengthOfLongestSubstring(s string) int {
	// sliding window perhaps?
	max := 0

	for i := 0; i < len(s); i++ {
		j := i + 1
		m := make(map[byte]bool)
		tmp := 1

		m[s[i]] = true

		for j < len(s) {
			if _, ok := m[s[j]]; ok {
				break
			}

			m[s[j]] = true

			tmp++
			j++
		}

		if tmp > max {
			max = tmp
		}
	}

	return max
}

func TestLengthOfLongestSubstring() {
	functions.PrintSolution(lengthOfLongestSubstring("abcabcbb"))
}
