package binarysearch

import "github.com/itka0526/problems-in-golang/functions"

// my attempt 90/150
// func findMin(nums []int) int {
// 	left, right := 0, len(nums)-1

// 	for left <= right {
// 		mid := (left + right) / 2

// 		if left == right {
// 			return nums[0]
// 		}

// 		if nums[mid] > nums[mid+1] {
// 			return nums[mid+1]
// 		}

// 		if nums[mid] < nums[mid+1] {
// 			left++
// 		}
// 	}

//		return 0
//	}
func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	res := nums[0]

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= nums[0] {
			left++
		} else {
			if nums[mid] < res {
				res = nums[mid]
			}
			right--
		}
	}

	return res
}

// func findMin(nums []int) int {
// 	l, r := 0, len(nums)-1

// 	for l <= r {
// 		m := (l + r) / 2
// 		if l == r {
// 			return nums[l]
// 		}
// 		if nums[l] < nums[m] {
// 			r = m - 1
// 		} else {
// 			l = m + 1
// 		}
// 	}

// 	return -1
// }

func TestFindMin() {
	// functions.PrintSolution(findMin( []int{3, 4, 5, 1, 2}))
	// functions.PrintSolution(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	// functions.PrintSolution(findMin([]int{11, 13, 15, 17}))
	functions.PrintSolution(findMin([]int{3, 1, 2}))
}
