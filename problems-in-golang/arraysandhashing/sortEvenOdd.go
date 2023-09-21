package arraysandhashing

func sortEvenOdd(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i%2 == 0 && j%2 == 0 && nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
			if i%2 == 1 && j%2 == 1 && nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}
