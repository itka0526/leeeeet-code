package seventyfive

import "fmt"

func pivotIndex(nums []int) int {
	s := func(a []int) int {
		sum := 0
		for i := range a {
			sum += a[i]
		}
		return sum
	}
	t := s(nums)
	ls := 0
	for i := 0; i < len(nums); i++ {
		if ls*2 == t-nums[i] {
			return i
		}
		ls += nums[i]
	}
	return -1
}

func TestPivotIndex() {
	fmt.Println("Exp: 3; Got: ", pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}
