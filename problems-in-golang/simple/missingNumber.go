package simple

import "fmt"

func missingNumber(nums []int) int {
	l := len(nums)
	sum := (l + 1) * l / 2
	realSum := 0
	for i := 0; i < l; i++ {
		realSum += nums[i]
	}
	return sum - realSum
}

func TestMissingNumber() {
	s := []int{3, 0}
	fmt.Println(missingNumber(s))
}
