package arraysandhashing

import (
	"fmt"
)

// Determine if a 9 x 9 Sudoku board is valid.
// Only the filled cells need to be validated according to the following rules:

// Each row must contain the digits 1-9 without repetition.
// Each column must contain the digits 1-9 without repetition.
// Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
// Note:

// A Sudoku board (partially filled) could be valid but is not necessarily solvable.
// Only the filled cells need to be validated according to the mentioned rules.

func isValidSudokuSolution(board [][]string) bool {
	m := make(map[string]bool)

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			current_val := string(board[r][c])

			if current_val == "." {
				continue
			}

			_, ok1 := m[current_val+"r"+fmt.Sprint(r)]
			_, ok2 := m[current_val+"c"+fmt.Sprint(c)]
			_, ok3 := m[current_val+"g"+fmt.Sprint(r/3)+"-"+fmt.Sprint(c/3)]

			if ok1 || ok2 || ok3 {
				return false
			}

			m[current_val+"found in row"+fmt.Sprint(r)] = true
			m[current_val+"found in col"+fmt.Sprint(c)] = true
			m[current_val+"found in grid"+fmt.Sprint(r/3)+"-"+fmt.Sprint(c/3)] = true
		}
	}

	return true
}

func stfuCompiler(a ...interface{}) {}

func TestIsValidSudoku() {
	sudoku := [][]string{
		{"5", "3", ".", ".", "7", ".", ".", ".", "."}, {"6", ".", ".", "1", "9", "5", ".", ".", "."}, {".", "9", "8", ".", ".", ".", ".", "6", "."}, {"8", ".", ".", ".", "6", ".", ".", ".", "3"}, {"4", ".", ".", "8", ".", "3", ".", ".", "1"}, {"7", ".", ".", ".", "2", ".", ".", ".", "6"}, {".", "6", ".", ".", ".", ".", "2", "8", "."}, {".", ".", ".", "4", "1", "9", ".", ".", "5"}, {".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}
	fmt.Println(isValidSudokuSolution(sudoku))

	sudoku2 := [][]string{
		{"8", "3", ".", ".", "7", ".", ".", ".", "."}, {"6", ".", ".", "1", "9", "5", ".", ".", "."}, {".", "9", "8", ".", ".", ".", ".", "6", "."}, {"8", ".", ".", ".", "6", ".", ".", ".", "3"}, {"4", ".", ".", "8", ".", "3", ".", ".", "1"}, {"7", ".", ".", ".", "2", ".", ".", ".", "6"}, {".", "6", ".", ".", ".", ".", "2", "8", "."}, {".", ".", ".", "4", "1", "9", ".", ".", "5"}, {".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}

	fmt.Println(isValidSudokuSolution(sudoku2))
}
