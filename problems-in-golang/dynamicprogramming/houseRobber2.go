package dynamicprogramming

func rob2(nums []int) int {
	if len(nums) == 0 {
		return nums[0]
	}
	return max(helper(nums[1:]), helper(nums[:len(nums)-1]))
}

func helper(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = nums[1]
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[n-1]
}
