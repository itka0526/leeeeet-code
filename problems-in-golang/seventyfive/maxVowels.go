package seventyfive

import "fmt"

func maxVowels(s string, k int) int {
	l, r := 0, 0
	ss := []rune(s)
	t, m := 0, 0
	v := func(c rune) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
	}
	for r < k {
		if v(ss[r]) {
			t++
		}
		r++
	}
	// We should initialize max to be the initial total
	// If we don't then, in case, if k is equal to the length - 1 of our string we cannot
	// check whether the maximum is the absolute maximum because the first letter could be a
	// vowel. Off by one ERRORS are the most fun...
	m = t
	for r < len(ss) {
		if v(ss[l]) {
			t--
		}
		if v(ss[r]) {
			t++
		}
		if t > m {
			m = t
		}
		l++
		r++
	}
	return m
}

func TestMaxVowels() {
	fmt.Println("Exp: 7; Got: ", maxVowels("ibpbhixfiouhdljnjfflpapptrxgcomvnb", 33))
	fmt.Println("Exp: 3; Got: ", maxVowels("abciiidef", 3))
	fmt.Println("Exp: 2; Got: ", maxVowels("leetcode", 3))
}
