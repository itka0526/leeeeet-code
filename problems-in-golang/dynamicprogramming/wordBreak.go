package dynamicprogramming

func wordBreak(s string, wordDict []string) bool {
	cache := make([]bool, len(s)+1)
	cache[len(s)] = true
	for i := len(s) - 1; i > -1; i -= 1 {
		for _, w := range wordDict {
			if (len(w)+i) <= len(s) && s[i:i+len(w)] == w {
				cache[i] = cache[i+len(w)]
			}
			if cache[i] {
				break
			}
		}
	}
	return cache[0]
}
