package backtracking

import (
	"github.com/itka0526/problems-in-golang/functions"
)

func existNaive(board [][]byte, word string) bool {
	vis := make([][]int, len(board))
	for i := range vis {
		vis[i] = make([]int, len(board[i]))
	}
	var dfs func(curr []byte, i, j int) bool
	dfs = func(curr []byte, i, j int) bool {
		if i < 0 || j < 0 || i > len(board)-1 || j > len(board[0])-1 || vis[i][j] == 1 {
			return false
		}
		vis[i][j] = 1
		curr = append(curr, board[i][j])
		if string(curr) == word {
			return true
		}
		for _, k := range dirs {
			di, dj := k[0], k[1]
			if dfs(curr, i+di, j+dj) {
				return true
			}
		}
		vis[i][j] = 0
		return false
	}
	for i := range board {
		for j := range board[0] {
			if dfs([]byte{}, i, j) {
				return true
			}
		}
	}
	return false
}

var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func exist(board [][]byte, word string) (ans bool) {
	n, m := len(board), len(board[0])
	ans, vis := false, make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, m)
	}
	var dfs func(c, i, j int) bool
	dfs = func(c, i, j int) bool {
		if c == len(word) {
			return true
		}
		if i < 0 || j < 0 || i > n-1 || j > m-1 || word[c] != board[i][j] || vis[i][j] {
			return false
		}
		vis[i][j] = true
		for _, d := range dirs {
			di, dj := d[0], d[1]
			if dfs(c+1, i+di, j+dj) {
				return true
			}
		}
		vis[i][j] = false
		return false
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if dfs(0, i, j) {
				ans = true
				return
			}
		}
	}
	return ans
}

func TestExist() {
	s := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	functions.LogAnswer(exist(s, "ABCCED"), true)

	s1 := [][]byte{{'b', 'b'}, {'a', 'b'}, {'b', 'a'}}
	functions.LogAnswer(exist(s1, "a"), true)

	s2 := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	functions.LogAnswer(exist(s2, "ABCB"), false)

	s3 := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	functions.LogAnswer(exist(s3, "SEE"), true)
}
