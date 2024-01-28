package seventyfive

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	rsa := func(s string) []string {
		res := []string{}
		w := []rune(s)
		l, r := 0, 0
		for l < len(w) {
			if w[l] != ' ' {
				r = l
				for r < len(w) {
					if w[r] == ' ' {
						break
					}
					r++
				}
				res = append(res, string(w[l:r]))
				l = r
			}
			l++
		}
		return res
	}
	ss := rsa(s)
	fmt.Println(ss)

	for i, j := 0, len(ss)-1; i < j; i, j = i+1, j-1 {
		ss[i], ss[j] = ss[j], ss[i]
	}
	return strings.Join(ss, " ")
}

func TestReverseWords() {
	fmt.Println("Expected: 'example good a'; Got: ", reverseWords("a good   example"))
}
