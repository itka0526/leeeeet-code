package seventyfive

import "fmt"

func productExceptSelf(nums []int) []int {
	po := make([]int, len(nums))
	pe := make([]int, len(nums))
	re := make([]int, len(nums))
	// Somehow keep the ahead products
	// [1,2,3,4]
	// 2 * 3 * 4 - 0 = 24
	// 3 * 4	 - 1 = 12
	// 4 		 - 2 = 4
	// 			 - 3 = 1
	for i, k := len(nums)-1, 1; i >= 0; i-- {
		po[i] = k
		k = nums[i] * k
	}
	// Somehow keep the before products
	// [1,2,3,4]
	// 			 - 0 = 1
	// 1		 - 1 = 1
	// 1 * 2	 - 2 = 2
	// 1 * 2 * 3 - 3 = 6
	for i, k := 0, 1; i < len(nums); i++ {
		pe[i] = k
		k = nums[i] * k
	}
	for i := 0; i < len(nums); i++ {
		re[i] = po[i] * pe[i]
	}
	return re
}

func TestProductExpectSelf() {
	fmt.Println("Expected: [24,12,8,6]; Got: ", productExceptSelf([]int{1, 2, 3, 4}))
}
