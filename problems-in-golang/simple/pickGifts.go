package simple

import (
	"container/heap"
	"fmt"
	"math"
)

type MaxHeap []int

func (a MaxHeap) Len() int           { return len(a) }
func (a MaxHeap) Less(i, j int) bool { return a[i] > a[j] }
func (a MaxHeap) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a *MaxHeap) Push(x any)        { *a = append(*a, x.(int)) }

func (a *MaxHeap) Pop() any {
	i := len(*a) - 1
	t := (*a)[i]
	*a = (*a)[:i]
	return t
}

func pickGifts(gifts []int, k int) (ans int64) {
	h := MaxHeap(gifts)
	heap.Init(&h)
	for k > 0 {
		m := heap.Pop(&h).(int)
		m = int(math.Sqrt(float64(m)))
		heap.Push(&h, m)
		k--
	}
	for i := 0; i < h.Len(); i++ {
		ans += int64(h[i])
	}
	return ans
}

func TestPickGifts() {
	s := []int{25, 64, 9, 4, 100}
	fmt.Println(pickGifts(s, 4))
	s1 := []int{1, 1, 1, 1}
	fmt.Println(pickGifts(s1, 4))
}
