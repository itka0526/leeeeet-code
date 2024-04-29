package medium

func minOperations(nums []int, k int) (ans int) {
	curr := nums[0]
	for i := 1; i < len(nums); i++ {
		curr = curr ^ nums[i]
	}
	for curr > 0 || k > 0 {
		if curr%2 != k%2 {
			ans += 1
		}
		curr /= 2
		k /= 2
	}
	return ans
}
