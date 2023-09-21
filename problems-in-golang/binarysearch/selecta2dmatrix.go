package binarysearch

import (
	"github.com/itka0526/problems-in-golang/functions"
)

// You are given an m x n integer matrix matrix with the following two properties:

// Each row is sorted in non-decreasing order.
// The first integer of each row is greater than the last integer of the previous row.
// Given an integer target, return true if target is in matrix or false otherwise.

// You must write a solution in O(log(m * n)) time complexity

// row and col are in increasing order
// hmmm I must use binary search here but I am not sure how to
// i could do for each row do binary search

// simple implementation
func searchMatrix(matrix [][]int, target int) bool {
	for row := 0; row < len(matrix); row++ {
		l, r := 0, len(matrix[row])-1

		for l <= r {
			m := (l + r) / 2

			if matrix[row][m] == target {
				return true
			} else if matrix[row][m] < target {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}

	return false
}

// But can be improved knowing this
// The first integer of each row is greater than the last integer of the previous row.
// This ****** approach does not work well, it did 126/133
// Should we be finding the best row to search using binary seach and then again using BS to find the target?
func searchMatrix1(matrix [][]int, target int) bool {
	// so we gotta try to find the right column first
	t, b := 0, len(matrix)-1

	for t <= b {
		row := (t + b) / 2

		if target > matrix[row][len(matrix[row])-1] {
			t = row + 1
		} else if target < matrix[row][0] {
			b = row - 1
		} else {
			break
		}
	}

	if t > b {
		return false
	}

	row := (t + b) / 2
	l, r := 0, len(matrix[0])-1

	for l <= r {
		m := (l + r) / 2
		if target > matrix[row][m] {
			l = m + 1
		} else if target < matrix[row][m] {
			r = m - 1
		} else {
			return true
		}
	}
	return false
}

func TestSearchMatrix() {
	functions.PrintSolution(
		searchMatrix1([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 11),
		searchMatrix1([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3),
	)
}
