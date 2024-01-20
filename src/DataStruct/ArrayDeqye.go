package DataStruct

/*
创建		NewArrayDeque(queCapacity int) *arrayDeque
取值		Back()		Front()		Len()
插入		PushBack()	PushFront()
取值 删除	 PopFront()	 PopBcck()
*/
type Element struct {
	first, second int
}

/* 基于环形数组(容量固定)实现的双向队列 */
type ArrayDeque struct {
	elements    []Element // 用于存储双向队列元素的数组
	front       int       // 队首指针，指向队首元素
	queSize     int       // 双向队列长度
	queCapacity int       // 队列容量（即最大容纳元素数量）
}

/* 初始化队列 */
func NewArrayDeque(queCapacity int) *ArrayDeque {
	return &ArrayDeque{
		elements:    make([]Element, 0, queCapacity),
		queCapacity: queCapacity,
		front:       0,
		queSize:     0,
	}
}

/* 获取双向队列的长度 */
func (q *ArrayDeque) Len() int {
	return q.queSize
}

/* 判断双向队列是否为空 */
func (q *ArrayDeque) IsEmpty() bool {
	return q.queSize == 0
}

/* 计算环形数组索引 */
func (q *ArrayDeque) index(i int) int {
	// 通过取余操作实现数组首尾相连
	// 当 i 越过数组尾部后，回到头部
	// 当 i 越过数组头部后，回到尾部
	return (i + q.queCapacity) % q.queCapacity
}

/* 队首入队 */
func (q *ArrayDeque) PushFront(ele Element) {
	if q.queSize == q.queCapacity {
		panic("双向队列已满")
	}
	// 队首指针向左移动一位
	// 通过取余操作实现 front 越过数组头部后回到尾部
	q.front = q.index(q.front - 1)
	// 将 Element 添加至队首
	q.elements[q.front] = ele
	q.queSize++
}

/* 队尾入队 */
func (q *ArrayDeque) PushBack(ele Element) {
	if q.queSize == q.queCapacity {
		panic("双向队列已满")
	}
	// 计算队尾指针，指向队尾索引 + 1
	rear := q.index(q.front + q.queSize)
	// 将 Element 添加至队首
	q.elements[rear] = ele
	q.queSize++
}

/* 访问队首元素 */
func (q *ArrayDeque) Front() any {
	if q.IsEmpty() {
		return nil
	}
	return q.elements[q.front]
}

/* 访问队尾元素 */
func (q *ArrayDeque) Back() any {
	if q.IsEmpty() {
		return nil
	}
	// 计算尾元素索引
	last := q.index(q.front + q.queSize - 1)
	return q.elements[last]
}

/* 队首出队 */
func (q *ArrayDeque) PopFront() any {
	ele := q.Front()
	// 队首指针向后移动一位
	q.front = q.index(q.front + 1)
	q.queSize--
	return ele
}

/* 队尾出队 */
func (q *ArrayDeque) PopBcck() any {
	ele := q.Back()
	q.queSize--
	return ele
}
