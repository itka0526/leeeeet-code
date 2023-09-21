package stack

import (
	"github.com/itka0526/problems-in-golang/functions"
)

// func isValid(s string) bool {
// 	pairs := map[byte]byte{
// 		'}': '{',
// 		']': '[',
// 		')': '(',
// 	}

// 	stack := make([]byte, 0)

// 	// I see now its trying to append only the start of the parenthesis
// 	// And then once we ran out of starting parenthesis we wouldn't be
// 	// Able to continue/skip and we would fall to next lines
// 	for _, char := range []byte(s) {
// 		pair, ok := pairs[char]
// 		if !ok {
// 			stack = append(stack, char)
// 			continue
// 		}

// 		// this is for just to save some performance
// 		if len(stack) == 0 {
// 			return false
// 		}

// 		// if last of the starting pair does not match the ending pair we return false
// 		if stack[len(stack)-1] != pair {
// 			return false
// 		}

// 		// here we just cut that part away if they are matching
// 		stack = stack[:len(stack)-1]
// 	}

// 	// result it determined by the length OFC! Genius indeed
// 	return len(stack) == 0
// }

func isValid1(s string) bool {
	stack := []byte{}
	end_to_start := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}

	for _, b := range []byte(s) {
		if b == '[' || b == '(' || b == '{' {
			stack = append(stack, b)
			continue
		}

		// it might start with ending parenthesis we have to take that into consideration
		if len(stack) == 0 || stack[len(stack)-1] != end_to_start[b] {
			return false
		}

		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}

func TestIsValid() {
	functions.PrintSolution(
		isValid1("()[]{}"),
	)
}
