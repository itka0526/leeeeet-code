package medium

import (
	"fmt"
	"strconv"
)

func getIntitialCombo(target string) []rune {
	s := make([]rune, len(target))
	for i := range s {
		s[i] = '0'
	}
	return s
}

func getDeadlockMap(s []string) (m map[string]bool) {
	m = make(map[string]bool)
	for i := range s {
		m[s[i]] = true
	}
	return m
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// func openLock(deadends []string, target string) (ans int) {
// 	s := getIntitialCombo(target)
// 	m := getDeadlockMap(deadends)
// 	i := 0
// 	l := len(target)
// 	cnt := 0
// 	for string(s) != target {
// 		d := runeToInt(s[i])
// 		t := runeToInt(rune(target[i]))

// 		if d == t {
// 			i += 1
// 			i %= l
// 			continue
// 		} else if abs(t-d) > abs(10-t+d) {
// 			d -= 1
// 		} else {
// 			d += 1
// 		}

// 		if d < 0 {
// 			d += 10
// 		}

// 		d %= 10
// 		cnt += 1

// 		s[i] = intToRune(d)

// 		if m[string(s)] {
// 			cnt -= 1
// 		}

// 		for m[string(s)] {
// 			cnt += 1
// 			prevIdx := i - 1%l
// 			p := (runeToInt(s[prevIdx]) + 1) % 10
// 			s[prevIdx] = intToRune(p)
// 		}

// 		i += 1
// 		i %= l
// 		fmt.Println(string(s))
// 	}
// 	fmt.Println(string(s), "=", target, m, "; Count: ", cnt)
// 	return ans
// }

// func openLock(deadends []string, target string) (ans int) {
// 	// BFS - finds the shortest path in unweighted graphs
// 	vis := make(map[string]bool, 10_000)

// 	for i := range deadends {
// 		vis[deadends[i]] = true
// 	}

// 	q := make([]string, 1)
// 	q[0] = "0000"

// 	for len(q) > 0 {
// 		c := q[0]
// 		q = q[1:]

// 		if vis[c] {
// 			continue
// 		}
// 		vis[c] = true

// 		if c == target {
// 			// Found
// 			return
// 		}
// 		// Generate paths
// 		for i := 0; i < 4; i++ {
// 			if p, n := generateP(c, i); p {
// 				q = append(q, n)
// 			}
// 			if p, n := generateM(c, i); p {
// 				q = append(q, n)
// 			}
// 		}
// 	}
// 	return ans
// }

func runeToInt(r rune) int {
	d, err := strconv.Atoi(string(r))
	if err != nil {
		panic("Cannot translate rune into int.")
	}
	return d
}

func intToRune(i int) rune {
	if i < 0 || i > 9 {
		panic(fmt.Sprintf("One rune one digit; Value: %d", i))
	}
	return []rune(fmt.Sprint(i))[0]
}

func generate(c string, idx int) (string, string) {
	pt := []rune(c)
	mt := []rune(c)
	p := runeToInt(pt[idx])
	m := p
	m -= 1
	if m < 0 {
		m += 10
	}
	p += 1
	if p > 9 {
		p %= 10
	}
	pt[idx] = intToRune(p)
	mt[idx] = intToRune(m)
	return string(pt), string(mt)
}

func generateM(c string, idx int) (bool, string) {
	t := []rune(c)
	d := runeToInt(t[idx]) - 1
	if d < 0 {
		return false, ""
	}
	t[idx] = intToRune(d)
	return true, string(t)
}

type Combo struct {
	value string
	moves int
}

func openLock(deadends []string, target string) int {
	vis := make(map[string]bool)
	for i := range deadends {
		if deadends[i] == "0000" {
			return -1
		}
		vis[deadends[i]] = true
	}
	q := make([]Combo, 1)
	q[0] = Combo{value: "0000", moves: 0}
	vis["0000"] = true
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if curr.value == target {
			return curr.moves
		}
		for i := 0; i < 4; i++ {
			p, m := generate(curr.value, i)
			if !vis[p] {
				vis[p] = true
				q = append(q, Combo{value: p, moves: curr.moves + 1})
			}
			if !vis[m] {
				vis[m] = true
				q = append(q, Combo{value: m, moves: curr.moves + 1})
			}
		}
	}
	return -1
}

func TestOpenLock() {
	s := []string{"0201", "0101", "0102", "1212", "2002"}
	fmt.Println(openLock(s, "0202"))

	s2 := []string{}
	fmt.Println(openLock(s2, "0004"))

	s1 := []string{"8888"}
	fmt.Println(openLock(s1, "0009"))
}
