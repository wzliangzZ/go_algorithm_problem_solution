package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF173B(in *bufio.Reader, out *bufio.Writer) {
	// var T int
	// for fmt.Fscan(in, &T); T > 0; T-- {

	// }
	const inf = 0x3f3f3f3f
	var n, m int
	fmt.Fscan(in, &n, &m)
	g := make([]string, n)
	f := make([][][]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &g[i])
		f[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			f[i][j] = make([]int, 4)
			for k := 0; k < 4; k++ {
				f[i][j][k] = inf
			}
		}
	}
	vi := []int{-1, 0, 1, 0}
	vj := []int{0, 1, 0, -1}

	dq := NewArrayDeque(6_000_000)
	f[0][0][1] = 0
	dq.PushBack(element{0, 0, 1})
	for !dq.IsEmpty() {
		t := dq.PopFront().(element)
		a, b, d := t.x, t.y, t.d
		nea, neb := a+vi[d], b+vj[d]
		if nea >= 0 && nea < n && neb >= 0 && neb < m && f[nea][neb][d] > f[a][b][d] {
			f[nea][neb][d] = f[a][b][d]
			dq.PushFront(element{nea, neb, d})
		}
		if g[a][b] == '#' {
			for k := 0; k < 4; k++ {
				if f[a][b][d]+1 < f[a][b][k] && k != d {
					f[a][b][k] = f[a][b][d] + 1
					dq.PushBack(element{a, b, k})
				}
			}
		}
	}
	if f[n-1][m-1][1] != inf {
		fmt.Fprintln(out, f[n-1][m-1][1])
	} else {
		fmt.Fprintln(out, -1)
	}
}

// func main() {
// 	in := bufio.NewReader(os.Stdin)
// 	out := bufio.NewWriter(os.Stdout)
// 	defer out.Flush()

// 	CF173B(in, out)
// }

type element struct {
	x, y, d int
}

/* 基于环形数组实现的双向队列 */
type arrayDeque struct {
	elements    []element // 用于存储双向队列元素的数组
	front       int       // 队首指针，指向队首元素
	queSize     int       // 双向队列长度
	queCapacity int       // 队列容量（即最大容纳元素数量）
}

/* 初始化队列 */
func NewArrayDeque(queCapacity int) *arrayDeque {
	return &arrayDeque{
		elements:    make([]element, queCapacity),
		queCapacity: queCapacity,
		front:       0,
		queSize:     0,
	}
}

/* 获取双向队列的长度 */
func (q *arrayDeque) Size() int {
	return q.queSize
}

/* 判断双向队列是否为空 */
func (q *arrayDeque) IsEmpty() bool {
	return q.queSize == 0
}

/* 计算环形数组索引 */
func (q *arrayDeque) Index(i int) int {
	// 通过取余操作实现数组首尾相连
	// 当 i 越过数组尾部后，回到头部
	// 当 i 越过数组头部后，回到尾部
	return (i + q.queCapacity) % q.queCapacity
}

/* 队首入队 */
func (q *arrayDeque) PushFront(ele element) {
	if q.queSize == q.queCapacity {
		fmt.Println("双向队列已满")
		return
	}
	// 队首指针向左移动一位
	// 通过取余操作实现 front 越过数组头部后回到尾部
	q.front = q.Index(q.front - 1)
	// 将 element 添加至队首
	q.elements[q.front] = ele
	q.queSize++
}

/* 队尾入队 */
func (q *arrayDeque) PushBack(ele element) {
	if q.queSize == q.queCapacity {
		fmt.Println("双向队列已满")
		return
	}
	// 计算队尾指针，指向队尾索引 + 1
	rear := q.Index(q.front + q.queSize)
	// 将 element 添加至队首
	q.elements[rear] = ele
	q.queSize++
}

/* 队首出队 */
func (q *arrayDeque) PopFront() any {
	ele := q.Front()
	// 队首指针向后移动一位
	q.front = q.Index(q.front + 1)
	q.queSize--
	return ele
}

/* 队尾出队 */
func (q *arrayDeque) PopBcck() any {
	ele := q.Back()
	q.queSize--
	return ele
}

/* 访问队首元素 */
func (q *arrayDeque) Front() any {
	if q.IsEmpty() {
		return nil
	}
	return q.elements[q.front]
}

/* 访问队尾元素 */
func (q *arrayDeque) Back() any {
	if q.IsEmpty() {
		return nil
	}
	// 计算尾元素索引
	last := q.Index(q.front + q.queSize - 1)
	return q.elements[last]
}
