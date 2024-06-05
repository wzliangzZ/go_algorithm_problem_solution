package DataStruct

import (
	"container/heap"
)

/*
对顶堆：在一段区间内，维护两个堆left, right 分别是大根堆和小根堆
用途：求一段区间内的中位数，前一段数(left)之和等等
*/

// 基本思路生成两个堆，通过交换两个堆的根节点，lfet和right满足初始特性
// 在删除时，通过balance参数判断是否需要交换两个堆的根节点

/*
1) 现将区间内的值都将入left
2) 从left堆顶取出指定数量的值加入right;至此，balance=0，即平衡状态，往后都将维持次状态
3）设 加入的数num[i],删除的数num[j]
	将删除的数做懒标记 hm[num[j]]++；
	如果 num[j] <= left[堆顶]，则 删除 的数在left：balance--,否则，balance++
	如果 num[i] <= left[堆顶]，则 加入 的数在left：balance++,否则，balance--
	如果balance < 0, 即 left 少了数，要把right[堆顶]加入left,反之亦然。
4）	for left; for right 判断堆顶是否为懒删除，如果是，将堆顶弹出，hm[num]--
*/

// 设在长度为n的数组中求长度为 dist 的所有中位数
func DigonalHeap(arr []int, dist, n int) {
	var left, right Heap // 小根堆，left操作的值添负号，变成大根堆
	hm := make(map[int]int)
	for i := 0; i < dist; i++ {
		heap.Push(&left, -arr[i])
	}
	for i := 0; i < dist/2; i++ {
		heap.Push(&right, -heap.Pop(&left).(int))
	}

	for i := dist; i < n; i++ {
		var balance int
		rm := arr[i-dist] // 要删除的数
		hm[rm]++
		// 考虑删除数的影响
		if left.Len() > 0 && rm <= -left.Top().(int) {
			balance--
		} else {
			balance++
		}

		// 考虑加入数的影响
		if left.Len() > 0 && arr[i] <= -left.Top().(int) {
			heap.Push(&left, -arr[i])
			balance++
		} else {
			heap.Push(&right, arr[i])
			balance--
		}
		// 考虑是不平衡
		if balance < 0 {
			heap.Push(&left, -heap.Pop(&right).(int))
		}
		if balance > 0 {
			heap.Push(&right, -heap.Pop(&left).(int))
		}

		//清除堆顶的懒标记
		for left.Len() > 0 && hm[-left.Top().(int)] > 0 {
			hm[-left.Top().(int)]--
			heap.Pop(&left)
		}
		for right.Len() > 0 && hm[right.Top().(int)] > 0 {
			hm[right.Top().(int)]--
			heap.Pop(&right)
		}
	}
}
