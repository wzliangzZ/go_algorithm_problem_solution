package api

import (
	"container/heap"
	"fmt"
)

// 创建一个堆要实现5个函数：3个排序函数(Len,Less,Swap)，2个增删函数(Push,Pop)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func THeap() {
	// 创建一个数组
	h := &IntHeap{2, 1, 5}
	// O(n) 时间将数组初始化为堆
	heap.Init(h)

	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}

	// O(log n)删除下标为1的值
	// heap.Remove(h, 1)
	// 在改变下标对应值时，O(log n)重序堆
	// (*h)[2] = 10
	// heap.Fix(h, 2)
}
