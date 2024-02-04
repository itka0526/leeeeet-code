package seventyfive

import "fmt"

func findDifference(nums1 []int, nums2 []int) [][]int {
	m1, m2 := map[int]bool{}, map[int]bool{}
	for i := range nums1 {
		m1[nums1[i]] = true
	}
	for i := range nums2 {
		m2[nums2[i]] = true
	}
	for k1 := range m1 {
		if _, ok := m2[k1]; ok {
			delete(m1, k1)
			delete(m2, k1)
		}
	}
	res := make([][]int, 2)
	r1, r2 := make([]int, len(m1)), make([]int, len(m2))
	i := 0
	for k1 := range m1 {
		r1[i] = k1
		i++
	}
	i = 0
	for k2 := range m2 {
		r2[i] = k2
		i++
	}
	res[0] = r1
	res[1] = r2
	return res
}

func TestFindDifference() {
	fmt.Println("Exp: [[-94 82 21 -43 -68 -80 -19] [-63]]; Got: ", findDifference([]int{-68, -80, -19, -94, 82, 21, -43}, []int{-63}))
	fmt.Println("Exp: [[1 3] [4 6]]; Got: ", findDifference([]int{1, 2, 3}, []int{2, 4, 6}))
}
