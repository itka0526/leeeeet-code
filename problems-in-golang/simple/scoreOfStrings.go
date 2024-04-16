package simple

import "fmt"

func scoreOfString(s string) int {
	v := []rune(s)
	score := 0
	for i := 0; i < len(v)-1; i++ {
		curr := int(v[i] - v[i+1])
		if curr < 0 {
			score += -curr
		} else {
			score += curr
		}
	}
	return score
}

func TestScoreOfString() {
	s := "hello"
	fmt.Println(scoreOfString(s))
}
