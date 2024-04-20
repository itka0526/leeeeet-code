package simple

func intersection(nums1 []int, nums2 []int) (ans []int) {
	a := make(map[int]bool)
	b := make(map[int]bool)
	for i := range nums1 {
		a[nums1[i]] = true
	}
	for i := range nums2 {
		b[nums2[i]] = true
	}
	if len(a) < len(b) {
		for k := range a {
			if _, ok := b[k]; ok {
				ans = append(ans, k)
			}
		}
	} else {
		for k := range b {
			if _, ok := a[k]; ok {
				ans = append(ans, k)
			}
		}
	}
	return ans
}
