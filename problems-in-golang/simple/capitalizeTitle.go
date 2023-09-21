package simple

import (
	"fmt"
	"strings"
	"unicode"
)

func capitalizeTitle(title string) string {
	words := strings.Split(title, " ")
	res := make([]string, len(words))

	for i, word := range words {
		allLower := strings.ToLower(word)
		if len(allLower) > 2 {
			capitalized := []rune(allLower)
			capitalized[0] = unicode.ToUpper(capitalized[0])
			res[i] = string(capitalized)
			continue
		}
		res[i] = allLower
	}

	return strings.Join(res, " ")
}

func complexCapitalizeTitle(title string) string {
	// ASCII value of 'a' (lowercase 'a') is 97, and the ASCII value of 'A' (uppercase 'A') is 65.
	// The difference between them is 32.

	res := []rune(title)

	var lastWhiteSpace, start int
	for i := range res {
		res[i] |= (1 << 5)

		if res[i] == ' ' {
			if lastWhiteSpace > 2 {
				res[start] -= (1 << 5)
			}
			start = i + 1
			lastWhiteSpace = 0
		} else {
			lastWhiteSpace++
		}
	}

	if lastWhiteSpace > 2 {
		res[start] -= (1 << 5)
	}

	return string(res)
}

func TestComplexCapitalizeTitle() {
	fmt.Println("Solution: ", complexCapitalizeTitle("HeLLo"))
}
