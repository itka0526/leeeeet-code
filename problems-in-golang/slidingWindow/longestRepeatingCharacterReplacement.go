package slidingwindow

import (
	"github.com/itka0526/problems-in-golang/functions"
)

func characterReplacement(s []rune, k int) int {
	frequentMap := make(map[rune]int)
	maxLength := 0

	left := 0
	maxf := 0
	for right := 0; right < len(s); right++ {
		frequentMap[s[right]] = 1 + frequentMap[s[right]]
		maxf = maxNumber(maxf, frequentMap[s[right]])
		if (right-left+1)-maxf > k {
			frequentMap[s[left]]--
			left++
		}
		maxLength = maxNumber(maxLength, right-left+1)
	}

	return maxLength
}

func maxNumber(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 34/39 cool :D This works! Even though it failed some.
// func characterReplacement(s string, k int) int {
// 	characters := []rune(s)
// 	maxLength := 0

// 	for i := 0; i < len(characters); i++ {
// 		local_k := k
// 		length := 1
// 		for j := i + 1; j < len(characters); j++ {
// 			if characters[i] != characters[j] {
// 				// we will check how many k's we got
// 				if local_k > 0 {
// 					// let it pass to the next if
// 					local_k--
// 					goto UseReplace
// 				} else {
// 					// not enough k's so we cannot replace any characters more
// 					// we should update
// 					break
// 				}
// 			}

// 		UseReplace:
// 			length++
// 		}

// 		if length > maxLength {
// 			maxLength = length
// 		}
// 	}

// 	return maxLength
// }

func TestCharacterReplacement() {
	functions.PrintSolution(characterReplacement([]rune("ABCCDC"), 1))
}
