package medium

import "fmt"

func dfsFarm(land *[][]int, ans *[][]int, i, j, n, m int) {
	if i < 0 || j < 0 || i > n-1 || j > m-1 {
		return
	}
	if (*land)[i][j] == 0 {
		return
	}
	if i > (*ans)[len(*ans)-1][2] {
		(*ans)[len(*ans)-1][2] = i
	}
	if j > (*ans)[len(*ans)-1][3] {
		(*ans)[len(*ans)-1][3] = j
	}
	(*land)[i][j] = 0
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, dir := range dirs {
		dfsFarm(land, ans, i+dir[0], j+dir[1], n, m)
	}
}

func findFarmland(land [][]int) (ans [][]int) {
	n, m := len(land), len(land[0])
	cp := make([][]int, len(land))
	copy(cp, land)
	for i := range cp {
		for j := range cp[i] {
			if cp[i][j] == 1 {
				ans = append(ans, []int{i, j, 0, 0})
				dfsFarm(&cp, &ans, i, j, n, m)
			}
		}
	}
	return ans
}

func TestFindFarmLand() {
	land := [][]int{{1, 0, 0}, {0, 1, 1}, {0, 1, 1}}
	fmt.Println(findFarmland(land))
}
