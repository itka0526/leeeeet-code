package simple

import "fmt"

func addRange(n []int, i, j int) string {
	if i == j {
		return fmt.Sprint(n[i])
	}
	return fmt.Sprint(n[i], "->", n[j])
}

func summaryRanges(nums []int) (ans []string) {
	m := 0
	p := 1
	for p < len(nums) {
		if nums[p-1]+1 != nums[p] {
			ans = append(ans, addRange(nums, m, p-1))
			m = p
		}
		p++
	}
	if len(nums) > 0 {
		ans = append(ans, addRange(nums, m, p-1))
	}
	return ans
}

func TestSummaryRanges() {
	n1 := []int{0, 1}
	fmt.Println(summaryRanges(n1))
	// n := []int{0, 1, 2, 4, 5, 7}
	// fmt.Println(summaryRanges(n))
}
