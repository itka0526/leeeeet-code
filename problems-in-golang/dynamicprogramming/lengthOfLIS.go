package dynamicprogramming

func lengthOfLIS(nums []int) int {
	n := len(nums)
	cache := make([]int, n)
	for i := range cache {
		cache[i] = 1
	}
	for i := len(nums) - 1; i > -1; i -= 1 {
		for j := i + 1; j < n; j++ {
			if nums[i] < nums[j] {
				cache[i] = max(1+cache[j], cache[i])
			}
		}
	}
	ans := cache[0]
	for i := 1; i < len(cache); i++ {
		ans = max(ans, cache[i])
	}
	return ans
}
