package seventyfive

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	// Time Limit Exceeded
	// for i := 0; i < len(nums); i++ {
	// 	for j := i + 1; j < len(nums); j++ {
	// 		if nums[i] < nums[j] {
	// 			for k := j + 1; k < len(nums); k++ {
	// 				if nums[j] < nums[k] {
	// 					return true
	// 				}
	// 			}
	// 		}
	// 	}
	// }
	// return false

	// We don't need to check each element is in increasing order because
	// if there is a number lesser than the already updated number
	// we can just ignore and still do the update
	// This logic is not intuitave
	// For example: 20, 100, 10, 12, 5, 13
	// a is gonna be 20
	// then 10 and then 5
	// THIS LOGIC
	//
	a, b := math.MaxInt32, math.MaxInt32
	for _, num := range nums {
		if num <= a {
			a = num
		} else if num <= b {
			b = num
		} else {
			return true
		}
	}
	return false
}

func TestIncreasingTriplet() {
	fmt.Println("Exp: true; Got: ", increasingTriplet([]int{20, 100, 10, 12, 5, 13}))
	fmt.Println("Exp: true; Got: ", increasingTriplet([]int{1, 2, 3, 4, 5}))
	fmt.Println("Exp: false; Got: ", increasingTriplet([]int{5, 4, 3, 2, 1}))
}
