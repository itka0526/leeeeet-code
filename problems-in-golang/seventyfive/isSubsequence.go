package seventyfive

func isSubsequence(s string, t string) bool {
	l, r := 0, 0
o:
	for l < len(s) {
		for r < len(t) {
			if s[l] == t[r] {
				l++
				r++
				continue o
			}
			r++
		}
		return false
	}
	return true
}
