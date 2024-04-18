package simple

func majorityElement(nums []int) int {
	m := (len(nums) + 1) / 2
	nc := make(map[int]int)
	for _, n := range nums {
		nc[n] += 1
	}
	for num, count := range nc {
		if count >= m {
			return num
		}
	}
	return -1
}
