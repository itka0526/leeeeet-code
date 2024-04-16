package simple

import "fmt"

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		t := make([]int, i+1)
		for j := 0; j < len(t); j++ {
			t[j] = 1
		}
		res[i] = t
	}
	for j := 2; j < numRows; j++ {
		for k := 1; k < len(res[j])-1; k++ {
			res[j][k] = res[j-1][k-1] + res[j-1][k]
		}
	}
	return res
}

func TestGenerate() {
	numRows := 5
	fmt.Println(generate(numRows))
}
