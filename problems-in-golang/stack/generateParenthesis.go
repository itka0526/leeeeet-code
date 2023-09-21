package stack

import (
	"strings"

	"github.com/itka0526/problems-in-golang/functions"
)

func generateParenthesisSolution(n int) []string {
	var res []string
	var stack []string

	var b func(int, int)

	b = func(o int, c int) {
		if o == n && o == c && c == n {
			res = append(res, strings.Join(stack, ""))
			return
		}
		if o < n {
			stack = append(stack, "(")
			b(o+1, c)
			p(&stack)
		}

		if c < o {
			stack = append(stack, ")")
			b(o, c+1)
			p(&stack)
		}
	}

	b(0, 0)
	return res
}

func p(s *[]string) {
	*s = (*s)[:len(*s)-1]
}

func TestGenerateParenthesisSolution() {
	functions.PrintSolution(generateParenthesisSolution(3))
}
