package simple

import "fmt"

func closetTarget(words []string, target string, startIndex int) int {
	// Formally, the next element of words[i] is words[(i + 1) % n]
	// and the previous element of words[i] is words[(i - 1 + n) % n],

	// where n is the length of words.
	left, right := startIndex, startIndex
	c, cc := 0, 0
	for {
		if len(words)-1 < right {
			if right-len(words) == len(words) {
				return -1
			}
			if words[right-len(words)] == target {
				fmt.Println("right")
				return cc
			}
		} else if words[right] == target {
			fmt.Println("right")
			return c
		}
		right++
		c++

		if left < 0 {
			if len(words)+left == 0 {
				return -1
			}
			if words[len(words)+left] == target {
				fmt.Println("left")
				return cc
			}
		} else if words[left] == target {
			fmt.Println("left")
			return cc
		}

		cc++
		left--
	}
}

func TestClosetTarget() {
	s := []string{"hsdqinnoha", "mqhskgeqzr", "zemkwvqrww", "zemkwvqrww", "daljcrktje", "fghofclnwp", "djwdworyka", "cxfpybanhd", "fghofclnwp", "fghofclnwp"}
	t := "zemkwvqrww"
	start := 8

	fmt.Println("Answer: ", closetTarget(s, t, start))
}
