package seventyfive

import "fmt"

func removeStars(s string) string {
	r := 0
	res := make([]rune, 0)
	for r < len(s) {
		c := s[r]
		if c == byte('*') {
			res = res[:len(res)-1]
		} else {
			res = append(res, rune(s[r]))
		}
		r++
	}
	return string(res)
}

func TestRemoveStars() {
	fmt.Println("Exp: lecoe; Got: ", removeStars("leet**cod*e"))
}
