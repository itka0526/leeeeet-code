package simple

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	strs := strings.Split(s, " ")
	return len(strs[len(strs)-1])
}

func TestLengthOfLastWord() {
	fmt.Println("Expt: 4; Got: ", lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println("Expt: 6; Got: ", lengthOfLastWord("luffy is still joyboy"))
}
