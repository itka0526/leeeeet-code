package stack

import "fmt"

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (t *MinStack) Push(val int) {
	t.stack = append(t.stack, val)

	// currVal := t.GetMin()
	if len(t.minStack) == 0 {
		t.minStack = append(t.minStack, val)
	} else {
		t.minStack = append(t.minStack, minVal(val, t.GetMin()))
	}
}

func minVal(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (t *MinStack) Pop() {
	t.stack = t.stack[:len(t.stack)-1]
	t.minStack = t.minStack[:len(t.minStack)-1]
}

func (t *MinStack) Top() int {
	return t.stack[len(t.stack)-1]
}

func (t *MinStack) GetMin() int {
	return t.minStack[len(t.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func TestMinStack() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin()) // return -3
	minStack.Pop()
	fmt.Println(minStack.Top())    // return 0
	fmt.Println(minStack.GetMin()) // return -2
}
