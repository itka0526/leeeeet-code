package simple

import "fmt"

func strStr(haystack string, needle string) int {
	// Sliding window problem?
	l, r, n, s, cs := -1, 0, 0, []rune(haystack), []rune(needle)
	for r < len(s) {
		if s[r] == cs[n] {
			if len(cs)-1 == n {
				break
			}
			if l == -1 {
				l = r
			}
			n++
		} else {
			n = 0
			l = -1
		}
		r++
	}
	if l == -1 {
		return -1
	}
	return l
}

func TestStrStr() {
	fmt.Println("Exp: 2; Got: ", strStr("hello", "ll"))
	fmt.Println("Exp: -1; Got: ", strStr("leetcode", "leeto"))
	fmt.Println("Exp: 6; Got: ", strStr("byecodsadbutsad", "sad"))
}
