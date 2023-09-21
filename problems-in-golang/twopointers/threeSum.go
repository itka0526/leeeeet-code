package twopointers

import (
	"fmt"
	"sort"
)

// func threeSumSolution(nums []int) [][]int {
// 	res := make([][]int, 0)
// 	m := make(map[string][]int)

// 	for i := range nums {
// 		for j := range nums {
// 			for k := range nums {
// 				if i != j && i != k && j != k && nums[i]+nums[j]+nums[k] == 0 {
// 					tmp := []int{nums[i], nums[j], nums[k]}

// 					sort.Ints(tmp)

// 					m[fmt.Sprint(tmp)] = tmp
// 				}
// 			}
// 		}
// 	}

// 	for key := range m {
// 		res = append(res, m[key])
// 	}

// 	return res
// }

func threeSumSolution(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)

	for x := range nums {
		if x > 0 && nums[x] == nums[x-1] {
			continue
		}
		y, z := x+1, len(nums)-1
		for y < z {
			sum := nums[x] + nums[y] + nums[z]

			if sum > 0 {
				z--
			}

			if sum < 0 {
				y++
			}

			if sum == 0 {
				res = append(res, []int{nums[x], nums[y], nums[z]})
				y++

				// This means that if the left pointer is the same as before just skip till
				// Left pointer is something different
				// And also right pointer must always be greater than left
				for nums[y-1] == nums[y] && y < z {
					y++
				}
			}
		}
	}

	return res
}

// func threeSumSolution(nums []int) [][]int {
// 	// a + b + c = 0,
// 	res := make([][]int, 0)
// 	sort.Ints(nums)

// 	for a := range nums {
// 		if a > 0 && nums[a] == nums[a-1] {
// 			continue
// 		}
// 		// -1, 0, 1, 2, -1, -4
// 		// since I got the first element
// 		// let us proceed
// 		// i dont get how did he figure out to use this technic but ok
// 		l, r := a+1, len(nums)-1

// 		for l < r {
// 			sum := nums[a] + nums[l] + nums[r]

// 			// -2 0 1 2 3 => (a: -2) + (l: 0) + (r: 2) = 0
// 			if sum > 0 {
// 				r--
// 			} else if sum < 0 {
// 				l++
// 			} else {
// 				res = append(res, []int{nums[a], nums[l], nums[r]})
// 				l++

// 				for nums[l] == nums[l-1] && l < r {
// 					l++
// 				}
// 			}

// 		}

// 	}

// 	return res
// }

// func threeSumSolution(nums []int) [][]int {
// 	// Maybe a map is a better way to solve this problem because nums can be too long
// 	// No impossible... I need the indexes right?
// 	sort.Ints(nums)

// 	res := make([][]int, 0)

// 	for i := range nums {
// 		if i > 0 && nums[i] == nums[i-1] {
// 			continue
// 		}

// 		l, r := i+1, len(nums)-1

// 		for l < r {
// 			sum := nums[i] + nums[l] + nums[r]

// 			if sum > 0 {
// 				r--
// 			} else if sum < 0 {
// 				l++
// 			} else {
// 				res = append(res, []int{nums[i], nums[l], nums[r]})

// 				l += 1
// 				for nums[l] == nums[l-1] && l < r {
// 					l++
// 				}
// 			}
// 		}
// 	}

// 	return res
// }

func TestThreeSumSolution() {
	nums := []int{-1, 0, 1, 2, -1, -4}

	fmt.Println("Solution: ", threeSumSolution(nums))
}
