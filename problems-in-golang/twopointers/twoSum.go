package twopointers

import (
	"fmt"
)

// func twoSumSolution(numbers []int, target int) []int {

// }

func twoSumSolution(numbers []int, target int) []int {
	// nums - 2,7,11,15 | target - 9
	i, j := 0, len(numbers)-1

	for {
		sum := numbers[i] + numbers[j]

		if sum == target {
			return []int{i + 1, j + 1}
		}

		if sum < target {
			i++
			continue
		}

		if sum > target {
			j--
			continue
		}
	}
}

func TestTwoSum() {
	// nums := []int{2, 7, 11, 15}
	// fmt.Println(twoSumSolution(nums, 9))

	// nums2 := []int{-3, 3, 4, 90}
	// fmt.Println(twoSumSolution(nums2, 0))

	nums3 := []int{5, 25, 75}
	fmt.Println(twoSumSolution(nums3, 100))
}
