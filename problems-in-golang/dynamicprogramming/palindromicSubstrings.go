package dynamicprogramming

func countSubstrings(s string) (ans int) {
	for i := range s {
		l, r := i, i
		ans += 1
		for r+1 < len(s) && s[l] == s[r+1] {
			ans += 1
			r++
		}
		for l-1 >= 0 && r+1 < len(s) && s[l-1] == s[r+1] {
			l--
			r++
			ans += 1
		}
	}
	return ans
}
