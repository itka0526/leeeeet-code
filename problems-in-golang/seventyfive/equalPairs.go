package seventyfive

import "fmt"

func equalPairs(grid [][]int) int {
	r := map[string]int{}
	c := map[string]int{}
	d := len(grid)
	for i := 0; i < d; i++ {
		k1 := ""
		k2 := ""
		for j := 0; j < d; j++ {
			k1 += fmt.Sprint(grid[i][j], " ")
			k2 += fmt.Sprint(grid[j][i], " ")
		}
		r[k1]++
		c[k2]++
	}
	k := 0
	for k1, v1 := range r {
		if v2, ok := c[k1]; ok {
			k += v1 * v2
		}
	}
	return k
}

func TestEqualPairs() {
	fmt.Println("Exp: 48; Got: ", equalPairs([][]int{{3, 3, 3, 6, 18, 3, 3, 3, 3, 3}, {3, 3, 3, 3, 1, 3, 3, 3, 3, 3}, {3, 3, 3, 3, 1, 3, 3, 3, 3, 3}, {3, 3, 3, 3, 1, 3, 3, 3, 3, 3}, {1, 1, 1, 11, 19, 1, 1, 1, 1, 1}, {3, 3, 3, 18, 19, 3, 3, 3, 3, 3}, {3, 3, 3, 3, 1, 3, 3, 3, 3, 3}, {3, 3, 3, 3, 1, 3, 3, 3, 3, 3}, {3, 3, 3, 1, 6, 3, 3, 3, 3, 3}, {3, 3, 3, 3, 1, 3, 3, 3, 3, 3}}))
	fmt.Println("Exp:  2; Got: ", equalPairs([][]int{{11, 1}, {1, 1}}))
	fmt.Println("Exp:  4; Got: ", equalPairs([][]int{{13, 13}, {13, 13}}))
	fmt.Println("Exp:  1; Got: ", equalPairs([][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}))
	fmt.Println("Exp:  3; Got: ", equalPairs([][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}))
}
