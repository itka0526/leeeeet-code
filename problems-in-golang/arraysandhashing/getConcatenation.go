package arraysandhashing

func getConcatenation(nums []int) []int {
	nums = append(nums, nums...)
	return nums
}
