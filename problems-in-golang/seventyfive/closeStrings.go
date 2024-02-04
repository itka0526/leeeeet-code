package seventyfive

import "sort"

func closeStrings(word1 string, word2 string) bool {
	m1 := map[rune]int{}
	m2 := map[rune]int{}
	for _, c := range word1 {
		m1[c]++
	}
	for _, c := range word2 {
		m2[c]++
	}
	if len(m1) != len(m2) {
		return false
	}
	for k1 := range m1 {
		if _, ok := m2[k1]; !ok {
			return false
		}
	}
	s1, s2 := make([]int, len(m1)), make([]int, len(m2))
	i := 0
	for _, v := range m1 {
		s1[i] = v
		i++
	}
	i = 0
	for _, v := range m2 {
		s2[i] = v
		i++
	}
	sort.Slice(s1, func(i2, j int) bool {
		return s1[i2] < s1[j]
	})
	sort.Slice(s2, func(i2, j int) bool {
		return s2[i2] < s2[j]
	})
	for x := 0; x < len(s1); x++ {
		if s1[x] != s2[x] {
			return false
		}
	}
	return len(s1) == len(s2)
}
