package DataStruct

/*
go lib: "container/heap"

func Init(h Interface)					将 h 变成堆				  	O(n)
func Remove(h Interface, i int) any		删除并返回下标为 i 的元素	 O(logn)

*/

//创建一个堆要实现5个函数：3个排序函数(Len,Less,Swap)，2个增删函数(Push,Pop)

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }

func (h *Heap) Push(x any) { *h = append(*h, x.(int)) }

func (h *Heap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return x
}

//可有可无
func (h *Heap) Top() any {
	return (*h)[0]
}
