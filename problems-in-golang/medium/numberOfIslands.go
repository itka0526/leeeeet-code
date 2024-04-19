package medium

import "fmt"

func dfsIsland(grid *[][]byte, i, j, n, m int) {
	if i < 0 || j < 0 || i > n-1 || j > m-1 || (*grid)[i][j] == '0' {
		return
	}
	if (*grid)[i][j] == '1' {
		(*grid)[i][j] = '0'
	}
	dirs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, dir := range dirs {
		dfsIsland(grid, i+dir[0], j+dir[1], n, m)
	}
}

func numIslands(grid [][]byte) (ans int) {
	n := len(grid)
	m := len(grid[0])
	vis := make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				dfsIsland(&grid, i, j, n, m)
				ans += 1
			}
		}
	}
	return ans
}

func TestNumsIslands() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(grid))
}
