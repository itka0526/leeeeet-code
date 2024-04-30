package dynamicprogramming

func sum(a []int) int {
	s := 0
	for _, n := range a {
		s += n
	}
	return s
}

type Set struct {
	items *map[int]bool
}

func newSet() Set {
	return Set{items: &map[int]bool{}}
}

func (s *Set) insert(x int) int {
	(*s.items)[x] = true
	return x
}

func (s *Set) has(x int) bool {
	return (*s.items)[x]
}

func canPartition(nums []int) bool {
	total := sum(nums)
	if total%2 != 0 {
		return false
	}
	target := total / 2
	s := newSet()
	s.insert(0)
	for i := len(nums) - 1; i > -1; i -= 1 {
		s1 := newSet()
		for t := range *(s.items) {
			if t+nums[i] == target {
				return true
			}
			s1.insert(t + nums[i])
			s1.insert(t)
		}
		s = s1
	}
	return false
}
