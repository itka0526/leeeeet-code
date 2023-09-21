package binarysearch

import "github.com/itka0526/problems-in-golang/functions"

func searchRotatedSortedArray(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		// do simple BS on the left beacause its sorted
		if nums[mid] >= nums[left] {
			if target > nums[mid] || target < nums[left] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			// right
			if target < nums[mid] || target > nums[right] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}

func TestSearchRotatedSortedArray() {
	functions.PrintSolution(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}
