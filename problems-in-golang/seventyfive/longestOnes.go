package seventyfive

import "fmt"

func longestOnes(nums []int, k int) int {
	l, r := 0, 0
	for r < len(nums) {
		if nums[r] == 0 {
			k--
		}
		if k < 0 {
			if nums[l] == 0 {
				k++
			}
			l++
		}
		r++
	}
	return r - l
}

func TestLongestOnes() {
	fmt.Println("Exp: 10; Got: ", longestOnes([]int{0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))
}
