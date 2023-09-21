package slidingwindow

import (
	"github.com/itka0526/problems-in-golang/functions"
)

// order matters !!
// func checkInclusion(s1 string, s2 string) bool {
// 	m := make(map[rune]int)

// 	for _, c := range s2 {
// 		m[c]++
// 	}

// 	count := len(s1)
// 	for _, c := range s1 {
// 		if _, ok := m[c]; ok {
// 			m[c]--
// 			count--
// 		}
// 	}

// 	return count == 0
// }

// on larger strings it does not work... 80/108... Naive approach.
// I think there is a trick to solving this... but I am not sure what method to use
// currently using HashMaps + Sliding Window.
// func checkInclusion(s1 string, s2 string) bool {
// 	m1 := make(map[byte]int)

// 	for i := 0; i < len(s1); i++ {
// 		m1[s1[i]]++
// 	}

// 	for i := 0; i < len(s2); i++ {
// 		if _, ok := m1[s2[i]]; ok {
// 			cp := copyMap(m1)
// 			cp[s2[i]]--

// 			if cp[s2[i]] == 0 {
// 				delete(cp, s2[i])
// 			}

// 			for j := i + 1; j < len(s2); j++ {
// 				if _, ok := cp[s2[j]]; ok {
// 					cp[s2[j]]--

// 					if cp[s2[j]] == 0 {
// 						delete(cp, s2[j])
// 					}
// 				} else {
// 					break
// 				}
// 			}

// 			if len(cp) == 0 {
// 				return true
// 			}
// 		}
// 	}

// 	return false
// }

// func copyMap[K comparable, V any](m map[K]V) map[K]V {
// 	newMap := make(map[K]V)

// 	for k, v := range m {
// 		newMap[k] = v
// 	}

// 	return newMap
// }

// if you could apply some logic I could have solved I think
// hints were size of the window never changes therefore we could have iterated over each window moving by 1 
// and not think it as dynamic size thats was my mistake
func checkInclusion(s1 string, s2 string) bool {
	a := make(map[rune]int)
	for _, c := range s1 {
		a[c]++
	}

	size := len(s1)
	for i := 0; i < len(s2); i++ {
		j := size + i
		b := make(map[rune]int)

		if j > len(s2) {
			break
		}
		window := s2[i:j]
		for _, c := range window {
			b[c]++
		}
		if compareMaps(a, b) {
			return true
		}
	}

	return false
}

func compareMaps[K comparable, V comparable](m1, m2 map[K]V) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k := range m1 {
		if m1[k] != m2[k] {
			return false
		}
	}

	return true
}

func TestCheckInclusion() {
	functions.PrintSolution(checkInclusion("adc", "dcda"))
}
