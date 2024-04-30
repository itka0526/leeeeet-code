package dynamicprogramming

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxSoFar := nums[0]
	minSoFar := nums[0]
	result := maxSoFar
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxSoFar, minSoFar = minSoFar, maxSoFar
		}
		maxSoFar = max(nums[i], maxSoFar*nums[i])
		minSoFar = min(nums[i], minSoFar*nums[i])
		result = max(result, maxSoFar)
	}
	return result
}
