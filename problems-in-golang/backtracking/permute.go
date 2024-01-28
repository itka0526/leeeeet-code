package backtracking

import (
	"fmt"
)

func permute(nums []int) [][]int {
	res := [][]int{}
	arrayOfBool := make([]bool, len(nums))

	var dfs func([]int)

	dfs = func(curr []int) {
		if len(curr) == len(nums) {
			cp := make([]int, len(curr))
			copy(cp, curr)
			res = append(res, cp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if !arrayOfBool[i] {
				arrayOfBool[i] = true
				curr = append(curr, nums[i])
				dfs(curr)
				arrayOfBool[i] = false
				curr = curr[:len(curr)-1]
			}
		}
	}

	dfs([]int{})

	return res
}

func TestPermute() {
	// void Backtrack (res, args)
	//		if ( GOAL REACHED )
	//			add solution to res
	// 			return
	//		for ( i = 0; i < n_choices; i++ )
	// 			if ( choices[i] is valid )
	// 				make choice
	// 				backtrack( res, args )
	// 				undo choices[i]

	fmt.Println(permute([]int{1, 2, 3}))
}
