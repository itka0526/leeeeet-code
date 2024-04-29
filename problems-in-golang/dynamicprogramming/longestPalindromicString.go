package dynamicprogramming

func longestPalindrome(s string) (ans string) {
	for i := 0; i < len(s); i++ {
		l, r := i, i
		for r+1 < len(s) && s[l] == s[r+1] {
			r++
		}
		for l-1 >= 0 && r+1 < len(s) && s[l-1] == s[r+1] {
			l--
			r++
		}
		tmp := s[l : r+1]
		if len(tmp) >= len(ans) {
			ans = tmp
		}
	}
	return ans
}
