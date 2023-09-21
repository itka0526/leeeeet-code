package simple

func findFinalValue(nums []int, original int) int {
	// sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	// for _, n := range nums {
	// 	if n == original {
	// 		original += original
	// 	}
	// }
	// return original

	s := make([]int, 1001)

	for _, n := range nums {
		s[n]++
	}

	for original <= 1001 && s[original] != 0 {
		s[original] = 0
		original += original
	}
	return original
}
