package backtracking

import (
	"fmt"
	"sort"

	"github.com/itka0526/problems-in-golang/functions"
)

func combinationSum2HalfSolution(candidates []int, target int) (ans [][]int) {
	var findSums func([]int, int)
	record := map[string]bool{}
	findSums = func(a []int, idx int) {
		if sum(a) == target {
			key := constructKey(a)
			if record[key] {
				return
			}
			record[key] = true
			ans = append(ans, a)
			return
		}
		if idx > len(candidates)-1 {
			return
		}
		findSums(a, idx+1)
		a = append(a, candidates[idx])
		findSums(a, idx+1)
	}
	findSums([]int{}, 0)
	return ans
}

func sum(a []int) int {
	s := 0
	for _, x := range a {
		s += x
	}
	return s
}

func constructKey(a []int) string {
	return fmt.Sprint(a)
}

func combinationSum2TLE(candidates []int, target int) (ans [][]int) {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	record := map[string]bool{}
	var bt func([]int, int, int)
	bt = func(a []int, idx int, curr int) {
		if curr > target {
			return
		}
		if curr == target {
			k := constructKey(a)
			if record[k] {
				return
			}
			record[k] = true
			cp := make([]int, len(a))
			copy(cp, a)
			ans = append(ans, cp)
			return
		}
		if idx > len(candidates)-1 {
			return
		}
		bt(a, idx+1, curr)
		a = append(a, candidates[idx])
		bt(a, idx+1, curr+candidates[idx])
	}
	bt([]int{}, 0, 0)
	return ans
}

func combinationSum2(candidates []int, target int) (ans [][]int) {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	var bt func([]int, int, int)
	bt = func(curr []int, pos int, targ int) {
		if targ == 0 {
			cp := make([]int, len(curr))
			copy(cp, curr)
			ans = append(ans, cp)
		}
		if targ <= 0 {
			return
		}
		prev := -1
		for i := pos; i < len(candidates); i++ {
			if candidates[i] == prev {
				continue
			}
			curr = append(curr, candidates[i])
			bt(curr, i+1, targ-candidates[i])
			curr = curr[:len(curr)-1]
			prev = candidates[i]
		}
	}
	bt([]int{}, 0, target)
	return ans
}

func combinationSum2FromMemory(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool { return candidates[i] < candidates[j] })
	ans := [][]int{}
	var bt func([]int, int, int)
	bt = func(curr []int, pos int, sum int) {
		if sum > target {
			return
		} else if sum == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			ans = append(ans, cp)
			return
		}
		prev := -1
		for i := pos; i < len(candidates); i++ {
			if candidates[i] == prev {
				continue
			}
			curr = append(curr, candidates[i])
			bt(curr, i+1, sum+candidates[i])
			curr = curr[:len(curr)-1]
			prev = candidates[i]
		}
	}
	bt([]int{}, 0, 0)
	return ans
}

func TestCombinationSum2() {
	s2 := []int{3, 1, 3, 5, 1, 1}
	functions.LogAnswer(combinationSum2FromMemory(s2, 8), "[[1,1,1,5],[1,1,3,3],[3,5]]")

	s := []int{10, 1, 2, 7, 6, 1, 5}
	functions.LogAnswer(combinationSum2FromMemory(s, 8), "[[1,1,6], [1,2,5], [1,7], [2,6]]")

	s1 := []int{2, 5, 2, 1, 2}
	functions.LogAnswer(combinationSum2FromMemory(s1, 5), "[[1,2,2],[5]]")
}
