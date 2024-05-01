package bitmanipulation

func singleNumber(nums []int) (ans int) {
	for i := 0; i < len(nums); i++ {
		ans ^= nums[i]
	}
	return ans
}
