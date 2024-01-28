package seventyfive

func moveZeroes(nums []int) {
	c := make([]int, len(nums))
	i := 0
	for _, n := range nums {
		if n != 0 {
			c[i] = n
			i++
		}
	}
	copy(nums, c)
}
