package seventyfive

import (
	"fmt"
	"unicode"
)

func reverseVowels(s string) string {
	ru := []rune(s)
	l := 0
	r := len(ru) - 1
	v := func(t rune) bool {
		ll := unicode.ToLower(t)
		return ll == 'a' || ll == 'e' || ll == 'i' || ll == 'o' || ll == 'u'
	}
	for l < r {
		for l < r && !v(ru[l]) {
			l++
		}
		for l < r && !v(ru[r]) {
			r--
		}
		if v(ru[l]) && v(ru[r]) {
			ru[l], ru[r] = ru[r], ru[l]
			l++
			r--
		}
	}
	return string(ru)
}

func TestReverseVowels() {
	fmt.Println("Expected: 'holle'; Got: ", reverseVowels("hello"))
}
