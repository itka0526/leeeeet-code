package slidingwindow

import (
	"fmt"
	"math"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]int{}
	for i, e := range nums {
		v, ok := m[e]
		if ok && math.Abs(float64(i-v)) <= float64(k) {
			return true
		} else {
			m[e] = i
		}
	}
	return false
}

func TestContainsNearbyDuplicate() {
	fmt.Println("Exp: true; Got: ", containsNearbyDuplicate([]int{0, 1, 2, 3, 2, 5}, 3))
	fmt.Println("Exp: true; Got: ", containsNearbyDuplicate([]int{99, 99}, 2))
	fmt.Println("Exp: true; Got: ", containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	fmt.Println("Exp: false; Got: ", containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}
