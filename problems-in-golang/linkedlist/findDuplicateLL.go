package linkedlist

func findDuplicate(nums []int) int {
	// create two pointers
	slow, fast := nums[0], nums[nums[0]]

	// till two pointers collide with each other
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	// reset the first pointer to 0
	slow = 0

	// now till the collided pointer meets the reseted pointer
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
