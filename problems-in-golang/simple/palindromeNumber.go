package simple

import "fmt"

func isPalindrome(x int) bool {
	s := []rune(fmt.Sprint(x))

	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func TestIsPalindrome() {
	fmt.Println(isPalindrome(121))
}
