package stack

import (
	"strconv"

	"github.com/itka0526/problems-in-golang/functions"
)

func evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		switch token {
		case "+":
			last, lastBefore := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			stack = append(stack, (last + lastBefore))
		case "-":
			last, lastBefore := stack[len(stack)-1], stack[len(stack)-2]

			stack = stack[:len(stack)-2]

			stack = append(stack, (lastBefore - last))
		case "*":
			last, lastBefore := stack[len(stack)-1], stack[len(stack)-2]

			stack = stack[:len(stack)-2]

			stack = append(stack, (last * lastBefore))
		case "/":
			last, lastBefore := stack[len(stack)-1], stack[len(stack)-2]

			stack = stack[:len(stack)-2]

			stack = append(stack, (lastBefore / last))
		default:
			n, _ := strconv.Atoi(token)
			stack = append(stack, n)
		}
	}

	return stack[0]
}

func TestEvalRPN() {
	functions.PrintSolution(evalRPN([]string{"2", "1", "+", "3", "*"}), evalRPN([]string{"4", "13", "5", "/", "+"}))
}
