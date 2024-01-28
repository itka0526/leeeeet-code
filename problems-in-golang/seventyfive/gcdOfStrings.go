package seventyfive

func gcdOfStrings(s1, s2 string) string {
	if s1+s2 != s2+s1 {
		return ""
	}
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	return s1[0:gcd(len(s1), len(s2))]
}

// func gcdOfStrings(s1, s2 string) string {
// 	// Get the shorter string of the two
// 	// Then find the longest repeating string
// 	// for the longer string and for the shorter string
// 	if s1+s2 != s2+s1 {
// 		return ""
// 	}
// 	m := func(s1, s2 string) (string, string, string) {
// 		if len(s1) > len(s2) {
// 			return s1, s2, s2
// 		}
// 		return s2, s1, s1
// 	}

// 	c := func(s1, s2 string) bool {
// 		if len(s1) != len(s2) {
// 			return false
// 		}
// 		for i := range s1 {
// 			if s1[i] != s2[i] {
// 				return false
// 			}
// 		}
// 		return true
// 	}

// 	// We will modify sh string to be the end result
// 	lo, sh, mo := m(s1, s2)kkkkkk

// 	for i := 0; i < len(lo) && len(mo) != 0; i += len(mo) {
// 		if i+len(mo) > len(lo) {
// 			mo = mo[:len(mo)-1]
// 			i = 0
// 			continue
// 		}
// 		cu := lo[i : i+len(mo)]
// 		if c(cu, mo) {
// 			continue
// 		} else {
// 			mo = mo[:len(mo)-1
// 			i = 0
// 		}
// 	}

// 	for i := 0; i < len(sh) && len(mo) != 0; i += len(mo) {
// 		if i+len(mo) > len(sh) {
// 			mo = mo[:len(mo)-1]
// 			i = 0
// 			continue
// 		}
// 		cu := sh[i : i+len(mo)]
// 		if c(cu, mo) {
// 			continue
// 		} else {
// 			mo = mo[:len(mo)-1]
// 			i = 0
// 		}
// 	}

// 	return mo
// }
