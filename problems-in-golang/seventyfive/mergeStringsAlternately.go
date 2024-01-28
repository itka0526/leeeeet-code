package seventyfive

func MergeAlternately(word1 string, word2 string) string {
	i := 0
	n := ""
	mm := func(a, b string) (int, string) {
		if len(a) < len(b) {
			return len(a), b
		}
		return len(b), a
	}
	m, x := mm(word1, word2)
	for m > 0 {
		c := string(word1[i]) + string(word2[i])
		n += c
		i++
		m--
	}
	n += x[i:]
	return n
}
