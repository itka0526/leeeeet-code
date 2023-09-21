package binarysearch

import "github.com/itka0526/problems-in-golang/functions"

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := (l + r) / 2

		if nums[m] == target {
			return m
		} else if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return -1
}

func TestSearch() {
	functions.PrintSolution(search([]int{-1, 0, 3, 5, 9, 12}, 9))
}
