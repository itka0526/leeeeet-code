package twopointers

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	re := regexp.MustCompile(`[a-z]|\d`)

	characters := re.FindAllString(strings.ToLower(s), -1)

	fmt.Println(characters)

	l := 0
	r := len(characters) - 1

	for l <= r {
		if characters[l] != characters[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func TestIsPalindrome() {
	s := "A man, a plan, a canal: Panama"

	fmt.Printf("Is palyndrome: %t \n", isPalindrome(s))

	s2 := "0P"

	fmt.Printf("Is palyndrome: %t \n", isPalindrome(s2))
}
