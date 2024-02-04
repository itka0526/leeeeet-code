package simple

import (
	"fmt"
	"math"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	f := []string{}
	for i, s1 := range strs {
		for j, s2 := range strs {
			if i == j {
				continue
			}
			k := 0
			for k < len(s1) && k < len(s2) && s1[k] == s2[k] {
				k++
			}
			f = append(f, s1[:k])
		}
	}
	l := math.MaxInt32
	ss := ""
	for _, s := range f {
		if len(s) < l {
			ss = s
			l = len(s)
		}
	}
	return ss
}

func TestLongestCommonPrefix() {
	fmt.Println("Exp: ''; Got: ", longestCommonPrefix([]string{"reflower", "flow", "flight"}))
	// fmt.Println("Exp: 'a'; Got: ", longestCommonPrefix([]string{"a"}))
	// fmt.Println("Exp: 'fl'; Got: ", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	// fmt.Println("Exp: ''; Got: ", longestCommonPrefix([]string{"dog", "racecar", "car"}))
	// fmt.Println("Exp: 'a'; Got: ", longestCommonPrefix([]string{"ab", "a"}))
}
