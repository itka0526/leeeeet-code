package binarysearch

import (
	"math"

	"github.com/itka0526/problems-in-golang/functions"
)

func canEat(piles []int, timeLimit, speed int) bool {
	hoursItWillTake := 0

	for _, pile := range piles {
		hoursItWillTake += (pile + speed - 1) / speed
	}

	return hoursItWillTake <= timeLimit
}

// BS approach
func minEatingSpeed1(piles []int, h int) int {
	lo, hi, ans := 1, 1_000_000_000, 1

	for lo <= hi {
		mid := (lo + hi) / 2
		if canEat(piles, h, mid) {
			ans = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return ans
}

// Naive approach 116/125
func minEatingSpeed(piles []int, h int) int {
	speed := 1

	for {
		hoursItWillTake := 0
		for _, pile := range piles {
			hoursItWillTake += int(math.Ceil(float64(pile) / float64(speed)))
		}
		if hoursItWillTake <= h {
			break
		}
		speed++
	}

	return speed
}

func TestMinEatingSpeed() {
	functions.PrintSolution(minEatingSpeed1([]int{3, 6, 7, 11}, 8))
}
