package simple

import "fmt"

func countPerimeter(m [][]int, i, j int) int {
	count := 4
	if m[i][j] == 1 {
		if i-1 >= 0 && m[i-1][j] == 1 {
			count--
		}
		if i+1 < len(m) && m[i+1][j] == 1 {
			count--
		}
		if j-1 >= 0 && m[i][j-1] == 1 {
			count--
		}
		if j+1 < len(m[i]) && m[i][j+1] == 1 {
			count--
		}
	}
	return count
}

func dfsIsland(
	m [][]int,
	i int,
	j int,
	p *int,
	vis *[][]int,
) {
	// out of bounds
	if i < 0 || i > len(m)-1 || j < 0 || j > len(m[i])-1 {
		return
	}
	// ignore non-island
	if m[i][j] == 0 {
		return
	}
	// visited block
	if (*vis)[i][j] == 7 {
		return
	}
	// count sides
	*p += countPerimeter(m, i, j)
	// mark as visited
	(*vis)[i][j] = 7
	dfsIsland(m, i+1, j, p, vis)
	dfsIsland(m, i-1, j, p, vis)
	dfsIsland(m, i, j+1, p, vis)
	dfsIsland(m, i, j-1, p, vis)
}

func islandPerimeter(grid [][]int) int {
	p := 0
	vis := make([][]int, len(grid))
	for i := range vis {
		vis[i] = make([]int, len(grid[i]))
	}
out:
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				// entry point
				dfsIsland(grid, i, j, &p, &vis)
				break out
			}
		}
	}
	return p
}

func TestIslandPerimeter() {
	s := [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}
	s1 := [][]int{{1}}

	fmt.Println(islandPerimeter(s1))
	fmt.Println(islandPerimeter(s))
}
